package domain

type TransferStatus string

const (
	TransferStatusPending   TransferStatus = "PENDING"
	TransferStatusRejected  TransferStatus = "REJECTED"
	TransferStatusReserved  TransferStatus = "RESERVED"
	TransferStatusCompleted TransferStatus = "COMPLETED"
	TransferStatusReversed  TransferStatus = "REVERSED"
	TransferStatusError     TransferStatus = "ERROR"
)

type Transfer struct {
	ID            ID             `json:"id" validate:"required,uuid"`
	PayerID       ID             `json:"payer_id" validate:"required,uuid"`
	PayerWalletID ID             `json:"payer_wallet_id" validate:"required,uuid"`
	PayeeID       ID             `json:"payee_id" validate:"required,uuid"`
	PayeeWalletID ID             `json:"payee_wallet_id" validate:"required,uuid"`
	Amount        Amount         `json:"amount" validate:"required,numeric"`
	Status        TransferStatus `json:"status" validate:"required,oneof=PENDING REFUSED RESERVED COMPLETED REVERSED ERROR"`
	Enabled       bool           `json:"enabled" validate:"required,bool"`
	CreatedAt     CreatedAt      `json:"created_at"`
	UpdatedAt     UpdatedAt      `json:"updated_at"`
	Version       Version        `json:"version" validate:"required,numeric"`
}
