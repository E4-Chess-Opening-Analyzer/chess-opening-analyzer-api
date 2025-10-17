package models


type Moves struct {
    WhiteWin     int                 `bson:"white_win" json:"white_win"`
    Draw         int                 `bson:"draw" json:"draw"`
    BlackWin     int                 `bson:"black_win" json:"black_win"`
    Piece        string              `bson:"piece" json:"piece"`
    PreviousRow  int                 `bson:"previous_row" json:"previous_row"`
    PreviousCol  int                 `bson:"previous_column" json:"previous_column"`
    NewRow       int                 `bson:"new_row" json:"new_row"`
    NewCol       int                 `bson:"new_column" json:"new_column"`
    Next         map[string]*Moves `bson:"next" json:"next"`
}