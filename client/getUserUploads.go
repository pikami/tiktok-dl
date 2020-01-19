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

// GetUserUploads - Get all uploads by user
func GetUserUploads(username string) []models.Upload {
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.UserDataDir(dir),
		chromedp.Flag("headless", models.Config.UserName),
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
		chromedp.Navigate(`https://www.tiktok.com/@`+username),
		// Execute url grabber script
		chromedp.EvaluateAsDevTools(utils.ReadFileAsString("getVidLinks.js"), &jsOutput),
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
