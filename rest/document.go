package papercli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
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

	var doc = CreateDocument{
		Title:    filepath.Base(*file_path),
		FileName: filepath.Base(*file_path),
		ParentID: *parent_id,
	}

	create_document(host, token, &doc)
}

func create_document(host *string, token *string, doc *CreateDocument) {

	jsonData, err := json.Marshal(*doc)
	if err != nil {
		log.Fatal(err)
	}

	url := fmt.Sprintf("%s/api/nodes/", *host)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *token))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Body: %s", body)

}

func upload_file() {

}
