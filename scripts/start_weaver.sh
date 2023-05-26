#!/bin/sh

pushd ../weaver-microservice/
printf "compiling project..."
go build
echo " done"

printf "generating weaver code... "
weaver generate
echo "done"

echo "starting service weaver as gke-local..."
weaver gke-local deploy weaver.toml
popd
