#!/bin/bash

pushd "functions"

go get all
go mod vendor

popd

pushd "common"

go get all
go mod vendor

popd