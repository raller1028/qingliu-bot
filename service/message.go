package service

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
)

type LarkService interface {
	SendCard(chartId string, title, content, url string)
}

type larkService struct {
}

func D(cli *lark.Lark) {
	// handle chat create callback
	cli.EventCallback.HandlerEventV1P2PChatCreate(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV1, event *lark.EventV1P2PChatCreate) (string, error) {
		_, _, err := cli.Message.Send().ToChatID(event.ChatID).SendText(ctx, "hi")
		return "", err
	})

	// handle message callback
	cli.EventCallback.HandlerEventV2IMMessageReceiveV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMMessageReceiveV1) (string, error) {
		content, err := lark.UnwrapMessageContent(event.Message.MessageType, event.Message.Content)
		if err != nil {
			return "", err
		}
		switch {
		case content.Text != nil:
			_, _, err = cli.Message.Reply(event.Message.MessageID).SendText(ctx, fmt.Sprintf("got text: %s", content.Text.Text))
		case content.File != nil:
			_, _, err = cli.Message.Reply(event.Message.MessageID).SendText(ctx, fmt.Sprintf("got file: %s, key: %s", content.File.FileName, content.File.FileKey))
		case content.Image != nil:
			_, _, err = cli.Message.Reply(event.Message.MessageID).SendText(ctx, fmt.Sprintf("got image: %s", content.Image.ImageKey))
		}
		return "", err
	})

	// handle add bot
	cli.EventCallback.HandlerEventV2IMChatMemberBotAddedV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMChatMemberBotAddedV1) (string, error) {
		return "", nil
	})
}

//oc_5db66c3947d8f5720d219e5b2a613c65
func (l *larkService) SendCard(chartId string, title, content, url string) {
	card := &lark.MessageContentCard{
		Header: &lark.MessageContentCardHeader{Template: lark.MessageContentCardHeaderTemplateOrange, Title: &lark.MessageContentCardObjectText{Content: title}},
		Config: nil,
		Modules: []lark.MessageContentCardModule{
			&lark.MessageContentCardModuleDIV{
				Text: &lark.MessageContentCardObjectText{
					Tag:     lark.MessageContentCardTextTypeLarkMd,
					Content: content,
					Lines:   0,
				},
				Fields: nil,
				Extra:  lark.MessageContentCardElementButton{URL: url, Type: "danger", Text: &lark.MessageContentCardObjectText{Content: "open", Tag: lark.MessageContentCardTextTypePlainText}},
			},
		},
	}

	a, b, c := LarkCli.Message.Send().ToChatID(chartId).SendCard(context.Background(), card.String())
	fmt.Println(a, b, c)
}
func NewLarkService() LarkService {
	return &larkService{}
}
