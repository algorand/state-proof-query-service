package main

import (
	"fmt"
	"time"

	"github.com/almog-t/state-proof-query-service/querier"
	"github.com/almog-t/state-proof-query-service/servicestate"
	"github.com/almog-t/state-proof-query-service/writer"
)

func fetchStateProof(state *servicestate.ServiceState, nodeQuerier querier.Querier, s3Writer writer.Writer) error {
	err := state.Load()
	if err != nil {
		return err
	}

	nextStateProofData, err := nodeQuerier.QueryNextStateProofData(state)
	if err != nil {
		return err
	}

	err = s3Writer.UploadStateProof(state, nextStateProofData)
	if err != nil {
		return err
	}

	err = state.Save()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	state, err := servicestate.InitializeState("state.txt")
	if err != nil {
		fmt.Printf("Could not initialize state: %s\n", err)
		return
	}

	nodeQuerier, err := querier.InitializeQuerier("/Users/almog/go/src/light-client-demo/demo_instance/demo_network/2")
	if err != nil {
		fmt.Printf("Could not initialize querier: %s\n", err)
	}

	s3Writer := writer.InitializeWriter("bucket", "key")

	for {
		time.Sleep(500 * time.Millisecond)
		err = fetchStateProof(state, *nodeQuerier, *s3Writer)
		if err != nil {
			fmt.Printf("Error while fetching state proof: %s\n", err)
			return
		}
	}
}
