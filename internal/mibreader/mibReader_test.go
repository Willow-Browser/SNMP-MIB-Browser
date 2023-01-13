package mibreader

import (
	"fmt"
	"os"
	"testing"

	"github.com/sleepinggenius2/gosmi/types"
	"github.com/stretchr/testify/assert"
	"github.com/willowbrowser/snmpmibbrowser/internal/oidstorage"
)

func TestAppendOidNumbers(t *testing.T) {
	parameters := []struct {
		oidNum    string
		newOidNum types.SmiSubId
		expected  string
	}{
		{".1", 3, ".1.3"},
		{".1.3.6.1.3.1.585.4", 6175, ".1.3.6.1.3.1.585.4.6175"},
	}

	for i, parameter := range parameters {
		t.Run(fmt.Sprintf("Testing [%v]", i), func(t *testing.T) {
			actual := appendOidNumber(parameter.oidNum, parameter.newOidNum)
			assert.Equal(t, parameter.expected, actual)
		})
	}
}

func TestReadMibWithStandardImports(t *testing.T) {
	// TODO : actually complete this test
}

func TestReadMibWithMultipleImports(t *testing.T) {
	// TODO : actually complete this test
	m := NewMibReader(oidstorage.NewLoadedOids())
	m.ReadMib(fmt.Sprintf("..%s..%stest%sCISCO-QOS-PIB-MIB.txt", string(os.PathSeparator), string(os.PathSeparator), string(os.PathSeparator)))
}
