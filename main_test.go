package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_testValidity(t *testing.T) {
	assert.Equal(t, true, testValidity("123-asdf-456-def"))
	assert.Equal(t, true, testValidity("123-asdf-456-def-0089-rty"))
	assert.Equal(t, true, testValidity("123-asdf"))
	assert.Equal(t, true, testValidity("346-random-466-string"))

	assert.Equal(t, false, testValidity("not-valid"))
	assert.Equal(t, false, testValidity("this-123-is-567-not-9875-valid"))
	assert.Equal(t, false, testValidity("this 954 is also not valid"))
	assert.Equal(t, false, testValidity("-not-valid-tooo-32554-"))
	assert.Equal(t, false, testValidity("-67-fedr-333-gh"))
	assert.Equal(t, false, testValidity("67-fedr-333-gh-"))
}
