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
	Result      []string
}

func New(logger *logrus.Logger, targetDir string, onlyFolders, targetIsCurrentDir bool) *Walker {
	walker := &Walker{
		logger:      logger,
		targetDir:   targetDir,
		onlyFolders: onlyFolders,
		Result:      make([]string, 0),
	}

	if !targetIsCurrentDir {
		// добавляем заданную директорию для "принта" в начале вывода результата
		walker.Result = append(walker.Result, targetDir)
	}

	return walker
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

func (w *Walker) walkSync(filePath string, prefix string) {
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
			w.Result = append(w.Result, fmt.Sprintf("%s└── %s", prefix, dirname))
			w.walkSync(subFilePath, prefix+"    ")
		} else {
			w.Result = append(w.Result, fmt.Sprintf("%s├── %s", prefix, dirname))
			w.walkSync(subFilePath, prefix+"│   ")
		}
	}
}

func (w *Walker) StartSync() {
	w.walkSync(w.targetDir, "")
}

func (w *Walker) Print() {
	for _, elem := range w.Result {
		fmt.Println(elem)
	}
}
