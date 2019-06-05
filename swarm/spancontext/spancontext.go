// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package spancontext

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"
)

func WithContext(ctx context.Context, sctx opentracing.SpanContext) context.Context {
	return context.WithValue(ctx, "span_context", sctx)
}

func FromContext(ctx context.Context) opentracing.SpanContext {
	sctx, ok := ctx.Value("span_context").(opentracing.SpanContext)
	if ok {
		return sctx
	}

	return nil
}

func StartSpan(ctx context.Context, name string) (context.Context, opentracing.Span) {
	tracer := opentracing.GlobalTracer()

	sctx := FromContext(ctx)

	var sp opentracing.Span
	if sctx != nil {
		sp = tracer.StartSpan(
			name,
			opentracing.ChildOf(sctx))
	} else {
		sp = tracer.StartSpan(name)
	}

	nctx := context.WithValue(ctx, "span_context", sp.Context())

	return nctx, sp
}

func StartSpanFrom(name string, sctx opentracing.SpanContext) opentracing.Span {
	tracer := opentracing.GlobalTracer()

	sp := tracer.StartSpan(
		name,
		opentracing.ChildOf(sctx))

	return sp
}
