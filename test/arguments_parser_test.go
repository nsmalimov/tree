package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	a "tree/internal/arguments_parser"
)

type TestCase struct {
	args               []string
	targetDir          string
	onlyFolders        bool
	targetIsCurrentDir bool
}

func TestParseArgs(t *testing.T) {
	targetDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error when try os.Getwd, err: %s", err)
	}

	fmt.Println(targetDir)
	test_cases := []TestCase{
		{
			args: []string{
				"program_name",
			},
			targetDir:          targetDir,
			onlyFolders:        false,
			targetIsCurrentDir: true,
		},
		{
			args: []string{
				"program_name",
				"-d",
			},
			targetDir:          targetDir,
			onlyFolders:        true,
			targetIsCurrentDir: true,
		},
		{
			args: []string{
				"program_name",
				"-d",
				"some_dir",
			},
			targetDir:          "some_dir",
			onlyFolders:        true,
			targetIsCurrentDir: false,
		},
		{
			args: []string{
				"program_name",
				"some_dir",
				"-d",
			},
			targetDir:          "some_dir",
			onlyFolders:        true,
			targetIsCurrentDir: false,
		},
		{
			args: []string{
				"program_name",
				"some_dir",
			},
			targetDir:          "some_dir",
			onlyFolders:        false,
			targetIsCurrentDir: false,
		},
	}

	arguments_parser := a.New()

	for _, test_case := range test_cases {
		targetDir, onlyFolders, targetIsCurrentDir, err := arguments_parser.Parse(test_case.args)

		if err != nil {
			t.Fatalf("Error when try arguments_parser.Parse, err: %s", err)
		}

		assert.Equal(t, test_case.targetDir, targetDir)
		assert.Equal(t, test_case.onlyFolders, onlyFolders)
		assert.Equal(t, test_case.targetIsCurrentDir, targetIsCurrentDir)
	}
}
