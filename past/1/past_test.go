package past_test

import (
	"testing"
	"time"

	"github.com/bitfield/past"

	"github.com/stretchr/testify/assert"
)

func TestOneHourAgo_ReturnsExpectedTime(t *testing.T) {
	t.Parallel()
	now := time.Now()
	then := past.OneHourAgo()
	if now.Hour() == 0 {
		assert.Equal(t, 23, then.Hour())
		assert.Equal(t, now.Day()-1, then.Day())
	} else {
		assert.Equal(t, now.Hour()-1, then.Hour())
		assert.Equal(t, now.Day(), then.Day())
	}
	assert.Equal(t, now.Month(), then.Month())
	assert.Equal(t, now.Year(), then.Year())
}
