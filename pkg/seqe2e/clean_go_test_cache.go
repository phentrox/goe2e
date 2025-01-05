package seqe2e

import (
	"log"
	"os/exec"
)

func CleanGoTestCache() {
	cmd := exec.Command("go", "clean", "-testcache")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error cleaning go test cache: %v", err)
	}
}
