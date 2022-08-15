module github.com/almog-t/state-proof-query-service

go 1.17

require (
	github.com/algorand/go-algorand-sdk v1.18.0
	github.com/aws/aws-sdk-go v1.44.73
)

require (
	github.com/algorand/go-codec/codec v1.1.8 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
)

replace github.com/algorand/go-algorand-sdk => github.com/almog-t/go-algorand-sdk v1.14.1-0.20220815130753-4f4f4a46360f
