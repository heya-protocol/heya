package types

const (
	ModuleName  = "tokenfactory"
	StoreKey    = ModuleName
	DenomPrefix = "factory"
)

var (
	DenomKeyPrefix = []byte{0x01}
)

func DenomKey(denom string) []byte {
	return append(DenomKeyPrefix, []byte(denom)...)
}
