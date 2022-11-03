#!/usr/bin/env bash

echo "remove old gen files"
rm -rf internal/gen/*
rm -rf pkg/gen/*

echo "generate swagger yaml by code annotations"
docker run \
  --rm \
  --user $(id -u):$(id -g) \
  -e GOPATH=$HOME/go:/go \
  -e XDG_CACHE_HOME=/tmp/.cache \
  -v $HOME:$HOME -w $(pwd) \
  quay.io/goswagger/swagger generate spec -w internal/docs/ -o ./swagger.yaml --scan-models

echo "generate api server by swagger file"
docker run \
  --rm \
  --user $(id -u):$(id -g) \
  -e GOPATH=$HOME/go:/go \
  -e XDG_CACHE_HOME=/tmp/.cache \
  -v $HOME:$HOME -w $(pwd) \
  quay.io/goswagger/swagger generate server \
  --exclude-main \
  --target=pkg/gen \
  --template-dir=./internal/templates \
  --regenerate-configureapi \
  -f ./swagger.yaml

echo "generating client by swagger spec"
# генерим клиента по спеке без авторизации, т.к. авторизация добавляется кастомная в транспорт
docker run --rm -it --user $(id -u):$(id -g) -e GOPATH=$HOME/go:/go -e XDG_CACHE_HOME=/tmp/.cache -v $HOME:$HOME -w $(pwd) \
  quay.io/goswagger/swagger generate client -f swagger.yaml  -t pkg/gen

git add pkg/gen/*