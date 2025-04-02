package cmd

import (
	"github.com/phentrox/goe2e/internal/useCases"
	"github.com/phentrox/goe2e/internal/useCases/collect_test_dirs"
)

func Run(testDir string) error {
	testDirs := collect_test_dirs.CollectTestDirs(testDir)
	useCases.CleanGoTestCache()
	err := useCases.RunTimedE2ETests(testDirs)
	return err
}
