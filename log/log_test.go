package log

import (
	"os"

	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel(InfoLevel)
	if debugLog.Writer() == os.Stdout || infoLog.Writer() != os.Stdout {
		t.Fatal("failed to set log level")
	}

	SetLevel(ErrorLevel)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() != os.Stdout {
		t.Fatal("failed to set log level")
	}

	SetLevel(Disabled)
	if infoLog.Writer() == os.Stdout || errorLog.Writer() == os.Stdout {
		t.Fatal("failed to set log level")
	}
}
