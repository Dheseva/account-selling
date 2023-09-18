package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableWorld(t *testing.T){
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "World",
			request: "World",
			expected: "Hello World",
		},
		{
			name: "Sekai",
			request: "Sekai",
			expected: "Hello Sekai",
		},
		{
			name: "Dunia",
			request: "Dunia",
			expected: "Hello Dunia",
		},
	}

	for _, tested := range tests {
		t.Run(tested.name, func(t *testing.T) {
			result := HelloWorld(tested.request)
			require.Equal(t, tested.expected, result)
		})
	}

}
func TestHelloAssert(t *testing.T){
	result := HelloWorld("World")
	assert.Equal(t, "HelloWorld", result, "Result must be Hello World") // fail test unit
	require.Equal(t, "Hello World", result, "Result must be Hello World") // failnow test unit
}