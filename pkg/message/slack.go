package message

import (
	"fmt"

	"github.com/ionut-maxim/notarry/pkg/radarr"
	"github.com/ionut-maxim/notarry/pkg/templates"
	"github.com/slack-go/slack"
)

func SendSlackMessage(webhook string, template string) error {
	r, err := radarr.New()
	if err != nil {
		return err
	}

	blocks, err := templates.NewSlackTemplate(r, template)
	if err != nil {
		return err
	}

	err = slack.PostWebhook(webhook, &slack.WebhookMessage{
		Text: fmt.Sprintf("%s (%s) %s", r.Movie.Title, r.Movie.Year, r.FormattedType),
		Blocks: &slack.Blocks{
			BlockSet: blocks.BlockSet,
		},
	})
	if err != nil {
		return err
	}

	return nil
}
