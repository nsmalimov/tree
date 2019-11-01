package test

import (
	"fmt"
	"os"
	"testing"

	walker2 "tree/internal/walker"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// todo: создавать и удалять тестовые файлы из кода

type testCase struct {
	path               string
	res                []string
	onlyFolders        bool
	targetIsCurrentDir bool
}

func TestWalk(t *testing.T) {
	logger := logrus.New()

	currentDir, err := os.Getwd()

	if err != nil {
		t.Fatalf("Error when try os.Getwd, err: %s", err)
	}

	testCases := []testCase{
		{
			path: fmt.Sprintf("%s/%s", currentDir, "/test_folders/1/www/theme"),
			res: []string{
				"├── classic",
				"│   └── views",
				"└── kiosk",
				"    ├── css",
				"    │   └── main.css",
				"    └── images",
				"        ├── android.png",
				"        ├── bg_share.jpg",
				"        ├── bg_share.png",
				"        ├── download.png",
				"        ├── ios.png",
				"        ├── share_fb.png",
				"        ├── share_tv.png",
				"        └── share_vk.png",
			},
			targetIsCurrentDir: true,
		},
		{
			path: fmt.Sprintf("%s/%s", currentDir, "/test_folders/2/pac"),
			res: []string{
				"├── lib",
				"│   ├── edit",
				"│   ├── ex",
				"│   │   ├── vte32",
				"│   │   │   └── auto",
				"│   │   │       └── Gnome2",
				"│   │   │           └── Vte",
				"│   │   ├── vte64",
				"│   │   │   └── auto",
				"│   │   │       ├── Gnome2",
				"│   │   │       │   └── Vte",
				"│   │   ├── vteARM",
				"│   │   │   └── auto",
				"│   │   │       └── Gnome2",
				"│   │   │           └── Vte",
				"│   │   ├── vteARMV7L",
				"│   │   │   └── auto",
				"│   │   │       └── Gnome2",
				"│   │   │           ├── Vte",
				"│   │   └── vtePPC64",
				"│   │       └── auto",
				"│   │           └── Gnome2",
				"│   │               └── Vte",
				"│   └── method",
				"├── res",
			},
			onlyFolders: true,
		},
	}

	for _, testCase := range testCases {
		walker := walker2.New(
			logger,
			testCase.path,
			testCase.onlyFolders,
			testCase.targetIsCurrentDir,
		)

		walker.StartSync()

		assert.True(t, len(walker.Result) == (len(testCase.res)+1))
		assert.Equal(t, testCase.res, walker.Result[1:])
	}
}
