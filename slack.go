package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

type slackClient struct {
	client  *slack.Client
	channel string
}

func newSlackClient(channel string) *slackClient {
	client := slack.New(os.Getenv("SLACK_APP_TOKEN"))
	return &slackClient{client: client, channel: channel}
}

func setFieldsBy(t *totalTable) []slack.AttachmentField {
	var fields []slack.AttachmentField

	fields = append(fields, slack.AttachmentField{
		Title: "収入",
		Value: t.Income,
		Short: true,
	})
	fields = append(fields, slack.AttachmentField{
		Title: "支出",
		Value: t.Expenses,
		Short: true,
	})
	fields = append(fields, slack.AttachmentField{
		Title: "収支",
		Value: t.Balance,
		Short: true,
	})

	return fields
}

func (s *slackClient) postMessage(t *totalTable) {
	attachment := slack.Attachment{
		Color:   "good",
		Pretext: "MoneyForwardの取得結果です",
		Fields:  setFieldsBy(t),
	}
	if _, _, err := s.client.PostMessage(s.channel, slack.MsgOptionAttachments(attachment)); err != nil {
		fmt.Println(err)
	}

	return
}
