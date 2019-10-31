package walker

import (
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/sirupsen/logrus"
)

type Walker struct {
	logger      *logrus.Logger
	targetDir   string
	onlyFolders bool
	res         []string
}

func New(logger *logrus.Logger, targetDir string, onlyFolders bool) *Walker {
	return &Walker{
		logger:      logger,
		targetDir:   targetDir,
		onlyFolders: onlyFolders,
		res:         make([]string, 0),
	}
}

func (w *Walker) dirnamesFromPath(path string) (dirnames []string) {
	file, err := os.Open(path)
	if err != nil {
		w.logger.Fatalf("Error when try os.Open, err: %s", err)
	}

	dirnames, err = file.Readdirnames(0)

	defer func() {
		err = file.Close()

		if err != nil {
			w.logger.Fatalf("Error when try file.Close, err: %s", err)
		}
	}()

	sort.Strings(dirnames)

	return
}

func (w *Walker) walk(filePath string, prefix string) {
	dirnames := w.dirnamesFromPath(filePath)

	for index, dirname := range dirnames {
		// пропускаем скрытые
		if dirname[0] == '.' {
			continue
		}

		subFilePath := path.Join(filePath, dirname)

		stat, err := os.Stat(subFilePath)

		if err != nil {
			logrus.Fatalf("Error when try os.Stat, err: %s", err)
		}

		if !stat.IsDir() && w.onlyFolders {
			continue
		}

		if index == len(dirnames)-1 {
			w.res = append(w.res, fmt.Sprintf("%s└──%s", prefix, dirname))
			w.walk(subFilePath, prefix+"   ")
		} else {
			w.res = append(w.res, fmt.Sprintf("%s├──%s", prefix, dirname))
			w.walk(subFilePath, prefix+"│  ")
		}
	}
}

func (w *Walker) Start() {
	w.walk(w.targetDir, "")
}

func (w *Walker) Print() {
	for _, elem := range w.res {
		fmt.Println(elem)
	}
}
