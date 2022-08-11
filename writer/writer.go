package writer

import (
	"fmt"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Writer struct {
	bucket string
	key    string
}

func (w *Writer) UploadStateProof(proof *models.StateProof) {
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(w.bucket),
		Key:    aws.String(w.key),
		Body:   proof[:],
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
}
