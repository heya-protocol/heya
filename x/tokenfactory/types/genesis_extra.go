package types

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DenomCreationFee: "1000000000uheya",
		Denoms:           []*DenomAuthority{},
	}
}

func (gs *GenesisState) Validate() error {
	return nil
}
