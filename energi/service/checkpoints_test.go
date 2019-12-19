package service

import (
	"testing"
)

func TestCheckpointsService(t *testing.T) {
	if val, ok := os.LookupEnv("SKIP_KNOWN_FAIL"); ok && val == "1" {
		t.Skip("unit test is broken: conditional test skipping activated")
	}

	t.Fail()

	// TODO: test coverage is required
}
