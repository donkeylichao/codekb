package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yanglichao/codekb/internal/exitcode"
)

func TestRunHelp(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := Run([]string{"--help"}, &stdout, &stderr, "test")

	if code != exitcode.Success {
		t.Fatalf("Run() code = %d, want %d; stderr = %q", code, exitcode.Success, stderr.String())
	}
	if !strings.Contains(stdout.String(), "Build and query a local code knowledge base") {
		t.Fatalf("help output missing description: %q", stdout.String())
	}
	if strings.Contains(stdout.String(), "init") {
		t.Fatalf("help advertises an unimplemented command: %q", stdout.String())
	}
}

func TestRunRejectsUnexpectedArguments(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := Run([]string{"unexpected"}, &stdout, &stderr, "test")

	if code != exitcode.Failure {
		t.Fatalf("Run() code = %d, want %d", code, exitcode.Failure)
	}
	if !strings.Contains(stderr.String(), "unknown command") {
		t.Fatalf("stderr = %q, want unknown command error", stderr.String())
	}
}
