package fruit

type Type int

const (
	AppleType  Type = 0
	BananaType Type = 1
)

type Fruit interface {
	GetName() string
	SetName(name string)

	GetType() Type
}
