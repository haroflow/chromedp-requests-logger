#!/bin/bash

# This starts the docker image (built with build-image.sh) and
# executes chromedp-requests-logger passing all parameters

# USAGE:
# ./run-image.sh -url https://www.youtube.com

IMAGE_NAME=chromedp-requests-logger

docker run --init chromedp-requests-logger /chromedp-requests-logger $@
