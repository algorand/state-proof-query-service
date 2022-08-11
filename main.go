package main

import (
	"fmt"
	"github.com/almog-t/state-proof-query-service/querier"
	"github.com/almog-t/state-proof-query-service/servicestate"
)

func main() {
	state, err := servicestate.InitializeState("/Users/almog/go/src/github.com/almog-t/state-proof-query-service")
	if err != nil {
		fmt.Printf("Could not initialize state: %s\n", err)
		return
	}

	querier, err := querier.InitializeQuerier("http://127.0.0.1")
}
