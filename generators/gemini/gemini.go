package gemini

import (
	"context"
	"errors"
	"thunderbird-ai-compose-server/config"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateResponse(prompt string) (string, error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(config.Config.APIKey))
	if err != nil {
		return "", err
	}

	defer client.Close()

	model := client.GenerativeModel(config.Config.Model)
	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockNone,
		},
		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockNone,
		},
	}

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 {
		return "", errors.New("no response candidates received")
	}

	candidate := resp.Candidates[0]

	if candidate.Content == nil {
		return "", errors.New("no content in the response candidate")
	}

	if len(candidate.Content.Parts) == 0 {
		return "", errors.New("no parts in the response content")
	}

	if textPart, ok := candidate.Content.Parts[0].(genai.Text); ok {
		result := string(textPart)
		return result, nil
	}

	return "", errors.New("response part is not text")
}
