# TikTok-DL

[![Go Report Card](https://goreportcard.com/badge/github.com/pikami/tiktok-dl)](https://goreportcard.com/report/github.com/pikami/tiktok-dl)
[![tiktok-dl_CI](https://github.com/pikami/tiktok-dl/workflows/tiktok-dl_CI/badge.svg?branch=master)](https://github.com/pikami/tiktok-dl/actions)

A simple tiktok video downloader written in go

```diff
- This tool is not working currenly, I will revive it in the future when I have some free time
```

## Basic usage examples
Download the executable from `https://github.com/pikami/tiktok-dl/releases`\
You can download all videos from user by running `./tiktok-dl TIKTOK_USERNAME`\
You can download single video by running `./tiktok-dl VIDEO_URL`\
You can download items listed in a text file by running `./tiktok-dl -batch-file path/to/items.txt`

## Usage Manual
```
Usage: tiktok-dl [OPTION]... TARGET
  or:  tiktok-dl [OPTION]... -batch-file BATCH_FILE

In the 1st form, download given `TARGET`.
In the 2nd form, download all targets listed in given `BATCH_FILE`.
```

## Available options
* `-archive` - Download only videos not listed in the archive file. Record the IDs of all downloaded videos in it.
* `-batch-file some_file` - File containing URLs/Usernames to download, one value per line. Lines starting with '#', are considered as comments and ignored.
* `-deadline` - Sets the timout for scraper logic in seconds (used as a workaround for context deadline exceeded error) (default 1500)
* `-debug` - enables debug mode
* `-fail-log some_file` - Write failed items to log file
* `-json` - Returns whole data, that was scraped from TikTok, in json
* `-limit` - Sets the max count of video that will be downloaded (default infinity)
* `-metadata` - Write video metadata to a .json file
* `-output some_directory` - Output path (default "./downloads")
* `-quiet` - Suppress output

## Build instructions
1. Clone this repository
2. Run `go build` to build the executable.

## Acknowledgments
This software uses the **chromedp** for web scraping, it can be found here: https://github.com/chromedp/chromedp \
For releases the JS code is minified by using **terser** toolkit, it can be found here: https://github.com/terser/terser
