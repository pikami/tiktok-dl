# TikTok-DL

[![Go Report Card](https://goreportcard.com/badge/github.com/pikami/tiktok-dl)](https://goreportcard.com/report/github.com/pikami/tiktok-dl)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/pikami/tiktok-dl/tiktok-dl_CI)

A simple tiktok video downloader written in go

## Basic usage
Download the executable from `https://github.com/pikami/tiktok-dl/releases`\
You can download all videos from user by running `./tiktok-dl [Options] TIKTOK_USERNAME`\
You can download single video by running `./tiktok-dl [Options] VIDEO_URL`\
You can download all videos by music by running `./tiktok-dl [Options] MUSIC_URL`\
You can download items listed in a text file by running `./tiktok-dl [OPTIONS] -batch-file path/to/items.txt`

## Build instructions
Clone this repository and run `go build` to build the executable.

## Available options
* `-archive` - Download only videos not listed in the archive file. Record the IDs of all downloaded videos in it.
* `-batch-file` - File containing URLs/Usernames to download, one value per line. Lines starting with '#', are considered as comments and ignored.
* `-deadline` - Sets the timout for scraper logic in seconds (used as a workaround for context deadline exceeded error) (default 1500)
* `-debug` - enables debug mode
* `-json` - Returns whole data, that was scraped from TikTok, in json
* `-limit` - Sets the max count of video that will be downloaded (default infinity)
* `-metadata` - Write video metadata to a .json file
* `-output some_directory` - Output path (default "./downloads")
* `-quiet` - Supress output

## Acknowledgments
This software uses the **chromedp** for web scraping, it can be found here: https://github.com/chromedp/chromedp \
For releases the JS code is minified by using **terser** toolkit, it can be found here: https://github.com/terser/terser
