package papercli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-resty/resty/v2"
)

type CreateDocument struct {
	Title    string
	FileName string `json: file_name`
	ParentID string `json:parent_id`
}

func Upload(
	host *string,
	token *string,
	file_path *string,
	parent_id *string) {

	var create_doc = CreateDocument{
		Title:    filepath.Base(*file_path),
		FileName: filepath.Base(*file_path),
		ParentID: *parent_id,
	}

	data, err := json.Marshal(create_doc)

	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fileBytes, _ := os.ReadFile(*file_path)

	client := resty.New()
	client.SetBaseURL(*host)

	resp, err := client.R().
		SetBody(data).
		SetAuthToken(*token).
		Post("/api/nodes")

	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("Status Code: ", resp.StatusCode())

	resp2, err2 := client.R().
		SetBody(fileBytes).
		SetAuthToken(*token).
		Post("/api/document")

	if err2 != nil {
		fmt.Println("err:", err2)
	}

	fmt.Println("Status Code: ", resp2.StatusCode())
}
