package servicestate

import (
	"encoding/json"
	"io/fs"
	"os"
)

type SavedServiceState struct {
	LatestCompletedAttestedRound uint64
}

type ServiceState struct {
	SavedState SavedServiceState
	filePath   string
}

func InitializeState(filePath string) (*ServiceState, error) {
	state := ServiceState{
		SavedState: SavedServiceState{LatestCompletedAttestedRound: 0},
		filePath:   filePath,
	}

	err := state.Load()
	return &state, err
}

func (s *ServiceState) Save() error {
	encodedData, err := json.Marshal(s.SavedState)
	if err != nil {
		return err
	}

	err = os.WriteFile(s.filePath, encodedData, fs.ModePerm)
	return err
}

func (s *ServiceState) Load() error {
	encodedData, err := os.ReadFile(s.filePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	if err != nil {
		return err
	}

	err = json.Unmarshal(encodedData, &s.SavedState)
	return err
}
