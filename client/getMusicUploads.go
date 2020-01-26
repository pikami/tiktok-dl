package client

import (
	"context"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"time"

	models "../models"
	utils "../utils"
)

// GetMusicUploads - Get all uploads by given music
func GetMusicUploads(url string) []models.Upload {
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.UserDataDir(dir),
		chromedp.Flag("headless", !models.Config.Debug),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 1500*time.Second)
	defer cancel()

	var jsOutput string
	err = chromedp.Run(ctx,
		// Navigate to user's page
		chromedp.Navigate(url),
		// Execute url grabber script
		chromedp.EvaluateAsDevTools(utils.ReadFileAsString("scraper.js"), &jsOutput),
		chromedp.EvaluateAsDevTools("bootstrapIteratingVideos()", &jsOutput),
		// Wait until custom js finishes
		chromedp.WaitVisible(`video_urls`),
		// Grab url links from our element
		chromedp.InnerHTML(`video_urls`, &jsOutput),
	)
	if err != nil {
		log.Fatal(err)
	}

	return models.ParseUploads(jsOutput)
}
