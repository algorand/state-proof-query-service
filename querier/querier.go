package querier

import (
	"context"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/encoding/msgpack"
	"github.com/algorand/go-algorand-sdk/stateproofs/transactionverificationtypes"
)

type Querier struct {
	client *algod.Client
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

func (q *Querier) QueryStateProofData(round uint64) (transactionverificationtypes.Message,
	*transactionverificationtypes.EncodedStateProof, error) {
	stateProofData, err := q.client.GetStateProof(round).Do(context.Background())
	if err != nil {
		return transactionverificationtypes.Message{}, nil, err
	}

	var attestedMessage transactionverificationtypes.Message
	err = msgpack.Decode(stateProofData.Message, &attestedMessage)

	if err != nil {
		return transactionverificationtypes.Message{}, nil, err
	}

	return attestedMessage, (*transactionverificationtypes.EncodedStateProof)(&stateProofData.Stateproof), nil
}
