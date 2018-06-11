package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	lambdaInit "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func handleEvent(ctx context.Context) (string, error) {
	if os.Getenv("INVOKE_FUNCTION") == "" {
		return "", errors.New("INVOKE_FUNCTION env required")
	}

	invokeLambda := lambda.New(session.Must(session.NewSession()))
	payload := os.Getenv("PAYLOAD")

	_, err := invokeLambda.Invoke(&lambda.InvokeInput{
		FunctionName:   aws.String(os.Getenv("INVOKE_FUNCTION")),
		Payload:        []byte(payload),
		InvocationType: aws.String("Event"),
	})

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Invoked lambda with payload: %s", payload), nil
}

func main() {
	lambdaInit.Start(handleEvent)
}
