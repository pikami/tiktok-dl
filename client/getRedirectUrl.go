package client

import (
	"context"
	"io/ioutil"
	"os"
	"time"

	"github.com/chromedp/chromedp"

	config "github.com/pikami/tiktok-dl/models/config"
	log "github.com/pikami/tiktok-dl/utils/log"
)

// GetRedirectURL - Returns URL that the given URL redirects to
func GetRedirectURL(url string) (string, error) {
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

	var jsOutput string
	if err := chromedp.Run(ctx,
		// Navigate to user's page
		chromedp.Navigate(url),
		// Wait until page loads
		chromedp.WaitReady(`div`),
		// Grab url links from our element
		chromedp.EvaluateAsDevTools(`window.location.href`, &jsOutput),
	); err != nil {
		return "", err
	}

	return jsOutput, err
}
