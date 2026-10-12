package logging

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"
)

func TestNewWritesJSON(t *testing.T) {
	var output bytes.Buffer

	New(&output, slog.LevelInfo).Info("ready", "component", "test")

	for _, want := range []string{`"level":"INFO"`, `"msg":"ready"`, `"component":"test"`} {
		if !strings.Contains(output.String(), want) {
			t.Fatalf("log output %q does not contain %q", output.String(), want)
		}
	}
}
