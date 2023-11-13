package gpt

import (
	"context"
	"dingchatgpt/config"
	"dingchatgpt/db"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func AskChatGpt(question, askUser string, askContext []db.GptInfo) (string, error) {

	appConfig := config.GetConfig()

	//gptToken := appConfig.GptToken

	gptConfig := openai.DefaultConfig(appConfig.GptToken)

	gptConfig.BaseURL = appConfig.GptBaseUrl

	//gptConfig.HTTPClient = &http.Client{}

	//gptConfig.HTTPClient.Transport.RoundTrip()

	c := openai.NewClientWithConfig(gptConfig)

	var askGptMessages []openai.ChatCompletionMessage

	// 判断是否有上下文提问
	if len(askContext) > 0 {
		for _, info := range askContext {
			askGptMessages = append(askGptMessages, []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: info.BotAnswer,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: info.AskBotQuestion,
				},
			}...)
		}
	}

	// 添加现在需要的提问
	askGptMessages = append(askGptMessages, []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: question,
		},
	}...)

	//log.Println(askGptMessages)

	resp, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: askGptMessages,
		})

	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return "", nil
	}
	//fmt.Println(resp.Choices[0].Message.Content)
	answer := resp.Choices[0].Message.Content
	fmt.Println(answer)
	timeStamp := int(resp.Created)
	askUseTokens := resp.Usage.PromptTokens
	answerUseTokens := resp.Usage.CompletionTokens
	//fmt.Println(askUseTokens, answerUseTokens, timeStamp)

	d := db.GptDBModel()
	//d.Init()
	err = d.InsertMessage(askUser, question, answer, timeStamp, askUseTokens+answerUseTokens)
	if err != nil {
		return "", err
	}

	//ding.AtUserSimpleReplyMarkdown(context.Background(), sessionWebhook, []byte("default"), []byte(answer), []string{askUser})

	return resp.Choices[0].Message.Content, nil
}
