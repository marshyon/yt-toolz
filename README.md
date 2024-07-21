# YouTube RSS Feed Finder

Finding the id of a YouTube channel is something obfuscated to the user where a YouTube channel is typically visible by the channel 'handle' for example

`https://www.youtube.com/@ThePrimeTimeagen` => which is the channel of The Primagen

But if you want to get an RSS feed to this or any other channel, you need to get the actual ID someehow and this is not displayed to users in the app or web page view, however it is to be found in the HTML source of the HTML page. The ID follows the following string :

https://www.youtube.com/channel

this ID may then be appended to another URL to get the RSS feed

https://www.youtube.com/feeds/videos.xml?channel_id=

# Usage

```bash
YouTube Toolz is a CLI tool for getting public YouTube data.
examples:

yt-toolz get --name "@somechannel"

Usage:
  yt-toolz [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         gets an RSS feed from a YouTube channel
  help        Help about any command

Flags:
  -h, --help   help for yt-toolz

Use "yt-toolz [command] --help" for more information about a command.
```

# About

this project is written in Go and initialised with `cobra-cli init` and `cobra-cli add get`

# References

https://mixedanalytics.com/blog/find-a-youtube-channel-id/
