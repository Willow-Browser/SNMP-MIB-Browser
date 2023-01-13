package mibreader

import (
	"fmt"
	"os"
	"strings"

	"github.com/sleepinggenius2/gosmi/parser"
	"github.com/sleepinggenius2/gosmi/types"
	"github.com/willowbrowser/snmpmibbrowser/internal/oidstorage"
	"github.com/willowbrowser/snmpmibbrowser/internal/utils"
)

type MibReader struct {
	types      []parser.Type
	loadedOids *oidstorage.LoadedOids
	newOids    []oidstorage.Oid
}

func NewMibReader(loadedOids *oidstorage.LoadedOids) *MibReader {
	return &MibReader{
		types:      []parser.Type{},
		loadedOids: loadedOids,
		newOids:    []oidstorage.Oid{},
	}
}

func (m *MibReader) ReadMib(fileName string) {
	// new import struct variable is kinda screwing with this
	module := m.parseMibFile(fileName)

	// load in stuff from already involved OIDs and types for the parsing process

	newImports := m.readNewImports(module)

	for _, newImport := range newImports {
		alreadyImported := false
		for _, existingImport := range m.loadedOids.GetLoadedMibs() {
			if newImport.Module.String() == existingImport {
				alreadyImported = true
				break
			}
		}

		if !alreadyImported {
			mibFileExtensions := []string{"", ".txt", ".mib"}
			pathToMibs := m.getBasePathOfMib(fileName)
			var newImportPath string
			fileExists := false

			for _, extension := range mibFileExtensions {
				newImportPath = pathToMibs + newImport.Module.String() + extension
				if utils.FileExists(newImportPath) {
					fileExists = true
					break
				}
			}

			if fileExists {
				m.ReadMib(newImportPath)
			} else {
				fmt.Printf("WARNING: following mib was not in path: %s\n", newImport.Module.String())
				// TODO : return error to frontend
			}
		}
	}

	// need to import the new stuff before going through types
	m.readNewTypes(module)

	if module.Body.Identity != nil {
		m.addIdentity(module.Body.Identity, module.Name.String())
	}
	m.readNewOids(module)

	m.loadedOids.AddNewOids(m.newOids)

	// TODO : if the imports are not located in the same folder as the mib, need to throw an error
	// by emitting an event for Vue to deal with

	fmt.Println("Did it")
}

func (m MibReader) getBasePathOfMib(filePath string) string {
	index := strings.LastIndex(filePath, string(os.PathSeparator))
	return filePath[0:(index + 1)]
}

func (m MibReader) parseMibFile(filePath string) *parser.Module {
	module, err := parser.ParseFile(filePath)
	if err != nil {
		fmt.Errorf("Error parsing mib: %v", err)
	}

	return module
}

func (m *MibReader) readNewImports(module *parser.Module) []parser.Import {
	newImports := []parser.Import{}

	// need information on what mibs we have loaded in, just in case we already have the proper information
	// worry about this last. Don't even know how we are going to store it yet
	for _, moduleImport := range module.Body.Imports {
		duplicateModule := false
		if moduleImport.Module.String() == "SNMPv2-CONF" {
			duplicateModule = true
		} else {
			for _, mib := range m.loadedOids.GetLoadedMibs() {
				if mib == moduleImport.Module.String() {
					duplicateModule = true
					break
				}
			}
		}

		if !duplicateModule {
			newImports = append(newImports, moduleImport)
		}
	}

	return newImports
}

func (m *MibReader) readNewTypes(module *parser.Module) {
	for _, moduleType := range module.Body.Types {
		duplicateType := false
		for i := range m.types {
			if m.types[i].Name == moduleType.Name {
				duplicateType = true
			}
		}

		// probably have to some recursive search for all imports to make sure we have them all

		if !duplicateType {
			m.types = append(m.types, moduleType)
		}
	}

	for _, newType := range m.types {
		fmt.Printf("New type is %s\n", newType.Name.String())
	}
}

func (m *MibReader) addIdentity(moduleIdentity *parser.ModuleIdentity, mibName string) {
	parentOid := m.loadedOids.FindDirectParent(moduleIdentity.Oid.SubIdentifiers[0].Name.String())
	oidNum := appendOidNumber(parentOid.OID, *moduleIdentity.Oid.SubIdentifiers[1].Number)

	newStoredOid := oidstorage.CreateNewOid(moduleIdentity.Name.String(), oidNum, mibName)
	newStoredOid.Description = moduleIdentity.Description
	m.newOids = append(m.newOids, newStoredOid)
	parentOid.AddChildren(&newStoredOid)
}

func (m *MibReader) readNewOids(module *parser.Module) {
	// TODO : fix this up
	mibName := module.Name.String()
	newStoredOid := oidstorage.CreateNewOid("", "", "")

	for _, newOid := range module.Body.Nodes {
		parentName := newOid.Oid.SubIdentifiers[0].Name.String()
		parentOid := m.loadedOids.FindDirectParent(parentName)

		if parentOid == nil {
			parentOid = m.findParentInNewOids(parentName)

			if parentOid == nil {
				fmt.Printf("No suitable parent found for oid: %s\n", newOid.Name.String())
			}
		}

		if newOid.ObjectIdentifier {
			oidNum := appendOidNumber(parentOid.OID, *newOid.Oid.SubIdentifiers[1].Number)
			newStoredOid.Name = newOid.Name.String()
			newOidStore := oidstorage.CreateNewOid(newOid.Name.String(), oidNum, mibName)
			m.newOids = append(m.newOids, newOidStore)

		} else if newOid.ObjectIdentity != nil {
			oidNum := appendOidNumber(parentOid.OID, *newOid.Oid.SubIdentifiers[1].Number)
			newOidStore := oidstorage.CreateNewOid(newOid.Name.String(), oidNum, mibName)
			newOidStore.Description = newOid.ObjectIdentity.Description
			newOidStore.Status = newOid.ObjectIdentity.Status.ToSmi().String()
			newOidStore.Type = oidstorage.ObjectIdentity
			parentOid.AddChildren(&newOidStore)
			m.newOids = append(m.newOids, newOidStore)
		} else if newOid.ObjectType != nil {
			oidNum := appendOidNumber(parentOid.OID, *newOid.Oid.SubIdentifiers[1].Number)
			newOidStore := oidstorage.CreateNewOid(newOid.Name.String(), oidNum, mibName)
			newOidStore.Description = newOid.ObjectType.Description
			newOidStore.Status = newOid.ObjectType.Status.ToSmi().String()
			newOidStore.Type = oidstorage.ObjectType
			newOidStore.Access = newOid.ObjectType.Access.ToSmi().String()
			parentOid.AddChildren(&newOidStore)
			m.newOids = append(m.newOids, newOidStore)
		} else if newOid.ModuleCompliance != nil {
			oidNum := appendOidNumber(parentOid.OID, *newOid.Oid.SubIdentifiers[1].Number)
			newOidStore := oidstorage.CreateNewOid(newOid.Name.String(), oidNum, mibName)
			newOidStore.Description = newOid.ModuleCompliance.Description
			newOidStore.Status = newOid.ModuleCompliance.Status.ToSmi().String()
			newOidStore.Type = oidstorage.ModuleCompliance
			parentOid.AddChildren(&newOidStore)
			m.newOids = append(m.newOids, newOidStore)
		} else if newOid.ObjectGroup != nil {
			oidNum := appendOidNumber(parentOid.OID, *newOid.Oid.SubIdentifiers[1].Number)
			newOidStore := oidstorage.CreateNewOid(newOid.Name.String(), oidNum, mibName)
			newOidStore.Description = newOid.ObjectGroup.Description
			newOidStore.Status = newOid.ObjectGroup.Status.ToSmi().String()
			newOidStore.Type = oidstorage.ObjectGroup
			parentOid.AddChildren(&newOidStore)
			m.newOids = append(m.newOids, newOidStore)
		} else if newOid.NotificationType != nil {
			oidNum := appendOidNumber(parentOid.OID, *newOid.Oid.SubIdentifiers[1].Number)
			newOidStore := oidstorage.CreateNewOid(newOid.Name.String(), oidNum, mibName)
			newOidStore.Description = newOid.NotificationType.Description
			newOidStore.Status = newOid.NotificationType.Status.ToSmi().String()
			newOidStore.Type = oidstorage.NotificationType
			parentOid.AddChildren(&newOidStore)
			m.newOids = append(m.newOids, newOidStore)
		} else if newOid.NotificationGroup != nil {
			oidNum := appendOidNumber(parentOid.OID, *newOid.Oid.SubIdentifiers[1].Number)
			newOidStore := oidstorage.CreateNewOid(newOid.Name.String(), oidNum, mibName)
			newOidStore.Description = newOid.NotificationGroup.Description
			newOidStore.Status = newOid.NotificationGroup.Status.ToSmi().String()
			newOidStore.Type = oidstorage.NotificationGroup
			parentOid.AddChildren(&newOidStore)
			m.newOids = append(m.newOids, newOidStore)
		} else {
			fmt.Printf("Still need some work to properly add oid: %s\n", newOid.Name.String())
		}
	}
}

func (m *MibReader) findParentInNewOids(parentName string) *oidstorage.Oid {
	var parentOid *oidstorage.Oid

	for _, oid := range m.newOids {
		if oid.Name == parentName {
			parentOid = &oid
			break
		}
	}

	return parentOid
}

// TODO : some way to overwrite duplicates. Assuming the mib is an updated mib

func appendOidNumber(oidNumStr string, newOidNum types.SmiSubId) string {
	var sb strings.Builder

	sb.WriteString(oidNumStr)
	sb.WriteString(fmt.Sprintf(".%d", newOidNum))

	return sb.String()
}
