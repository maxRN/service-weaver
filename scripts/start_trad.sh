#!/bin/sh

echo "starting main service..."
go run ../trad-microservice/services/main-service/main.go &

echo "starting worker..."
go run ../trad-microservice/services/work-service/main.go
