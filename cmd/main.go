package main

import (
	"os"

	a "tree/internal/arguments_parser"
	w "tree/internal/walker"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	arguments_parser := a.New()
	targetDir, onlyFolders, targetIsCurrentDir, err := arguments_parser.Parse(os.Args)

	if err != nil {
		logrus.Fatalf("Error when try arguments_parser.Parse, err: %s", err)
	}

	walker := w.New(
		logger,
		targetDir,
		onlyFolders,
		targetIsCurrentDir,
	)

	walker.StartSync()
	walker.Print()
}
