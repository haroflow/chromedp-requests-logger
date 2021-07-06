FROM chromedp/headless-shell:latest

WORKDIR /
COPY chromedp-requests-logger /

# Clears the entrypoint for execution of the chromedp-requests-logger application
ENTRYPOINT []
