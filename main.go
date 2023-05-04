package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"golang.org/x/net/context"
	"net/http"
)

func main() {
	// 创建一个服务
	r := gin.Default()
	// 访问地址 处理请求 Request Response
	r.GET("/api", func(c *gin.Context) {
		// c.JSON(200, gin.H{"msg": "喜喜"})
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	r.POST("/chatgpt", func(c *gin.Context) {
		msg := c.Query("msg")

		client := openai.NewClient("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")
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: msg,
					},
				},
			},
		)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}
		fmt.Println(resp.Choices[0].Message.Content)

		c.JSON(200, gin.H{"returnMessage": resp.Choices[0].Message.Content})
	})
	// 服务器端口
	r.Run(":8012") // 默认8080
}
