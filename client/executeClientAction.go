package client

import (
	"context"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"time"

	config "../models/config"
	utils "../utils"
)

// GetMusicUploads - Get all uploads by given music
func executeClientAction(url string, jsAction string) string {
	dir, err := ioutil.TempDir("", "chromedp-example")
	utils.CheckErr(err)
	defer os.RemoveAll(dir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.UserDataDir(dir),
		chromedp.Flag("headless", !config.Config.Debug),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Duration(config.Config.Deadline)*time.Second)
	defer cancel()

	var jsOutput string
	jsOutput = runScrapeWithInfo(ctx, jsAction, url)

	return jsOutput
}

func runScrapeQuiet(ctx context.Context, jsAction string, url string) string {
	var jsOutput string
	err := chromedp.Run(ctx,
		// Navigate to user's page
		chromedp.Navigate(url),
		// Execute url grabber script
		chromedp.EvaluateAsDevTools(utils.ReadFileAsString("scraper.js"), &jsOutput),
		chromedp.EvaluateAsDevTools(jsAction, &jsOutput),
		// Wait until custom js finishes
		chromedp.WaitVisible(`video_urls`),
		// Grab url links from our element
		chromedp.InnerHTML(`video_urls`, &jsOutput),
	)
	utils.CheckErr(err)
	return jsOutput
}

func runScrapeWithInfo(ctx context.Context, jsAction string, url string) string {
	var jsOutput string
	err := chromedp.Run(ctx,
		// Navigate to user's page
		chromedp.Navigate(url),
		// Execute url grabber script
		chromedp.EvaluateAsDevTools(utils.ReadFileAsString("scraper.js"), &jsOutput),
		chromedp.EvaluateAsDevTools(jsAction, &jsOutput),
	)
	utils.CheckErr(err)

	for {
		err = chromedp.Run(ctx, chromedp.EvaluateAsDevTools("currentState.preloadCount.toString()", &jsOutput))
		utils.CheckErr(err)
		if jsOutput != "0" {
			utils.Logf("\rPreloading... Currently loaded %s items.", jsOutput)
		} else {
			utils.Logf("\rPreloading...")
		}

		err = chromedp.Run(ctx, chromedp.EvaluateAsDevTools("currentState.finished.toString()", &jsOutput))
		utils.CheckErr(err)
		if jsOutput == "true" {
			break
		}

		time.Sleep(50 * time.Millisecond)
	}

	utils.Log("\nRetrieving items...")
	err = chromedp.Run(ctx,
		// Wait until custom js finishes
		chromedp.WaitVisible(`video_urls`),
		// Grab url links from our element
		chromedp.InnerHTML(`video_urls`, &jsOutput),
	)
	utils.CheckErr(err)

	return jsOutput
}
