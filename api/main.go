package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/maybemanolo/api/search"
)

type EventResponse struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}

type EventRequest struct {
	Number string `json:"number"`
}

func apiGWHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var evtReq EventRequest
	err := json.Unmarshal([]byte(req.Body), &evtReq)
	if err != nil {
		return apiGWResponse("", 0, http.StatusInternalServerError, err), err
	}
	number, index := search.Get(evtReq.Number)
	return apiGWResponse(number, index, http.StatusOK, nil), nil
}

func apiGWResponse(result string, index int, status int, err error) events.APIGatewayProxyResponse {
	resp := events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	if err != nil {
		resp.Body = err.Error()
		return resp
	}

	evtResp := EventResponse{
		Index: index,
		Text:  result,
	}

	data, err := json.Marshal(evtResp)
	if err != nil {
		resp.StatusCode = http.StatusInternalServerError
		resp.Body = err.Error()
	}
	resp.Body = string(data)
	return resp
}

func main() {
	lambda.Start(apiGWHandler)
}
