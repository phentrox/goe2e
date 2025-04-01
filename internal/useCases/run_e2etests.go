package useCases

import (
	"time"
)

func RunE2ETests(testDirs []string) error {
	for _, testDir := range testDirs {
		err := RunGoTest(testDir)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
