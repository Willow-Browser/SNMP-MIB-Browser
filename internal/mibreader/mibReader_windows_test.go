//go:build windows
// +build windows

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
		{"Windows", "C:\\Users\\test user\\AppData\\Local\\Mibs\\SNMP.txt", "C:\\Users\\test user\\AppData\\Local\\Mibs\\"},
	}

	for _, parmeter := range parameters {
		t.Run(fmt.Sprintf("Testing: %s", parmeter.testName), func(t *testing.T) {
			l := oidstorage.NewLoadedOids(db)
			mibReader := NewMibReader(l)
			actual := mibReader.getBasePathOfMib(parmeter.filePath)
			assert.Equal(t, parmeter.expected, actual)
		})
	}
}
