package seqe2e

import (
	"fmt"
	"time"
)

func RunTimedE2ETest(testDirs []string) {
	println("### Testing End To End ###")

	start := time.Now()

	RunE2ETests(testDirs)

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Println("### Testing Time: ", elapsed.Round(time.Second).String(), " ###")
}
