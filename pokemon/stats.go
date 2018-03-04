package pokemon

//StatsList contins list of stats
type StatsList struct {
	BaseStat int `json:"base_stat"`
	Effort   int
	Stat     statSimple
}

//StatSimple contains a single stat
type statSimple struct {
	Name string
	URL  string
}

//StatAll contins details infomration about a stat
type StatAll struct {
	ID         int
	Name       string
	GameIndex  int  `json:"game_index"`
	BattleOnly bool `json:"is_battle_only"`
	// AffectingMoves  []affmoves  `json:"affecting_moves"`
	// AffectingNature []affnature `json:"affecting_natures"`
}

// type affmoves struct {
// 	Increase []
// 	Decrease
// }

// type affnature struct {

// }
