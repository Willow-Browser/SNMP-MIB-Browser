package oidstorage

type OidType string
type OidObjectType string

const (
	ObjectIdentity    OidType = "ObjectIdentity"
	ModuleIdentity    OidType = "ModuleIdentity"
	ObjectType        OidType = "ObjectType"
	ModuleCompliance  OidType = "ModuleCompliance"
	ObjectGroup       OidType = "ObjectGroup"
	NotificationType  OidType = "NotificationType"
	NotificationGroup OidType = "NotificationGroup"
)

const (
	Scalar OidObjectType = "Scalar"
	Table  OidObjectType = "Table"
	Entry  OidObjectType = "Entry"
)

type Oid struct {
	Name        string   `json:"name"`
	OID         string   `json:"oid"`
	Mib         string   `json:"mib"`
	Syntax      string   `json:"syntax"`
	Access      string   `json:"access"`
	Status      string   `json:"status"`
	DefVal      string   `json:"def_val"`
	Indexes     []string `json:"indexes"`
	Description string   `json:"description"`
	IsIndex     bool     `json:"is_index"`
	IsRow       bool     `json:"is_row"`
	Type        OidType  `json:"type"`
	children    []*Oid
}

func CreateNewOid(name, OID, mib string) Oid {
	return Oid{
		Name:     name,
		OID:      OID,
		Mib:      mib,
		children: []*Oid{},
	}
}

func (o *Oid) FindDirectChild(parentName string) {

}

func (o *Oid) AddChildren(childrenOids ...*Oid) {
	o.children = append(o.children, childrenOids...)
}
