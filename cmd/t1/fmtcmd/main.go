package fmtcmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/natefinch/atomic"
	"github.com/senforsce/t1/cmd/t1/processor"
	parser "github.com/senforsce/t1/parser/v2"
)

const workerCount = 4

func Run(w io.Writer, args []string) (err error) {
	if len(args) > 0 {
		return formatDir(w, args[0])
	}
	return formatReader(w, os.Stdin)
}

func formatReader(w io.Writer, r io.Reader) (err error) {
	var bytes []byte
	bytes, err = io.ReadAll(r)
	if err != nil {
		return
	}
	t, err := parser.ParseString(string(bytes))
	if err != nil {
		return fmt.Errorf("parsing error: %w", err)
	}
	err = t.Write(w)
	if err != nil {
		return fmt.Errorf("formatting error: %w", err)
	}
	return nil
}

func formatDir(w io.Writer, dir string) (err error) {
	start := time.Now()
	results := make(chan processor.Result)
	go processor.Process(dir, format, workerCount, results)
	var successCount, errorCount int
	for r := range results {
		if r.Error != nil {
			err = errors.Join(err, fmt.Errorf("%s: %w", r.FileName, r.Error))
			errorCount++
			continue
		}
		fmt.Printf("%s complete in %v\n", r.FileName, r.Duration)
		successCount++
	}
	fmt.Printf("Formatted %d templates with %d errors in %s\n", successCount+errorCount, errorCount, time.Since(start))
	return
}

func format(fileName string) (err error) {
	contents, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("failed to read file %q: %w", fileName, err)
	}
	t, err := parser.ParseString(string(contents))
	if err != nil {
		return fmt.Errorf("%s parsing error: %w", fileName, err)
	}
	w := new(bytes.Buffer)
	err = t.Write(w)
	if err != nil {
		return fmt.Errorf("%s formatting error: %w", fileName, err)
	}
	if string(contents) == w.String() {
		return nil
	}
	err = atomic.WriteFile(fileName, w)
	if err != nil {
		return fmt.Errorf("%s file write error: %w", fileName, err)
	}
	return
}
