package useCases

import (
	"fmt"
	"time"
)

func RunTimedE2ETests(testDirs []string) error {
	println("### Testing End To End ###")

	start := time.Now()

	err := RunE2ETests(testDirs)
	if err != nil {
		return err
	}

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Println("### Testing Time: ", elapsed.Round(time.Second).String(), " ###")
	return nil
}
