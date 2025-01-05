package collect_test_dirs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectTestDirs(t *testing.T) {
	// given
	root := "test-data"

	// when
	actual := CollectTestDirs(root)

	// then expected
	expected := []string{
		"test-data/data/cleanup/cleanupAll",
		"test-data/data/cleanup/cleanupEmpty",
		"test-data/truncation/truncationMultipleSchemasTest",
	}

	// then assert
	assert.Equal(t, expected, actual)
}
