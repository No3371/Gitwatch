# GitWatch

A discord bot to watch git activity for you, works everywhere as long as it's a git repo.

Useful for outdated git service deployments that does not support Discord, or where webhooks is not available.

**Git is required to be on system path.**

**To build:** `go build`

**To run:** `gitwatch.exe --repo [path] --token [bot_token] --cid [channelId]` (gitwatch.exe is the built binary file)

- `path`: The absolute path to the repo to be watch.
- `bot_token`: The discord bot token. Get yours on Discord Developer Portal.
- `cid`: The channel Id (number) where the notification will be sent.