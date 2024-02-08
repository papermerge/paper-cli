package papercli

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

func Upload(
	host *string,
	token *string,
	file_path *string,
	parent_id *string) {

	fileBytes, _ := os.ReadFile(*file_path)

	uri := fmt.Sprintf("%s/api/document", *host)

	client := resty.New()
	resp, err := client.R().
		SetBody(fileBytes).
		SetAuthToken(*token).
		Post(uri)

	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("Status Code: ", resp.StatusCode())
}
