module github.com/MSU-Bot/Serverless/functions

go 1.13

require (
	cloud.google.com/go v0.53.0
	cloud.google.com/go/firestore v1.1.1
	github.com/MSU-Bot/Serverless/common v0.0.0
	github.com/plivo/plivo-go v4.1.5+incompatible
	github.com/sirupsen/logrus v1.4.2
)

replace github.com/MSU-Bot/Serverless/common => ../common
