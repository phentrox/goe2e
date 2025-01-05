package seqe2e

import (
	"log"
	"time"
)

func RunE2ETests(testDirs []string) {
	for _, testDir := range testDirs {
		err := RunGoTest(testDir)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}
}
