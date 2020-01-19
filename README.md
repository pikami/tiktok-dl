# TikTok-DL

[![Go Report Card](https://goreportcard.com/badge/github.com/pikami/tiktok-dl)](https://goreportcard.com/report/github.com/pikami/tiktok-dl)

A simple tiktok video downloader written in go

## Basic usage
Clone this repository and run `go build` to build the executable.\
You can download all videos from user by running `./tiktok-dl [Options] TIKTOK_USERNAME`\
You can download single video by running `./tiktok-dl [Options] VIDEO_URL`

## Available options
* `-debug` - enables debug mode
* `-output some_directory` - Output path (default "./downloads")

## Acknowledgments
This software uses the chromedp for web scraping, it can be found here: https://github.com/chromedp/chromedp