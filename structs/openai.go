package structs

import (
	"context"
	"log"

	"github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	APIKey string
	client *openai.Client
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	client := openai.NewClient(apiKey)

	return &OpenAIClient{
		APIKey: apiKey,
		client: client,
	}
}

func (c *OpenAIClient) GetEmbeddingForText(text string) openai.Embedding {
	queryReq := openai.EmbeddingRequest{
		Input: []string{text},
		Model: openai.AdaEmbeddingV2,
	}
	queryResponse, err := c.client.CreateEmbeddings(context.Background(), queryReq)
	if err != nil {
		log.Fatal("Error creating query embedding:", err)
	}
	return queryResponse.Data[0]
}
