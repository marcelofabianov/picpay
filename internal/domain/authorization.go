package domain

type AuthorizationTransferStatus string

const (
	AuthorizationTransferStatusPending AuthorizationTransferStatus = "PENDING"
	AuthorizationTransferStatusOK      AuthorizationTransferStatus = "OK"
	AuthorizationTransferRejected      AuthorizationTransferStatus = "REJECTED"
	AuthorizationTransferStatusError   AuthorizationTransferStatus = "ERROR"
)

type AuthorizationTransfer struct {
	ID         ID
	TransferID ID
	Status     AuthorizationTransferStatus
	Enabled    bool
	CreatedAt  CreatedAt
	UpdatedAt  UpdatedAt
	Version    Version
}
