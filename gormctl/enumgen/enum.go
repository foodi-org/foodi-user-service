package enumgen

type (
	CardType   int8
	GenderType string
)

const (
	FIRSTGENERATIONIDCARD  CardType = 1
	SECONDGENERATIONIDCARD CardType = 2
	BUNSINESSLICENSE       CardType = 3
)

const (
	MALE   = "male"
	FEMALE = "female"
)
