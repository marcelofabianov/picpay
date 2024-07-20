package presenter

type UserPresenter struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	DocumentRegistry string `json:"document_registry"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	Enabled          bool   `json:"enabled"`
}
