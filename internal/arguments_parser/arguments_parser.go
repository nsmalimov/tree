package arguments_parser

import (
	"os"
)

type ArgumentsParser struct {
}

func New() *ArgumentsParser {
	return &ArgumentsParser{}
}

func (a *ArgumentsParser) Parse(args []string) (targetDir string,
	onlyFolders bool,
	targetIsCurrentDir bool,
	err error) {

	targetDir, err = os.Getwd()
	if err != nil {
		return
	}

	targetIsCurrentDir = true

	if len(args) == 1 {
		targetIsCurrentDir = true
		return
	} else if len(args) > 1 && len(args) <= 3 {
		for _, arg := range args[1:] {
			if arg == "-d" {
				onlyFolders = true
			} else {
				targetIsCurrentDir = false
				targetDir = arg
			}
		}
	}

	return
}
