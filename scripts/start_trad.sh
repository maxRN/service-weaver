#!/bin/bash

echo "starting main service..."
go run ../trad-microservice/main-service/main.go 8070 8069 &
go run ../trad-microservice/main-service/main.go 8071 8068 &

echo "starting worker..."
go run ../trad-microservice/work-service/main.go 8069 &
go run ../trad-microservice/work-service/main.go 8068
