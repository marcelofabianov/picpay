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
	ID            ID
	PayerID       ID
	PayerWalletID ID
	PayeeID       ID
	PayeeWalletID ID
	Amount        Amount
	Status        TransferStatus
	Enabled       bool
	CreatedAt     CreatedAt
	UpdatedAt     UpdatedAt
	Version       Version
}
