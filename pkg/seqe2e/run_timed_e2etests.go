package seqe2e

import (
	"time"
)

func RunTimedE2ETest(testDirs []string) {
	println("### Testing End To End ###")

	start := time.Now()

	RunE2ETests(testDirs)

	end := time.Now()
	elapsed := end.Sub(start)

	println("### Testing Time: %v ###", elapsed)
}
