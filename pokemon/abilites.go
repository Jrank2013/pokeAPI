package pokemon

//AbilitesList contains list of abilities
type AbilitesList struct {
	IsHidden bool `json:"is_hidden"`
	Slot     int
	Ability  AbilitySimple
}

//AbilitySimple contains information about a given ability
type AbilitySimple struct {
	Name string
	URL  string
}
