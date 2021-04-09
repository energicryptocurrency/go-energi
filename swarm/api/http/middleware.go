// Copyright 2021 The Energi Core Authors
// Copyright 2018 The go-ethereum Authors
// This file is part of the Energi Core library.
//
// The Energi Core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Energi Core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Energi Core library. If not, see <http://www.gnu.org/licenses/>.

package http

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"energi.world/core/gen3/metrics"
	"energi.world/core/gen3/swarm/api"
	"energi.world/core/gen3/swarm/log"
	"energi.world/core/gen3/swarm/sctx"
	"energi.world/core/gen3/swarm/spancontext"
	"github.com/pborman/uuid"
)

// Adapt chains h (main request handler) main handler to adapters (middleware handlers)
// Please note that the order of execution for `adapters` is FIFO (adapters[0] will be executed first)
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for i := range adapters {
		adapter := adapters[len(adapters)-1-i]
		h = adapter(h)
	}
	return h
}

type Adapter func(http.Handler) http.Handler

func SetRequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(SetRUID(r.Context(), uuid.New()[:8]))
		metrics.GetOrRegisterCounter(fmt.Sprintf("http.request.%s", r.Method), nil).Inc(1)
		log.Info("created ruid for request", "ruid", GetRUID(r.Context()), "method", r.Method, "url", r.RequestURI)

		h.ServeHTTP(w, r)
	})
}

func SetRequestHost(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(sctx.SetHost(r.Context(), r.Host))
		log.Info("setting request host", "ruid", GetRUID(r.Context()), "host", sctx.GetHost(r.Context()))

		h.ServeHTTP(w, r)
	})
}

func ParseURI(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri, err := api.Parse(strings.TrimLeft(r.URL.Path, "/"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			respondError(w, r, fmt.Sprintf("invalid URI %q", r.URL.Path), http.StatusBadRequest)
			return
		}
		if uri.Addr != "" && strings.HasPrefix(uri.Addr, "0x") {
			uri.Addr = strings.TrimPrefix(uri.Addr, "0x")

			msg := fmt.Sprintf(`The requested hash seems to be prefixed with '0x'. You will be redirected to the correct URL within 5 seconds.<br/>
			Please click <a href='%[1]s'>here</a> if your browser does not redirect you within 5 seconds.<script>setTimeout("location.href='%[1]s';",5000);</script>`, "/"+uri.String())
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(msg))
			return
		}

		ctx := r.Context()
		r = r.WithContext(SetURI(ctx, uri))
		log.Debug("parsed request path", "ruid", GetRUID(r.Context()), "method", r.Method, "uri.Addr", uri.Addr, "uri.Path", uri.Path, "uri.Scheme", uri.Scheme)

		h.ServeHTTP(w, r)
	})
}

func InitLoggingResponseWriter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tn := time.Now()

		writer := newLoggingResponseWriter(w)
		h.ServeHTTP(writer, r)

		ts := time.Since(tn)
		log.Info("request served", "ruid", GetRUID(r.Context()), "code", writer.statusCode, "time", ts)
		metrics.GetOrRegisterResettingTimer(fmt.Sprintf("http.request.%s.time", r.Method), nil).Update(ts)
		metrics.GetOrRegisterResettingTimer(fmt.Sprintf("http.request.%s.%d.time", r.Method, writer.statusCode), nil).Update(ts)
	})
}

func InstrumentOpenTracing(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := GetURI(r.Context())
		if uri == nil || r.Method == "" || (uri != nil && uri.Scheme == "") {
			h.ServeHTTP(w, r) // soft fail
			return
		}
		spanName := fmt.Sprintf("http.%s.%s", r.Method, uri.Scheme)
		ctx, sp := spancontext.StartSpan(r.Context(), spanName)

		defer sp.Finish()
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RecoverPanic(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Error("panic recovery!", "stack trace", string(debug.Stack()), "url", r.URL.String(), "headers", r.Header)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
