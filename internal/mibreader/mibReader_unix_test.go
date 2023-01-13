//go:build !windows
// +build !windows

package mibreader

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willowbrowser/snmpmibbrowser/internal/oidstorage"
)

func TestGetBasePath(t *testing.T) {
	parameters := []struct {
		testName string
		filePath string
		expected string
	}{
		{"Unix", "/home/test user/.mibs/SNMP.txt", "/home/test user/.mibs/"},
	}

	for _, parmeter := range parameters {
		t.Run(fmt.Sprintf("Testing: %s", parmeter.testName), func(t *testing.T) {
			l := oidstorage.NewLoadedOids()
			mibReader := NewMibReader(l)
			actual := mibReader.getBasePathOfMib(parmeter.filePath)
			assert.Equal(t, parmeter.expected, actual)
		})
	}
}
