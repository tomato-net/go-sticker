package alphavantage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsError(t *testing.T) {
	t.Parallel()

	t.Run("returns true if given error", func(t *testing.T) {
		giveError := ErrorMessage{Message: "error"}
		gotBool := IsError(giveError)
		assert.True(t, gotBool)
	})

	t.Run("returns false if not given error", func(t *testing.T) {
		giveError := ErrorMessage{}
		gotBool := IsError(giveError)
		assert.False(t, gotBool)
	})
}
