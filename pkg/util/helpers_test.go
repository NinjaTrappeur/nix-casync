package util

import (
	"testing"

	"github.com/numtide/go-nix/nixbase32"
	"github.com/stretchr/testify/assert"
)

func TestStorePathHelpers(t *testing.T) {
	storePath := "/nix/store/dr76fsw7d6ws3pymafx0w0sn4rzbw7c9-etc-os-release"

	t.Run("getHashFromStorePath", func(t *testing.T) {
		hash, err := GetHashFromStorePath(storePath)
		assert.NoError(t, err)
		assert.Equal(t, "dr76fsw7d6ws3pymafx0w0sn4rzbw7c9", nixbase32.EncodeToString(hash))
	})

	t.Run("getNameFromStorePath", func(t *testing.T) {
		name := GetNameFromStorePath(storePath)
		assert.Equal(t, "etc-os-release", name)
	})
}
