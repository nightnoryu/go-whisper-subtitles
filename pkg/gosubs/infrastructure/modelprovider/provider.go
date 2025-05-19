package modelprovider

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/schollz/progressbar/v3"
)

func NewModelProvider(modelsDir string) *ModelProvider {
	return &ModelProvider{
		modelsDir: modelsDir,
	}
}

type ModelProvider struct {
	modelsDir string
}

const (
	modelURLTemplate   = "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-%s.bin?download=true"
	outputPathTemplate = "ggml-%s.bin"
)

func (p *ModelProvider) DownloadModel(model string) (modelPath string, err error) {
	outputPath := filepath.Join(p.modelsDir, fmt.Sprintf(outputPathTemplate, model))
	if _, err := os.Stat(outputPath); err == nil {
		fmt.Printf("model %s found\n", model)
		return outputPath, nil
	}

	fmt.Printf("model %s not found locally, downloading\n", model)

	url := fmt.Sprintf(modelURLTemplate, model)

	output, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download model %s: %w", model, err)
	}
	defer response.Body.Close()

	bar := progressbar.NewOptions64(
		response.ContentLength,
		progressbar.OptionSetDescription("Downloading model..."),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowBytes(true),
		progressbar.OptionShowTotalBytes(true),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	_, err = io.Copy(io.MultiWriter(output, bar), response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save model: %w", err)
	}

	return outputPath, nil
}
