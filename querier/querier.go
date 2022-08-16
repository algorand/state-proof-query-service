package querier

import (
	"context"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/almog-t/state-proof-query-service/servicestate"
	"os"
	"path/filepath"
	"strings"
)

type Querier struct {
	client *algod.Client
}

func readFromNodeFile(nodepath string) (string, error) {
	contentBytes, err := os.ReadFile(nodepath)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(contentBytes), "\n"), nil
}

func InitializeQuerier(nodePath string) (*Querier, error) {
	algodAddress, err := readFromNodeFile(filepath.Join(nodePath, "algod.net"))
	if err != nil {
		return nil, err
	}

	apiToken, err := readFromNodeFile(filepath.Join(nodePath, "algod.admin.token"))
	if err != nil {
		return nil, err
	}

	client, err := algod.MakeClient("http://"+algodAddress, apiToken)
	if err != nil {
		return nil, err
	}

	return &Querier{
		client: client,
	}, nil
}

func (q *Querier) QueryNextStateProofData(state *servicestate.ServiceState) (*models.StateProof, error) {
	proof, err := q.client.GetStateProof(state.LatestCompletedAttestedRound + 1).Do(context.Background())
	if err != nil {
		return nil, err
	}

	return &proof, nil
}
