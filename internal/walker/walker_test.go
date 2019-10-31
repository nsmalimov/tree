package walker

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func prepareTempDirWithFiles() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	testDirName := "test_dir"

	testDirPath, err := ioutil.TempDir(fmt.Sprintf("%s", currentDir), testDirName)

	//tmpfile, err := ioutil.TempFile("", "example")
	//if err != nil {
	//	log.Fatal(err)
	//}

	return testDirPath, err
}

func deleteTempDirAndFiles() {
	//	defer os.RemoveAll(dir) // clean up
}

func TestWalk(t *testing.T) {
	logger := logrus.New()
	targetDir := ""
	onlyFolders := false

	walker := New(
		logger,
		targetDir,
		onlyFolders,
	)

	// walker.Start()

	mainTestDir, err := prepareTempDirWithFiles()

	if err != nil {
		t.Fatalf("Error when try prepareTempDirWithFiles, err: %s", err)
	}

	fmt.Println(walker, mainTestDir)
	assert.Equal(t, 1, 1)
}
