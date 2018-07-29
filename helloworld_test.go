package gomodules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HelloWorld(t *testing.T) {
	tests := map[string]struct {
		expectedResponse string
	}{
		"success": {
			expectedResponse: "Hello, world",
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		response := sayHello()
		assert.Equal(t, test.expectedResponse, response)
	}
}
