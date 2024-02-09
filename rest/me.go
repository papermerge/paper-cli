package papercli

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	types "paper-cli/types"
)

func Me(host *string, token *string) *types.User {
	url := fmt.Sprintf("%s/api/users/me", *host)
	req, err := http.NewRequest("GET", url, nil)

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
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)

	return &user
}
