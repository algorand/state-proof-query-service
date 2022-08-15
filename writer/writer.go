package writer

import (
	"bytes"
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/almog-t/state-proof-query-service/servicestate"
)

type Writer struct {
	bucket    string
	region    string
	key       string
	secretKey string
}

func InitializeWriter(bucket string, region string, key string, secretKey string) *Writer {
	return &Writer{
		bucket:    bucket,
		region:    region,
		key:       key,
		secretKey: secretKey,
	}
}

func (w *Writer) UploadStateProof(state *servicestate.ServiceState, proof *models.StateProof) error {
	encodedProof := json.Encode(proof)
	proofReader := bytes.NewReader(encodedProof)
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(w.region),
		Credentials: credentials.NewStaticCredentials(w.key, w.secretKey, ""),
	}))
	uploader := s3manager.NewUploader(sess)

	objectName := fmt.Sprintf("proof_%d_to_%d", proof.Message.Firstattestedround, proof.Message.Lastattestedround)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(w.bucket),
		Key:    aws.String(objectName),
		Body:   proofReader,
	})

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	state.SavedState.LatestCompletedAttestedRound = proof.Message.Lastattestedround

	return nil
}
