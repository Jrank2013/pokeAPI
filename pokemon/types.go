package pokemon

//TypesList contains list of types
type TypesList struct {
	Slot int
	Type TypeSimple
}

//TypeSimple contains information about a given type
type TypeSimple struct {
	Name string
	URL  string
}
