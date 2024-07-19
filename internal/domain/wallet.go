package domain

type WalletType string

const (
	WalletTypeCommon   WalletType = "user"
	WalletTypeMerchant WalletType = "merchant"
)

type Wallet struct {
	ID        ID         `json:"id" validate:"required,uuid"`
	UserID    ID         `json:"user_id" validate:"required,uuid"`
	Amount    Amount     `json:"amount" validate:"required,numeric"`
	Type      WalletType `json:"type" validate:"required,oneof=user merchant"`
	Enabled   bool       `json:"enabled" validate:"required,bool"`
	CreatedAt CreatedAt  `json:"created_at"`
	UpdatedAt UpdatedAt  `json:"updated_at"`
	DeletedAt *DeletedAt `json:"deleted_at" validate:"omitempty"`
	Version   Version    `json:"version" validate:"required,numeric"`
}
