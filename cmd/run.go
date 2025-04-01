package cmd

import (
	"github.com/phentrox/goseq/internal/useCases"
	"github.com/phentrox/goseq/internal/useCases/collect_test_dirs"
)

func Run(testDir string) error {
	testDirs := collect_test_dirs.CollectTestDirs(testDir)
	useCases.CleanGoTestCache()
	err := useCases.RunTimedE2ETests(testDirs)
	return err
}
