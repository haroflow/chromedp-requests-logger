// Command visible is a chromedp example demonstrating how to wait until an
// element is visible.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var scanUrl = flag.String("url", "", "url to scan")
var outputFormat = flag.String("outputformat", "text", `"text" or "json"`)
var printUrls = flag.Bool("printurls", false, "print urls for each domain")

func main() {
	flag.Parse()

	// Check args
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	// Init chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Allocate domain map, where each domain has a list of request urls
	domains := map[string][]string{}

	// Add event listeners to get requests and urls
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent: // Outgoing requests
			u2, err := url.ParseRequestURI(ev.Request.URL)
			if err != nil {
				fmt.Printf("failed to parse url %q: %v\n", ev.Request.URL, err)
				break
			}

			domains[u2.Host] = append(domains[u2.Host], ev.Request.URL)
		}
	})

	// Run the scan
	err := chromedp.Run(ctx, chromedp.Navigate(*scanUrl))
	if err != nil {
		log.Fatal(err)
	}

	// Output
	switch *outputFormat {
	case "text":
		for domain, urls := range domains {
			fmt.Println(domain)

			if !*printUrls {
				continue
			}

			for _, url := range urls {
				fmt.Println("-", url)
			}
		}
	case "json":
		if *printUrls {
			json.NewEncoder(os.Stdout).Encode(domains)
		} else {
			d := []string{}
			for domain := range domains {
				d = append(d, domain)
			}
			json.NewEncoder(os.Stdout).Encode(d)
		}
	default:
		fmt.Println("Unknown output format:", *outputFormat)
		os.Exit(1)
	}
}
