package writer

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/encoding/json"

	"github.com/almog-t/state-proof-query-service/servicestate"
)

type Writer struct {
	bucket string
	key    string
}

func InitializeWriter(bucket string, key string) *Writer {
	return &Writer{
		bucket: bucket,
		key:    key,
	}
}

func (w *Writer) UploadStateProof(state servicestate.ServiceState, proof *models.StateProof) error {
	encodedProof := json.Encode(proof)
	proofReader := bytes.NewReader(encodedProof)
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(w.bucket),
		Key:    aws.String(w.key),
		Body:   proofReader,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Uploaded proof to %s\n", result.Location)
	state.SavedState.LatestCompletedAttestedRound = proof.Message.Lastattestedround

	return nil
}
