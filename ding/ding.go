package ding

import (
	"context"
	"dingchatgpt/config"
	"dingchatgpt/db"
	"dingchatgpt/gpt"
	"fmt"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/chatbot"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/client"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/logger"
	"strings"
)

// 添加@用户的回调机器人markdown方法
func AtUserSimpleReplyMarkdown(ctx context.Context, sessionWebhook string, title, content []byte, userIds []string) error {
	requestBody := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": string(title),
			"text":  string(content),
		},
		"at": map[string]interface{}{
			"atUserIds": userIds,
		},
	}
	return chatbot.NewChatbotReplier().ReplyMessage(ctx, sessionWebhook, requestBody)
}

func OnChatBotMessageReceived(ctx context.Context, data *chatbot.BotCallbackDataModel) ([]byte, error) {
	//replyMsg := []byte(fmt.Sprintf("msg received: [%s]", data.Text.Content))
	//marshal, _ := json.Marshal(data)
	//
	//logger.GetLogger().Infof(string(marshal) + "\r\n")

	senderStaffId := data.SenderStaffId

	question := strings.TrimSpace(data.Text.Content)

	logger.GetLogger().Infof(data.Text.Content + "\r\n")

	//askChan := make(chan string, 5)

	// 查询用户最后5次提问

	//db.GptDBModel().Init()
	ask, _ := db.GptDBModel().SelectLast5Ask(senderStaffId)

	// 附带上下文提问chat-gpt
	gptAnswer, err := gpt.AskChatGpt(question, senderStaffId, ask)

	// 机器人回答信息
	answer := fmt.Sprintf("@%s \r\n %s", senderStaffId, gptAnswer)

	if err != nil {
		return []byte(""), err
	}

	// 发送gpt回答，并@用户
	AtUserSimpleReplyMarkdown(context.Background(), data.SessionWebhook, []byte("gpt"), []byte(answer), []string{senderStaffId})

	//db.GptDBModel().InsertMessage(senderStaffId, question, gptAnswer)

	return []byte(""), nil
}

// 启动钉钉机器人stream 监听消息
func StartDingTalk() {
	//var clientId, clientSecret string
	appConfig := config.GetConfig()

	clientId := appConfig.DingAppKey
	clientSecret := appConfig.DingAppSecret

	//
	//flag.Parse()
	logger.SetLogger(logger.NewStdTestLogger())

	cli := client.NewStreamClient(client.WithAppCredential(client.NewAppCredentialConfig(clientId, clientSecret)))

	//注册事件类型的处理函数
	//cli.RegisterAllEventRouter(OnEventReceived)
	//注册callback类型的处理函数
	cli.RegisterChatBotCallbackRouter(OnChatBotMessageReceived)
	//注册插件的处理函数
	//cli.RegisterPluginCallbackRouter(OnPluginMessageReceived)

	//db.GptDBModel().Init()

	err := cli.Start(context.Background())
	if err != nil {
		panic(err)
	}

	defer cli.Close()
}
