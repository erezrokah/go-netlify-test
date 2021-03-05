package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"microappsdev.com/go-test/internal/pkg/utils"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Print("Got request for '/.netlify/functions/goodbye', this message is dumpled by 'cmd/goodbye/main.go'")
	lc, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return &events.APIGatewayProxyResponse{
			StatusCode: 503,
			Body:       "Something went wrong :(",
		}, nil
	}

	cc := lc.ClientContext
	fmt.Printf("cc: %#v\n", cc)
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       utils.IntroduceYourself(cc.Client.AppTitle),
	}, nil
}

func main() {
	lambda.Start(handler)
}
