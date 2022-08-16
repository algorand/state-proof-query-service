package servicestate

import (
	"github.com/almog-t/state-proof-query-service/utilities"
)

type ServiceState struct {
	LatestCompletedAttestedRound uint64
	filePath                     string
}

func InitializeState(filePath string, genesisRound uint64) (*ServiceState, error) {
	state := ServiceState{
		LatestCompletedAttestedRound: genesisRound,
		filePath:                     filePath,
	}

	err := state.Load()
	return &state, err
}

func (s *ServiceState) Save() error {
	return utilities.EncodeToFile(&s, s.filePath)
}

func (s *ServiceState) Load() error {
	return utilities.DecodeFromFile(&s, s.filePath)
}
