package papercli

import "time"

type User struct {
	ID            string    `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	HomeFolderID  string    `json:"home_folder_id"`
	InboxFolderID string    `json:"inbox_folder_id"`
}
