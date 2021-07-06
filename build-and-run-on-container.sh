#!/bin/bash

# For use in development.
# This builds the go application, runs the container,
# executes chromedp-requests-logger passing parameters,
# prints output and stops container.

# USAGE:
# ./build-and-run-on-container.sh -url https://www.youtube.com

CONTAINER_NAME=chromedp-requests-logger

go build . && 
	docker run -d --rm --name $CONTAINER_NAME -v $PWD:/home/ chromedp/headless-shell >/dev/null &&
	docker exec -w /home $CONTAINER_NAME /home/chromedp-requests-logger $@

docker stop -t 0 $CONTAINER_NAME >/dev/null
