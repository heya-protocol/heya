package types

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DenomCreationFee: "1000" + "000000uheya",
		Denoms:           []*DenomAuthority{},
	}
}

func (gs *GenesisState) Validate() error {
	return nil
}
