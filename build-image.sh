#!/bin/bash

# This builds the go application, and builds the docker image.

IMAGE_NAME=chromedp-requests-logger

go build . && docker build -t $IMAGE_NAME .
