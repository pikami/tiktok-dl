package client

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"

	config "github.com/pikami/tiktok-dl/models/config"
	res "github.com/pikami/tiktok-dl/resources"
	utils "github.com/pikami/tiktok-dl/utils"
	log "github.com/pikami/tiktok-dl/utils/log"
)

// GetMusicUploads - Get all uploads by given music
func executeClientAction(url string, jsAction string) (string, error) {
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		return "", err
	}
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
		chromedp.WithLogf(log.Logf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Duration(config.Config.Deadline)*time.Second)
	defer cancel()

	jsOutput, err := runScrapeWithInfo(ctx, jsAction, url)
	if strings.HasPrefix(jsOutput, "\"ERR:") {
		err = errors.New(jsOutput)
	}
	return jsOutput, err
}

func runScrapeQuiet(ctx context.Context, jsAction string, url string) (string, error) {
	var jsOutput string
	if err := chromedp.Run(ctx,
		// Navigate to user's page
		chromedp.Navigate(url),
		// Execute url grabber script
		chromedp.EvaluateAsDevTools(utils.GetScraper(), &jsOutput),
		chromedp.EvaluateAsDevTools(jsAction, &jsOutput),
		// Wait until custom js finishes
		chromedp.WaitVisible(`video_urls`),
		// Grab url links from our element
		chromedp.InnerHTML(`video_urls`, &jsOutput),
	); err != nil {
		return "", err
	}

	return jsOutput, nil
}

func runScrapeWithInfo(ctx context.Context, jsAction string, url string) (string, error) {
	var jsOutput string
	if err := chromedp.Run(ctx,
		// Navigate to user's page
		chromedp.Navigate(url),
		// Execute url grabber script
		chromedp.EvaluateAsDevTools(utils.GetScraper(), &jsOutput),
		chromedp.EvaluateAsDevTools(jsAction, &jsOutput),
	); err != nil {
		return "", err
	}

	for {
		if err := chromedp.Run(ctx, chromedp.EvaluateAsDevTools("currentState.preloadCount.toString()", &jsOutput)); err != nil {
			return "", err
		}

		if jsOutput != "0" {
			log.Logf(res.PreloadingItemsFound, jsOutput)
		} else {
			log.Logf(res.Preloading)
		}

		if err := chromedp.Run(ctx, chromedp.EvaluateAsDevTools("currentState.finished.toString()", &jsOutput)); err != nil {
			return "", err
		}

		if jsOutput == "true" {
			break
		}

		time.Sleep(50 * time.Millisecond)
	}

	log.Log(res.Retrieving)
	if err := chromedp.Run(ctx,
		// Wait until custom js finishes
		chromedp.WaitVisible(`video_urls`),
		// Grab url links from our element
		chromedp.InnerHTML(`video_urls`, &jsOutput),
	); err != nil {
		return "", err
	}

	return jsOutput, nil
}
