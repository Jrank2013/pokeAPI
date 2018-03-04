package pokemon

//Pokemon contains all the information about a given pokemon
type Pokemon struct {
	ID             int
	Name           string
	BaseExperience int `json:"base_experience"`
	Height         int
	IsDefault      bool `json:"is_default"`
	Order          int
	Weight         int
	Abilities      []AbilitesList
	Forms          []FormsList
	Stats          []StatsList
	Types          []TypesList
}
