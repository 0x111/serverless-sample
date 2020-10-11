package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var gorillaLambda *gorillamux.GorillaMuxAdapter

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/today", getCurrentDate).Methods("GET")
	gorillaLambda = gorillamux.New(r)
}

func getCurrentDate(writer http.ResponseWriter, request *http.Request) {
	date := time.Now().String()
	fmt.Fprintf(writer, date)
}

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return gorillaLambda.ProxyWithContext(ctx, req)
}
