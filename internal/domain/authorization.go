package domain

type AuthorizationTransferStatus string

const (
	AuthorizationTransferStatusPending AuthorizationTransferStatus = "PENDING"
	AuthorizationTransferStatusOK      AuthorizationTransferStatus = "OK"
	AuthorizationTransferNegative      AuthorizationTransferStatus = "NEGATIVE"
	AuthorizationTransferStatusError   AuthorizationTransferStatus = "ERROR"
)

type AuthorizationTransfer struct {
	ID         ID                          `json:"id" validate:"required,uuid"`
	TransferID ID                          `json:"transfer_id" validate:"required,uuid"`
	Status     AuthorizationTransferStatus `json:"status" validate:"required,oneof=PENDING OK NEGATIVE ERROR"`
	Enabled    bool                        `json:"enabled" validate:"required,bool"`
	CreatedAt  CreatedAt                   `json:"created_at"`
	UpdatedAt  UpdatedAt                   `json:"updated_at"`
	Version    Version                     `json:"version" validate:"required,numeric"`
}
