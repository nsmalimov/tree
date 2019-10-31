package test

import (
	"fmt"
	"os"
	"testing"

	w "tree/internal/walker"

	"github.com/sirupsen/logrus"
)

// mac os, 4 cors, 8gb, sync processing
// BenchmarkSample-4           3135            384393 ns/op           28767 B/op        212 allocs/op

func BenchmarkSample(b *testing.B) {
	logger := logrus.New()

	currentDir, err := os.Getwd()

	if err != nil {
		b.Fatalf("Error when try os.Getwd, err: %s", err)
		return
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		walker := w.New(
			logger,
			fmt.Sprintf("%s/%s", currentDir, "/test_folders/1/www/theme"),
			false,
			false,
		)

		walker.StartSync()
	}
}
