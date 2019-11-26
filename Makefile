.PHONY: build clean deploy

build:
	go get github.com/aws/aws-lambda-go/lambda
	go get github.com/aws/aws-lambda-go/events
	go get github.com/line/line-bot-sdk-go/linebot
	env GOOS=linux go build -ldflags="-s -w" -o bin/bot bot/*.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
