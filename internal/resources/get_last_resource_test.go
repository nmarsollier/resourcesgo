package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		version string
		semVer  string
		valid   bool
	}{
		{"1.2.3", "1.2.3", true},
		{"1.2.3", "1.2.4", true},
		{"1.2.3", "1.2.*", true},
		{"1.2.3", "1.2.+", true},
		{"1.2.3", "1.3.*", false},
		{"1.2.3", "1.3.+", false},
		{"1.2.3", "2.0.0", false},
		{"1.2.3", "1.*", true},
		{"1.2.3", "2.*", false},
		{"1.2.3", "*", true},
		{"1.2.3", "", false},
		{"1.2.3", "asd", false},
		{"1.2.3", "1.+", true},
		{"1.2.3", "1.2", false},
		{"1.2", "1.2.3", false},
		{"1.2.3", "1.2.3.4", false},
	}

	for _, test := range tests {
		t.Run(test.version+"-"+test.semVer, func(t *testing.T) {
			result := isValidSemver(test.version, test.semVer)
			assert.Equal(t, test.valid, result)
		})
	}
}
