package generator

import (
	"archive/zip"
	"context"
	"encoding/json"
	"fmt"
	"github.com/5GCoreNet/openapi-generator-go/pkg/utils"
	"github.com/google/go-github/github"
	"io"
	"net/http"
	"strings"
)

type Generator struct {
	// Repository is the URL to the OpenAPI specification repository
	Repository string
	// RepositoryOwner is the owner of the OpenAPI specification repository
	RepositoryOwner string
	// Ref is the commit hash or branch name of the OpenAPI specification repository
	Ref string
	// Path is the path to the OpenAPI specification folder under the repository
	Path string
	// Output is the output directory for the generated code
	Output string
	// ExitOnFailure is a flag to exit on error
	ExitOnFailure bool
}

type input struct {
	SwaggerUrl string            `json:"openAPIUrl"`
	Options    map[string]string `json:"options"`
}

type output struct {
	Link string `json:"link"`
}

type Mode string

const (
	// ClientMode generates a client for the API
	ClientMode Mode = "client"
	// ServerMode generates a server for the API
	ServerMode Mode = "server"
)

// Validate validates the generator configuration
func (g *Generator) Validate() error {
	if g.Output == "" {
		return fmt.Errorf("output is required")
	}
	return nil
}

func (g *Generator) Generate(ctx context.Context, mode Mode) error {
	var generatorUrl string
	var subFolder string
	switch mode {
	case ClientMode:
		generatorUrl = utils.OpenApiClientGeneratorUrl
		subFolder = utils.ClientSubFolder
	case ServerMode:
		generatorUrl = utils.OpenApiGinServerGeneratorUrl
		subFolder = utils.ServerGinSubFolder
	default:
		return fmt.Errorf("unknown mode: %s", mode)
	}
	urls, err := g.getSpecFromUrl(ctx)
	if err != nil {
		return err
	}
	for _, url := range urls {
		err := g.generate(ctx, url, generatorUrl, subFolder)
		if err != nil && g.ExitOnFailure {
			return err
		}
	}
	return nil
}

func (g *Generator) getSpecFromUrl(ctx context.Context) ([]string, error) {
	client := github.NewClient(nil)
	_, files, _, err := client.Repositories.GetContents(
		ctx,
		g.RepositoryOwner,
		g.Repository,
		g.Path,
		&github.RepositoryContentGetOptions{
			Ref: g.Ref,
		},
	)
	if err != nil {
		return nil, err
	}
	var urls []string
	for _, file := range files {
		if strings.HasSuffix(file.GetName(), ".yaml") || strings.HasSuffix(file.GetName(), ".yml") {
			urls = append(urls, retrieveUrl(file.GetHTMLURL()))
		}
	}
	return urls, nil
}

func retrieveUrl(url string) string {
	return strings.Replace(
		strings.Replace(
			url,
			"github.com",
			"raw.githubusercontent.com",
			1,
		),
		"/blob",
		"",
		1,
	)
}

func (g *Generator) generate(ctx context.Context, url string, generatorUrl string, subFolder string) error {
	specName := utils.GetSpecNameFromUrl(url)
	body, err := json.Marshal(input{
		SwaggerUrl: url,
		Options: map[string]string{
			"packageName":   "openapi_" + specName,
			"isGoSubmodule": "true",
		},
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, generatorUrl, strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var output output
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, &output); err != nil {
		return err
	}
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, output.Link, nil)
	if err != nil {
		return err
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	reader, err := zip.NewReader(strings.NewReader(string(b)), int64(len(b)))
	if err != nil {
		return err
	}
	for _, file := range reader.File {
		// Skip go.mod and go.sum files to avoid conflicts with the main project
		if strings.HasSuffix(file.Name, ".mod") || strings.HasSuffix(file.Name, ".sum") {
			continue
		}
		if err := utils.SaveFile(file, g.Output, specName, subFolder); err != nil {
			return err
		}
	}
	return nil
}
