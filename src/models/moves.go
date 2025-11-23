package models

type NextMove struct {
	Name       string `bson:"name"`
	WhiteWin   uint   `bson:"white_win"`
	Draw       uint   `bson:"draw"`
	BlackWin   uint   `bson:"black_win"`
	TotalGames uint   `bson:"total_games"`
}

type Move struct {
	ID           string     `bson:"_id"`
	MoveSequence []string   `bson:"move_sequence"`
	Depth        int        `bson:"depth"`
	WhiteWin     uint       `bson:"white_win"`
	Draw         uint       `bson:"draw"`
	BlackWin     uint       `bson:"black_win"`
	TotalGames   uint       `bson:"total_games"`
	NextMoves    []NextMove `bson:"next_moves"`
}
