package main

import (
	"fmt"
	"time"

	"github.com/almog-t/state-proof-query-service/querier"
	"github.com/almog-t/state-proof-query-service/servicestate"
	"github.com/almog-t/state-proof-query-service/writer"
)

func fetchStateProof(state servicestate.ServiceState, nodeQuerier querier.Querier, s3Writer writer.Writer) error {
	nextStateProofData, err := nodeQuerier.QueryNextStateProofData(state)
	if err != nil {
		return err
	}

	err = s3Writer.UploadStateProof(nextStateProofData)
	if err != nil {
		return err
	}

	state.SavedState.LatestCompletedAttestedRound = 3
	err = state.Save()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	state, err := servicestate.InitializeState("/Users/almog/go/src/github.com/almog-t/state-proof-query-service")
	if err != nil {
		fmt.Printf("Could not initialize state: %s\n", err)
		return
	}

	nodeQuerier, err := querier.InitializeQuerier("http://127.0.0.1")

	s3Writer := writer.InitializeWriter("bucket", "key")

	queryTicker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			select {
			case _ = <-queryTicker.C:
				err = fetchStateProof(*state, *nodeQuerier, *s3Writer)
				if err != nil {
					fmt.Printf("Error while fetching state proof: %s\n", err)
					return
				}
			}
		}
	}()
}
