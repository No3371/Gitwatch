package main

import (
	// "encoding/json"
	"flag"
	"fmt"

	// "strconv"

	// "fmt"
	// "os"
	// "path/filepath"
	"time"

	"github.com/No3371/gitwatch/gitlog"
	// "./gitlog"
	"go.uber.org/zap"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
)

var logger *zap.SugaredLogger

func init () {
    l, err := zap.NewDevelopment()
    if err != nil {
        panic("failed to create logger")
    }
    logger = l.Sugar()
}

func main() {
    // apiKey := flag.String("api_key", "")
    botToken := flag.String("token", "", "")
    channelId := flag.Uint64("cid", 0, "")
    repo := flag.String("repo", "", "")
    flag.Parse()

    if *repo == "" {
        logger.Error("no repo ")
    }

	s := state.New("Bot " + *botToken)
	s.AddIntents(gateway.IntentGuilds | gateway.IntentGuildMessages)


	var cache string
    timer := time.NewTimer(time.Minute)

    for {
		// New gitlog
		git := gitlog.New(&gitlog.Config{
			Path: *repo,        // default "."
		})
	
		// List git-log
		commits, err := git.Log(nil, &gitlog.Params { AllRefs: true })
		if err != nil {
			logger.Fatalf("failed to git log: %v", err)
		}
	
		collected := make([]*gitlog.Commit, 0)
		// Output
		for _, commit := range commits {
			if cache == "" {
				cache = commit.Hash.Long
				break
			}
			
			if cache == commit.Hash.Long {
				break
			}
			
			logger.Infof(
				"%s %s %s\n",
				commit.Hash.Short,
				commit.Author.Name,
				commit.Subject,
			)
			collected = append(collected, commit)

			if len(collected) > 10 {
				break
			}
		}

		if len(collected) > 0 {
			cache = collected[0].Hash.Long
			err = send(s, discord.ChannelID(*channelId), collected, len(collected))
			if err != nil {
				logger.Fatalf("failed to send: %v", err)
			}
			collected = collected[:1]
		}

        // logger.Info("One minue delay...")
        <-timer.C
        timer.Reset(time.Minute)
		err = git.FetchAll()
		if err != nil {
			logger.Errorf("failed to fetch all: %v", err)
		} else {
			logger.Infof("fetched all.")
		}
    }

}

func send (s *state.State, cId discord.ChannelID, collected []*gitlog.Commit, total int) error {
	embeds := make([]discord.Embed, 0)
	for _, c := range collected {
		embeds = append(embeds, discord.Embed{
			Type:  discord.NormalEmbed,
			Color: discord.DefaultEmbedColor,
			Title: c.Subject,
			Description: c.Body,
			Fields: []discord.EmbedField {
				discord.EmbedField {
					Name: "作者",
					Value: c.Author.Name,
					Inline: true,
				},
				discord.EmbedField {
					Name: "時間",
					Value: c.Author.Date.String(),
					Inline: true,
				},
				discord.EmbedField {
					Name: "HASH",
					Value: c.Hash.Short,
					Inline: true,
				},
			},
			
		})
	}
	data := api.SendMessageData{
		Content:         fmt.Sprintf("發現 %d 個 git 動態：", total),
		Embeds:          embeds,
	}

	_, err := s.SendMessageComplex(cId, data)
	if err != nil {
		logger.Errorf("failed to send: %v", err)
		_, err = s.SendMessageComplex(cId, data)
		if err != nil {
			logger.Errorf("failed to send again, abort: %v", err)
			return err
		}
	}

	return nil
}

// func commitHandler (c *gogit.Commit, end string, collected map[string]*gogit.Commit) error {
    
//     pCount := c.ParentCount()
//     for i := 0; i < pCount; i++ {
//         commitHandler(c.Parent(i))
//     }
//     logger.Infof("%s %s %s", c.Author.Email, c.Author.When, c.CommitMessage)
//     return nil
// }