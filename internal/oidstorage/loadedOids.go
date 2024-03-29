package oidstorage

type LoadedOids struct {
	oids       []Oid
	loadedMibs []string
	db         *DB
}

func NewLoadedOids(db *DB) *LoadedOids {
	// TODO : load oids from data cache

	loadedMibs := []string{"SNMPv2-SMI"}

	return &LoadedOids{
		oids:       createBaseOids(),
		loadedMibs: loadedMibs,
		db:         db,
	}
}

func createBaseOids() []Oid {
	baseOids := []Oid{}

	iso := CreateNewOid("iso", ".1", "SNMPv2-SMI")
	iso.Type = ModuleIdentity

	org := CreateNewOid("org", ".1.3", "SNMPv2-SMI")
	org.Type = ObjectIdentity

	dod := CreateNewOid("dod", ".1.3.6", "SNMPv2-SMI")
	dod.Type = ObjectIdentity

	internet := CreateNewOid("internet", ".1.3.6.1", "SNMPv2-SMI")
	internet.Type = ObjectIdentity

	directory := CreateNewOid("directory", ".1.3.6.1.1", "SNMPv2-SMI")
	directory.Type = ObjectIdentity

	mgmt := CreateNewOid("mgmt", ".1.3.6.1.2", "SNMPv2-SMI")
	mgmt.Type = ObjectIdentity

	mib_2 := CreateNewOid("mib-2", ".1.3.6.1.2.1", "SNMPv2-SMI")
	mib_2.Type = ObjectIdentity

	system := CreateNewOid("system", ".1.3.6.1.2.1.1", "SNMPv2-MIB")
	system.Type = ObjectIdentity

	sysDescr := CreateNewOid("sysDescr", ".1.3.6.1.2.1.1.1", "SNMPv2-MIB")
	sysDescr.Access = "read-only"
	sysDescr.Type = ObjectType

	sysObjectID := CreateNewOid("sysObjectID", ".1.3.6.1.2.1.1.2", "SNMPv2-MIB")
	sysObjectID.Access = "read-only"
	sysObjectID.Type = ObjectType

	sysUpTime := CreateNewOid("sysUpTime", ".1.3.6.1.2.1.1.3", "SNMPv2-MIB")
	sysUpTime.Access = "read-only"
	sysUpTime.Type = ObjectType

	sysContact := CreateNewOid("sysContact", ".1.3.6.1.2.1.1.4", "SNMPv2-MIB")
	sysContact.Access = "read-write"
	sysContact.Type = ObjectType

	sysName := CreateNewOid("sysName", ".1.3.6.1.2.1.1.5", "SNMPv2-MIB")
	sysName.Access = "read-write"
	sysName.Type = ObjectType

	sysLocation := CreateNewOid("sysLocation", ".1.3.6.1.2.1.1.6", "SNMPv2-MIB")
	sysLocation.Access = "read-write"
	sysLocation.Type = ObjectType

	sysServices := CreateNewOid("sysServices", ".1.3.6.1.2.1.1.7", "SNMPv2-MIB")
	sysServices.Access = "read-only"
	sysServices.Type = ObjectType

	experimental := CreateNewOid("experimental", ".1.3.6.1.3", "SNMPv2-SMI")
	experimental.Type = ObjectIdentity

	private := CreateNewOid("private", ".1.3.6.1.4", "SNMPv2-SMI")
	private.Type = ObjectIdentity

	enterprises := CreateNewOid("enterprises", ".1.3.6.1.4.1", "SNMPv2-SMI")
	enterprises.Type = ObjectIdentity

	security := CreateNewOid("security", ".1.3.6.1.5", "SNMPv2-SMI")
	security.Type = ObjectIdentity

	snmpV2 := CreateNewOid("snmpV2", ".1.3.6.1.6", "SNMPv2-SMI")
	snmpV2.Type = ObjectIdentity

	snmpDomains := CreateNewOid("snmpDomains", ".1.3.6.1.6.1", "SNMPv2-SMI")
	snmpDomains.Type = ObjectIdentity

	snmpProxys := CreateNewOid("snmpProxys", ".1.3.6.1.6.2", "SNMPv2-SMI")
	snmpProxys.Type = ObjectIdentity

	snmpModules := CreateNewOid("snmpModules", ".1.3.6.1.6.3", "SNMPv2-SMI")
	snmpModules.Type = ObjectIdentity

	iso.AddChildren(&org)
	org.AddChildren(&dod)
	dod.AddChildren(&internet)
	snmpV2.AddChildren(&snmpDomains, &snmpProxys, &snmpModules)
	private.AddChildren(&enterprises)
	internet.AddChildren(&directory, &mgmt, &experimental, &private, &security, &snmpV2)
	mgmt.AddChildren(&mib_2)
	mib_2.AddChildren(&system)
	system.AddChildren(&sysDescr, &sysObjectID, &sysUpTime, &sysContact, &sysName, &sysLocation, &sysServices)

	appendOid(&baseOids, iso)
	appendOid(&baseOids, org)
	appendOid(&baseOids, dod)
	appendOid(&baseOids, internet)
	appendOid(&baseOids, directory)
	appendOid(&baseOids, mgmt)
	appendOid(&baseOids, mib_2)
	appendOid(&baseOids, system)
	appendOid(&baseOids, sysDescr)
	appendOid(&baseOids, sysObjectID)
	appendOid(&baseOids, sysUpTime)
	appendOid(&baseOids, sysContact)
	appendOid(&baseOids, sysName)
	appendOid(&baseOids, sysLocation)
	appendOid(&baseOids, sysServices)
	appendOid(&baseOids, experimental)
	appendOid(&baseOids, private)
	appendOid(&baseOids, enterprises)
	appendOid(&baseOids, security)
	appendOid(&baseOids, snmpV2)
	appendOid(&baseOids, snmpDomains)
	appendOid(&baseOids, snmpProxys)
	appendOid(&baseOids, snmpModules)

	return baseOids
}

func appendOid(oidSlice *[]Oid, newOid Oid) {
	*oidSlice = append(*oidSlice, newOid)
}

func (l *LoadedOids) FindDirectParent(parentName string) *Oid {
	var parentOid *Oid

	for _, oid := range l.oids {
		if oid.Name == parentName {
			parentOid = &oid
			break
		}
	}

	// TODO : should really return an error if the parent is not found

	return parentOid
}

func (l *LoadedOids) AddNewOids(newOids []Oid) {
	newMibs := []string{}
	newLoadedOids := []Oid{}
	overwriteOids := []Oid{}

	for _, oid := range newOids {
		foundMib := false
		for _, mib := range l.loadedMibs {
			if mib == oid.Mib {
				foundMib = true
				break
			}
		}

		if !foundMib {
			newMibs = append(newMibs, oid.Mib)
		}

		foundOid := false
		for _, oldOids := range l.oids {
			if (oldOids.Name == oid.Name) && (oldOids.OID == oid.OID) {
				overwriteOids = append(overwriteOids, oid)
				foundOid = true
				break
			}
		}

		if !foundOid {
			newLoadedOids = append(newLoadedOids, oid)
		}
	}

	l.oids = append(l.oids, newLoadedOids...)
	l.loadedMibs = append(l.loadedMibs, newMibs...)
}

func (l *LoadedOids) MarkIndexOids(newIndices []string) {
	for _, index := range newIndices {
		oid := l.findOidByName(index)
		if oid != nil {
			oid.IsIndex = true
		}
	}
}

func (l *LoadedOids) findOidByName(oidName string) *Oid {
	for i, oid := range l.oids {
		if oid.Name == oidName {
			return &l.oids[i] // have to do this weird hack since we range won't give us a reference
		}
	}
	return nil
}

func (l *LoadedOids) GetLoadedOids() []Oid {
	return l.oids
}

func (l *LoadedOids) GetLoadedMibs() []string {
	return l.loadedMibs
}
