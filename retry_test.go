package nimble

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRetryTerminalError(t *testing.T) {

	e := func() error {
		return fmt.Errorf("Error")
	}

	delay := time.Millisecond * 250
	start := time.Now()
	err := Retry(&RetryConfig{
		NumAttempts:         3,
		DelayBetweenRetries: delay,
	}, e)
	end := time.Now()
	dur := end.Sub(start)
	assert.GreaterOrEqual(t, dur, delay*2)
	assert.Less(t, dur, delay*3)
	assert.NotNil(t, err)
}

func TestRetryNoError(t *testing.T) {

	e := func() error {
		return nil
	}

	delay := time.Millisecond * 250
	start := time.Now()
	err := Retry(&RetryConfig{
		NumAttempts:         3,
		DelayBetweenRetries: delay,
	}, e)
	end := time.Now()
	dur := end.Sub(start)
	assert.Less(t, dur, delay)
	assert.Nil(t, err)
}
