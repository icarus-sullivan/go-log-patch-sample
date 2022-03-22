package main

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	methodArn := strings.Join(strings.Split(request.MethodArn, "/")[0:2], "/") + "/*/*"
	authResponse := events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: "test-id",
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   "Allow",
					Resource: []string{methodArn},
				},
			},
		},
	}

	return authResponse, nil
}

func main() {
	lambda.Start(Handler)
}
