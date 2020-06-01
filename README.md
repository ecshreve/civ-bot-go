# civ-bot-go

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/ecshreve/civ-bot-go/Go)
![Go Report Card](https://goreportcard.com/badge/github.com/ecshreve/civ-bot-go)

This a multipurpose Discord bot for various operations related to Civilization 5, written in Go.

## Running Locally

### Bot Setup

- make a discord application in the [discord developer portal][1]
- create a bot for you application and set permissions
  ![bot perms](static/botperms.png "bot perms")
- add the bot to your discord server
  - go to the "OAuth2" tab, select scope "bot", navigate to the URL that gets generated
- copy the bot token from the "Bot" tab and set the `CIV_BOT_TOKEN` environment variable

### Application Setup

- make sure you have `go` installed by running `go version`, you should see some output like `/usr/local/bin/go`
- clone the repository
- `cd` into the root directory
- install dependencies with `go get github.com/ecshreve/civ-bot-go`
- run the application with `make run`
- now you can go to whatever channel you added the bot to in the previous section and interact with the bot, enter `/civ help` in the channel to see information about how to use the bot

## References

Here's some links that I used as reference throughout the project:

repositories:

- [discordgo](https://github.com/bwmarrin/discordgo)
- [discord-checkers](https://github.com/jmsheff/discord-checkers)

helpful golang links

- [using go modules](https://blog.golang.org/using-go-modules)
- [golang slice tricks](https://github.com/golang/go/wiki/SliceTricks)

helpful tooling links

- [hosting a discord bot on heroku](https://medium.com/@mason.spr/hosting-a-discord-js-bot-for-free-using-heroku-564c3da2d23f)
- [automate deploys to heroku from github](https://devcenter.heroku.com/articles/github-integration)
- [basic makefiles](https://tutorialedge.net/golang/makefiles-for-go-developers/)

[1]: https://discord.com/developers/applications
