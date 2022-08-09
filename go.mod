module github.com/almog-t/state-proof-query-service

go 1.17

require github.com/algorand/go-algorand-sdk v1.18.0

require (
	github.com/algorand/go-codec/codec v1.1.8 // indirect
	github.com/algorand/go-sumhash v0.1.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/sys v0.0.0-20220808155132-1c4a2a72c664 // indirect
)

replace github.com/algorand/go-algorand-sdk => github.com/almog-t/go-algorand-sdk v1.14.1-0.20220808115008-3b50e4b05790
