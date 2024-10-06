package propsmapper

// Protection is the representation of a pdf protection.
type Protection struct {
	Type          byte
	UserPassword  string
	OwnerPassword string
}
