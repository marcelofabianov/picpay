-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS authorization_transfers (
    id UUID PRIMARY KEY,
    transfer_id UUID NOT NULL,
    status VARCHAR(20) NOT NULL,
    enabled BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    version INT DEFAULT 1 NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_authorization_transfers_transfer_id ON authorization_transfers (transfer_id);

CREATE INDEX IF NOT EXISTS idx_authorization_transfers_status ON authorization_transfers (status);

CREATE INDEX IF NOT EXISTS idx_authorization_transfers_enabled ON authorization_transfers (enabled);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_authorization_transfers_enabled;

DROP INDEX IF EXISTS idx_authorization_transfers_status;

DROP INDEX IF EXISTS idx_authorization_transfers_transfer_id;

DROP TABLE IF EXISTS authorization_transfers;
-- +goose StatementEnd
