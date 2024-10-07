package teststringerrs

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"testing"

	t1 "github.com/senforsce/tndr"
	"github.com/senforsce/tndr/generator/htmldiff"
)

//go:embed expected.html
var expected string

func Test(t *testing.T) {
	t.Run("can render without error", func(t *testing.T) {
		component := render(nil)

		_, err := htmldiff.Diff(component, expected)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
	t.Run("string expressions can return errors", func(t *testing.T) {
		errSomethingBad := errors.New("bad error")

		err := render(errSomethingBad).Render(context.Background(), &bytes.Buffer{})
		if err == nil {
			t.Fatalf("expected error, but got nil")
		}

		t.Run("the errors are t1 errors", func(t *testing.T) {
			var templateErr t1.Error
			if !errors.As(err, &templateErr) {
				t.Fatalf("expected error to be t1.Error, but got %T", err)
			}
			if templateErr.FileName != `generator/test-string-errors/template.t1` {
				t.Errorf("expected error in `generator/test-string-errors/template.t1`, but got %v", templateErr.FileName)
			}
			if templateErr.Line != 17 {
				t.Errorf("expected error on line 17, but got %v", templateErr.Line)
			}
			if templateErr.Col != 26 {
				t.Errorf("expected error on column 26, but got %v", templateErr.Col)
			}
		})

		t.Run("the underlying error can be unwrapped", func(t *testing.T) {
			if !errors.Is(err, errSomethingBad) {
				t.Errorf("expected error: %v, but got %v", errSomethingBad, err)
			}
		})

	})
}
