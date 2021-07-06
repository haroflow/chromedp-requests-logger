# chromedp-requests-logger
Opens a site and logs all requests, printing accessed domains and URLs.

Tested on WSL only, should work on other systems.

## What it solves

This tool was made as a proof of concept, to be run when the user needs to know which domains are required by a website, so they can be allowed through the firewall or proxy for example.

`wget` has options to do this, but content loaded via JS is not shown.

Disclaimer: This is not being used in production, so there may be some special cases to handle.

## How it works

We use [chromedp](https://github.com/chromedp/chromedp) to navigate to a site and log all subsequent request's URLs.

This is done inside a [chromedp/headless-shell](https://hub.docker.com/r/chromedp/headless-shell/) container, which contains all dependencies needed by `chromedp`.

If you have google-chrome installed, you don't need to use the docker image. Just run chromedp-requests-logger directly.

## Requirements

- go 1.16
- docker, if you don't have chrome installed.

## Command line arguments
```
  -outputformat string
        "text" or "json" (default "text")
  -printurls
        print urls for each domain
  -url string
        url to scan
```

## Example output:
```
$ go run . -url https://www.youtube.com

fonts.googleapis.com
www.google.com.br
yt3.ggpht.com
www.youtube.com
accounts.google.com
fonts.gstatic.com
googleads.g.doubleclick.net
www.gstatic.com
www.google.com
i.ytimg.com

$ go run . -url https://www.google.com -printurls

www.google.com
- https://www.google.com/
- https://www.google.com/tia/tia.png
- ...
www.gstatic.com
- https://www.gstatic.com/inputtools/images/tia.png
- https://www.gstatic.com/og/_/js/k=og.qtm.en_US.Mc...
- ...
apis.google.com
- https://apis.google.com/_/scs/abc-static/_/js/k=g...
ogs.google.com
- https://ogs.google.com/widget/app/so?bc=1&origin=...
adservice.google.com
- https://adservice.google.com/adsid/google/ui
```

## How to run

If you have google-chrome installed, you don't need to use the docker image.

Just run chromedp-requests-logger:

```
git clone https://github.com/haroflow/chromedp-requests-logger
cd chromedp-requests-logger
go run .
```

## For development, with docker image

If you need to run in environments without google-chrome (for example, inside WSL), you can use a docker image.

Steps:
- Build the go application
- Start the chromedp/headless-shell container
- Run chromedp-requests-logger on the container and print the output
- Stop the container

On linux this is done using the helper script:
```
./build-and-run-on-container.sh -url https://www.youtube.com
```

## Building a docker image

Steps:
- Build `chromedp-requests-logger`
- Build an image based on `chromedp/headless-shell`

On linux this is done using the helper script:
```
./build-image.sh
```

To execute the image:
```
./run-image.sh -url https://www.youtube.com
```