package main

import (
	"os"
	"bytes"
	"context"
	"encoding/json"
	"time"
	"math/rand"
	"encoding/hex"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/s3"
)

type Response events.APIGatewayProxyResponse

func TempFileName(prefix, suffix string) string {
    randBytes := make([]byte, 16)
    rand.Read(randBytes)
    return prefix+hex.EncodeToString(randBytes)+suffix
}

func JSON(t interface{}) ([]byte, error) {
    buffer := &bytes.Buffer{}
    encoder := json.NewEncoder(buffer)
    encoder.SetEscapeHTML(false)
    err := encoder.Encode(t)
    return buffer.Bytes(), err
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	//var buf bytes.Buffer
        svc := s3.New(session.New())

	s3key := TempFileName("tmp-","")

	var bucketName = os.Getenv("UPLOAD_BUCKET")

        req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
	  Bucket: aws.String(bucketName),
          Key:    &s3key,
        })
        urlStr, err := req.Presign(15 * time.Minute)

	body, err := JSON(map[string]interface{}{
                "url": urlStr,
                "file": s3key,
        })


	if err != nil {
		return Response{StatusCode: 404}, err
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:		 string(body),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Access-Control-Allow-Headers": "*",
			"Content-Type":           "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
