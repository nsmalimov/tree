package main

import (
	"os"

	w "tree/internal/walker"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	var targetDir string

	var err error

	var onlyFolders bool

	targetDir, err = os.Getwd()
	if err != nil {
		logger.Fatalf("Error when try os.Getwd, err: %s", err)
	}

	if len(os.Args) >= 1 && len(os.Args) < 3 {
		for _, arg := range os.Args[1:] {
			if arg == "-d" {
				onlyFolders = true
			} else {
				targetDir = os.Args[1]
			}
		}
	} else {
		logger.Fatal("Count of given arguments > 2")
	}

	walker := w.New(
		logger,
		targetDir,
		onlyFolders,
	)

	walker.Start()
	walker.Print()
}
