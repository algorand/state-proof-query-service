package querier

import (
	"context"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/almog-t/state-proof-query-service/servicestate"
)

type Querier struct {
	client                  *algod.Client
	lastCompletedProofRound uint64
}

func InitializeQuerier(algodAddress string, apiToken string) (*Querier, error) {
	client, err := algod.MakeClient(algodAddress, apiToken)
	if err != nil {
		return nil, err
	}

	return &Querier{
		client: client,
	}, nil
}

func (q *Querier) QueryNextStateProofData(state *servicestate.ServiceState) (models.StateProof, error) {
	return q.client.GetStateProof(state.SavedState.LastCompletedStateProof).Do(context.Background())
}
