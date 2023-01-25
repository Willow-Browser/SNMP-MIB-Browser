package mibreader

import (
	"fmt"
	"os"
	"testing"

	"github.com/sleepinggenius2/gosmi/parser"
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
	m := NewMibReader(oidstorage.NewLoadedOids(db))
	m.ReadMib(fmt.Sprintf("..%s..%stest%sCISCO-QOS-PIB-MIB.txt", string(os.PathSeparator), string(os.PathSeparator), string(os.PathSeparator)))
}

func TestScalarOids(t *testing.T) {
	subTests := []struct {
		name        string
		oidName     string
		eStatus     string
		status      parser.Status
		description string
		eType       oidstorage.OidType
		eAccess     string
		access      parser.Access
		oid         string
		mib         string
		parentOid   string
		subId       types.SmiSubId
		syntax      string
	}{
		{
			name:        "First Test",
			oidName:     "testOid",
			eStatus:     "Current",
			status:      parser.StatusCurrent,
			description: "test description",
			eAccess:     "ReadOnly",
			access:      parser.AccessReadOnly,
			oid:         ".1.3.6",
			parentOid:   ".1.3",
			subId:       6,
			mib:         "TestMib",
			syntax:      "Unsigned32",
		},
		{
			name:        "Read-Write Oid Test",
			oidName:     "testOid3",
			eStatus:     "Current",
			status:      parser.StatusCurrent,
			description: "test description",
			eAccess:     "ReadWrite",
			access:      parser.AccessReadWrite,
			oid:         ".1.3.6.5.1.7",
			parentOid:   ".1.3.6.5.1",
			subId:       7,
			mib:         "TestMib",
			syntax:      "Unsigned32",
		},
	}

	for _, subtest := range subTests {
		t.Run(subtest.name, func(t *testing.T) {
			l := oidstorage.NewLoadedOids(db)
			m := NewMibReader(l)

			expected := oidstorage.CreateNewOid(subtest.oidName, subtest.oid, subtest.mib)
			expected.Description = subtest.description
			expected.Status = subtest.eStatus
			expected.Type = oidstorage.ObjectType
			expected.Access = subtest.eAccess
			expected.Syntax = subtest.syntax

			parentOid := oidstorage.Oid{
				OID: subtest.parentOid,
			}

			identifier := types.SmiIdentifier(subtest.oidName)

			testNode := parser.Node{
				Oid: &parser.Oid{
					SubIdentifiers: []parser.SubIdentifier{
						{
							Name: &identifier,
						},
						{
							Number: &subtest.subId,
						},
					},
				},
				Name: identifier,
				ObjectType: &parser.ObjectType{
					Description: subtest.description,
					Status:      subtest.status,
					Access:      subtest.access,
					Syntax: parser.Syntax{
						Type: &parser.SyntaxType{
							Name: types.SmiIdentifier(subtest.syntax),
						},
					},
				},
			}

			m.parseObjectType(&testNode, &parentOid, subtest.mib)

			l.AddNewOids(m.newOids)

			oids := l.GetLoadedOids()

			assert.Equal(t, expected, oids[len(oids)-1])
		})
	}

}

func TestParseTableOid(t *testing.T) {
	subTests := []struct {
		name        string
		oidName     string
		eStatus     string
		status      parser.Status
		description string
		eType       oidstorage.OidType
		eAccess     string
		access      parser.Access
		oid         string
		mib         string
		parentOid   string
		subId       types.SmiSubId
		sequence    string
	}{
		{
			name:        "First Test",
			oidName:     "testOidTable",
			eStatus:     "Current",
			status:      parser.StatusCurrent,
			description: "test description",
			eAccess:     "NotAccessible",
			access:      parser.AccessNotAccessible,
			oid:         ".1.3.6",
			parentOid:   ".1.3",
			subId:       6,
			mib:         "TestMib",
			sequence:    "testOidEntry",
		},
	}

	for _, subtest := range subTests {
		t.Run(subtest.name, func(t *testing.T) {
			l := oidstorage.NewLoadedOids(db)
			m := NewMibReader(l)

			expected := oidstorage.CreateNewOid(subtest.oidName, subtest.oid, subtest.mib)
			expected.Description = subtest.description
			expected.Status = subtest.eStatus
			expected.Type = oidstorage.ObjectType
			expected.Access = subtest.eAccess
			expected.Syntax = fmt.Sprintf("SEQUENCE OF %s", subtest.sequence)

			parentOid := oidstorage.Oid{
				OID: subtest.parentOid,
			}

			identifier := types.SmiIdentifier(subtest.oidName)

			testNode := parser.Node{
				Oid: &parser.Oid{
					SubIdentifiers: []parser.SubIdentifier{
						{
							Name: &identifier,
						},
						{
							Number: &subtest.subId,
						},
					},
				},
				Name: identifier,
				ObjectType: &parser.ObjectType{
					Description: subtest.description,
					Status:      subtest.status,
					Access:      subtest.access,
					Syntax: parser.Syntax{
						Sequence: (*types.SmiIdentifier)(&subtest.sequence),
					},
				},
			}

			m.parseObjectType(&testNode, &parentOid, subtest.mib)

			l.AddNewOids(m.newOids)

			oids := l.GetLoadedOids()

			assert.Equal(t, expected, oids[len(oids)-1])
		})
	}

}
