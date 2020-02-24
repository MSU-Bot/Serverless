module github.com/MSU-Bot/msubot-serverless/functions

go 1.13

require (
	cloud.google.com/go/firestore v1.1.1
	github.com/MSU-Bot/msubot-serverless/common v0.0.0
	github.com/sirupsen/logrus v1.4.2
)

replace github.com/MSU-Bot/msubot-serverless/common => ../common
