package country

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCountry(t *testing.T) {
	tests := []struct {
		Expected bool
		Value    string
	}{
		{false, "FRE"},
		{true, "ZMB"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.Expected, HasCountry(tt.Value))
	}

}
