package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Handler is
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("ログ出ているか確認")
	line := Line{}
	err := line.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"))
	if err != nil {
		//TODO ちゃんとエラーハンドリングする
		fmt.Println(err)
	}
	eve, err := ParseRequest(line.ChannelSecret, request)
	if err != nil {
		log.Println("エラーログ")
		log.Println(err)
		status := 200
		if err == linebot.ErrInvalidSignature {
			status = 400
		} else {
			status = 500
		}
		return events.APIGatewayProxyResponse{StatusCode: status}, errors.New("Bad Request")
	}
	line.EventRouter(eve)
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
