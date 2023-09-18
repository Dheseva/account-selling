package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloAssert(t *testing.T){
	result := HelloWorld("World")
	assert.Equal(t, "HelloWorld", result, "Result must be Hello World")
}