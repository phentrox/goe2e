package seqe2e

import "github.com/phentrox/go-seq-e2e-testing/pkg/seqe2e/collect_test_dirs"

func Run() {
	testDirs := collect_test_dirs.CollectTestDirs("e2eTesting")
	CleanGoTestCache()
	RunTimedE2ETest(testDirs)
}
