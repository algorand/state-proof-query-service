package main

import (
	"fmt"
	"strings"
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

// TODO: Logging
// TODO: What if the latest block is not in the network?
// TODO: Separate file for bucket and key
func main() {
	state, err := servicestate.InitializeState("state.txt")
	if err != nil {
		fmt.Printf("Could not initialize state: %s\n", err)
		return
	}

	nodeQuerier, err := querier.InitializeQuerier("/Users/almog/go/src/light-client-demo/demo_instance/demo_network/2")
	if err != nil {
		fmt.Printf("Could not initialize querier: %s\n", err)
		return
	}

	s3Writer := writer.InitializeWriter(BUCKET, ACCESS_KEY)

	for {
		err = fetchStateProof(state, *nodeQuerier, *s3Writer)
		if err == nil {
			continue
		}

		if strings.Contains(err.Error(), "HTTP 404") {
			time.Sleep(500 * time.Millisecond)
			continue
		}

		fmt.Printf("Error while fetching state proof: %s\n", err)
		break
	}
}
