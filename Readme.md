# GitWatch

A discord bot to watch git activity for you, works everywhere as long as it's a git repo.

Useful for outdated git service deployments that does not support Discord, or where webhooks is not available.

## Usage

**Git is required to be on system path.**

**To build:** `go build`

**To run:** `gitwatch.exe --repo [path] --token [bot_token] --cid [channelId]` (gitwatch.exe is the built binary file)

- `path`: The absolute path to the repo to be watch.
- `bot_token`: The discord bot token. Get yours on Discord Developer Portal.
- `channelId`: The channel Id (number) where the notification will be sent.

\*Make sure the bot is invited and have access to the channel given.

## Note

#### It's local

GitWatch works given any git repo, however, the best practice is give it a dedicated repo, so that it does not notify commits not pushed yet:

- Given RepoA, where you works: notify any new commits created locally.
- Given RepoB, which is RepoA's clone where you leave alone: notify only when commits pushed to remote.

#### Only For Cloned Branches

GitWatch basically just do `git fetch --all` and `git log --all`, these only works for already cloned branches.

## Credits

This project is packed with a customized version of [go-gitlog](https://github.com/wadackel/go-gitlog).

## License

MIT
