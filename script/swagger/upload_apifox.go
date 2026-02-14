package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	projectID   = "5185764"
	accessToken = "afxp_d61b1beVqkVIcdsyg4nM8rIYwRzR16tBCk2a"
	apiVersion  = "2024-03-28"
	baseURL     = "https://api.apifox.com"
)

type ImportRequest struct {
	Input   string      `json:"input"`
	Options interface{} `json:"options,omitempty"`
}

func main() {

	files, err := filepath.Glob("./docs/swagger/*.json")
	if err != nil {
		panic(err)
	}

	if len(files) == 0 {
		fmt.Println("No swagger json files found")
		return
	}

	for _, file := range files {
		fmt.Println("Uploading:", file)

		err := uploadSwagger(file)
		if err != nil {
			fmt.Println("❌ Failed:", err)
			continue
		}

		fmt.Println("✅ Success:", file)
	}
}

func uploadSwagger(filePath string) error {

	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	reqBody := ImportRequest{
		Input: string(content),
		Options: map[string]interface{}{
			"endpointOverwriteBehavior": "OVERWRITE_EXISTING",
			"schemaOverwriteBehavior":   "OVERWRITE_EXISTING",
			"deleteUnmatchedResources":  false,
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/v1/projects/%s/import-openapi", baseURL, projectID)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Set("X-Apifox-Api-Version", apiVersion)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 300 {
		return fmt.Errorf("status: %d, body: %s", resp.StatusCode, string(body))
	}

	fmt.Println("upload successfully")
	return nil
}
