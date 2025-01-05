package collect_test_dirs

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CollectTestDirs(root string) []string {
	var testDirs []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// is file
		if !info.IsDir() {
			if strings.HasSuffix(info.Name(), "test.go") {
				dir := filepath.Dir(path)
				testDirs = append(testDirs, dir)
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return testDirs
}
