package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlagValue_parse(t *testing.T) {
	flagVal := flagValue{
		startOffset: 2,
		endOffset:   3,
	}

	a := assert.New(t)
	a.Equal(flagVal.parse(), uint64(100))
}
