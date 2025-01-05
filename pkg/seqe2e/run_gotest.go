package seqe2e

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunGoTest(testDir string) error {
	cmd := exec.Command("go", "test", "./"+testDir)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running test | dir: %v | error: %v | output: %v", testDir, err, string(output))
	} else {
		outputWithoutNewLine, _ := strings.CutSuffix(string(output), "\n")
		println(outputWithoutNewLine)
	}
	return nil
}
