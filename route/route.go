/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-03-28 14:22:24
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-08-20 14:31:26
 * @FilePath: /feishu-bot/route/route.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package route

import (
	"net/http"
	v1 "qingliuBot/route/v1"

	"github.com/gorilla/mux"
)

func InitRoute() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/v1").Subrouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	// handle chat create callback
	// service.LarkCli.EventCallback.HandlerEventV1P2PChatCreate(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV1, event *lark.EventV1P2PChatCreate) (string, error) {
	// 	_, _, err := cli.Message.Send().ToChatID(event.ChatID).SendText(ctx, "hi")
	// 	return "", err
	// })

	// handle message callback
	// service.LarkCli.EventCallback.HandlerEventV2IMMessageReceiveV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMMessageReceiveV1) (string, error) {
	// 	content, err := lark.UnwrapMessageContent(event.Message.MessageType, event.Message.Content)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	switch {
	// 	case content.Text != nil:
	// 		_, _, err = cli.Message.Reply(event.Message.MessageID).SendText(ctx, fmt.Sprintf("got text: %s", content.Text.Text))
	// 	case content.File != nil:
	// 		_, _, err = cli.Message.Reply(event.Message.MessageID).SendText(ctx, fmt.Sprintf("got file: %s, key: %s", content.File.FileName, content.File.FileKey))
	// 	case content.Image != nil:
	// 		_, _, err = cli.Message.Reply(event.Message.MessageID).SendText(ctx, fmt.Sprintf("got image: %s", content.Image.ImageKey))
	// 	}
	// 	return "", err
	// })

	// handle add bot
	// service.LarkCli.EventCallback.HandlerEventV2IMChatMemberBotAddedV1(func(ctx context.Context, cli *lark.Lark, schema string, header *lark.EventHeaderV2, event *lark.EventV2IMChatMemberBotAddedV1) (string, error) {
	// 	return "", nil
	// })

	// s.HandleFunc("/api", v1.FeiShuEvent)
	// s.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	// 	//获取机器人所在列表
	// 	rest, _, err := service.LarkCli.Chat.GetChatListOfSelf(context.Background(), &lark.GetChatListOfSelfReq{})
	// 	fmt.Println(rest, err)

	// card := &lark.MessageContentCard{
	// 	Header: &lark.MessageContentCardHeader{Template: lark.MessageContentCardHeaderTemplateOrange, Title: &lark.MessageContentCardObjectText{Content: "title"}},
	// 	Config: nil,
	// 	Modules: []lark.MessageContentCardModule{
	// 		&lark.MessageContentCardModuleDIV{
	// 			Text: &lark.MessageContentCardObjectText{
	// 				Tag:     lark.MessageContentCardTextTypeLarkMd,
	// 				Content: "这是内容",
	// 				Lines:   1,
	// 			},
	// 			Fields: nil,
	// 			Extra:  lark.MessageContentCardElementButton{URL: "https://baidu.com", Type: "danger", Text: &lark.MessageContentCardObjectText{Content: "open", Tag: lark.MessageContentCardTextTypePlainText}},
	// 		},
	// 	},
	// }
	// fmt.Println(card)
	// 	a, b, c := service.LarkCli.Message.Send().ToChatID("oc_3b7ab4655bd6de673cf30cce8ce58b19").SendText(context.Background(), "发送测试内容")
	// 	fmt.Println(a, b, c)
	// 	json.NewEncoder(w).Encode(rest)
	// })
	s.HandleFunc("/4px", v1.GetSign)
	s.HandleFunc("/4px/import", v1.Import)
	s.HandleFunc("/4px/cookie", v1.Cookie)
	// s.HandleFunc("/px/login", v1.Login)
	// githubGroup := s.PathPrefix("/github").Subrouter()
	// githubGroup.HandleFunc("/payload", v1.GithubHook)
	// packageGroup := s.PathPrefix("/package").Subrouter()
	// packageGroup.HandleFunc("/test", v1.Test)
	// packageGroup.HandleFunc("/import", v1.PostVersion).Methods("POST")
	// packageGroup.HandleFunc("/", v1.GetMainPackage).Methods("GET")
	// packageGroup.HandleFunc("/migration", v1.GetUpgradePackage).Methods("GET")
	// This will serve files under http://localhost:8000/static/<filename>
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	return r
}
