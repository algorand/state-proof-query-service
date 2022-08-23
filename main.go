package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/almog-t/state-proof-query-service/querier"
	"github.com/almog-t/state-proof-query-service/servicestate"
	"github.com/almog-t/state-proof-query-service/utilities"
	"github.com/almog-t/state-proof-query-service/writer"
)

type ServiceConfiguration struct {
	LogPath         string
	GenesisRound    uint64
	StatePath       string
	NodePath        string
	BucketName      string
	BucketRegion    string
	BucketKey       string
	BucketSecretKey string
	BackoffTimeMs   time.Duration
}

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

	log.Printf("Fetched proof for round %d", nextStateProofData.Message.Lastattestedround)
	err = state.Save()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var config ServiceConfiguration
	err := utilities.DecodeFromFile(&config, "config.json")

	logFile, err := os.OpenFile(config.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file log file: %s", err)
	}

	log.SetOutput(logFile)
	state, err := servicestate.InitializeState(config.StatePath, config.GenesisRound)
	if err != nil {
		log.Printf("Could not initialize state: %s", err)
		return
	}

	nodeQuerier, err := querier.InitializeQuerier(config.NodePath)
	if err != nil {
		log.Printf("Could not initialize querier: %s", err)
		return
	}

	s3Writer := writer.InitializeWriter(config.BucketName, config.BucketRegion, config.BucketKey, config.BucketSecretKey)

	for {
		err = fetchStateProof(state, *nodeQuerier, *s3Writer)
		if err == nil {
			continue
		}

		if strings.Contains(err.Error(), "HTTP 404") || strings.Contains(err.Error(), "given round is greater than the latest round") {
			time.Sleep(config.BackoffTimeMs * time.Millisecond)
			continue
		}

		log.Printf("Error while fetching state proof: %s", err)
		break
	}
}
