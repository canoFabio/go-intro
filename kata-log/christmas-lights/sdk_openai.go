package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	openai "github.com/sashabaranov/go-openai"
)

func main() {

	config := openai.DefaultConfig("")
	proxyUrl, err := url.Parse("https://openai-proxy.melioffice.com/v1/")
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}

	c := openai.NewClientWithConfig(config)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:       openai.GPT3Ada,
		MaxTokens:   60,
		Temperature: 0,
		TopP:        1.0,
		Prompt:      "Lorem ipsum",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
