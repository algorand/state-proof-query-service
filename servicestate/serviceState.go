package servicestate

import (
	"github.com/almog-t/state-proof-query-service/utilities"
)

type SavedServiceState struct {
	LatestCompletedAttestedRound uint64
}

type ServiceState struct {
	SavedState SavedServiceState
	filePath   string
}

func InitializeState(filePath string, genesisRound uint64) (*ServiceState, error) {
	state := ServiceState{
		SavedState: SavedServiceState{LatestCompletedAttestedRound: genesisRound},
		filePath:   filePath,
	}

	err := state.Load()
	return &state, err
}

func (s *ServiceState) Save() error {
	return utilities.EncodeToFile(s.SavedState, s.filePath)
}

func (s *ServiceState) Load() error {
	return utilities.DecodeFromFile(&s.SavedState, s.filePath)
}
