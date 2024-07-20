package domain

type WalletType string

const (
	WalletTypeCommon   WalletType = "user"
	WalletTypeMerchant WalletType = "merchant"
)

type Wallet struct {
	ID        ID
	UserID    ID
	Amount    Amount
	Type      WalletType
	Enabled   bool
	CreatedAt CreatedAt
	UpdatedAt UpdatedAt
	DeletedAt *DeletedAt
	Version   Version
}
