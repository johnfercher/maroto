package propsmapper

// Protection is the representation of a pdf protection.
type Protection struct {
	Type          byte
	UserPassword  string
	OwnerPassword string
}

// NewPageNumber is responsible for creating the pageNumber, if the pageNumber fields cannot be
// converted, an invalid value is set.
func NewProtection(protection interface{}) *Protection {
	protectionMap, ok := protection.(map[string]interface{})
	if !ok {
		return nil
	}

	return &Protection{
		Type:          factoryTypeProtection(*convertFields(protectionMap["type"], "None")),
		UserPassword:  *convertFields(protectionMap["user_password"], ""),
		OwnerPassword: *convertFields(protectionMap["owner_password"], ""),
	}
}

func factoryTypeProtection(typeProtection string) byte {
	switch typeProtection {
	case "Print":
		return 4
	case "Modify":
		return 8
	case "Copy":
		return 16
	case "AnnotForms":
		return 32
	}

	return 0
}
