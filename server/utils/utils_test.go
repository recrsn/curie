package utils

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMust(t *testing.T) {
	t.Run("Should panic if error is not nil", func(t *testing.T) {
		defer func() {
			err := recover()
			assert.NotNil(t, err)
		}()
		Must(errors.New("some error"))
	})

	t.Run("Should not panic if error is nil", func(t *testing.T) {
		defer func() {
			err := recover()
			assert.Nil(t, err)
		}()
		Must(nil)
	})
}
