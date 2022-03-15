package entity

// domain では技術的ロジックは何も入れない
type User struct {
	ID        string `validate:"required"`
	AuthToken string `validate:"required"`
	Name      string `validate:"required"`
	HighScore int32  `validate:"gte=0"`
	Coin      int32  `validate:"gte=0"`
}

type Users *[]User
