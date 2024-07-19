-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transfers (
    id UUID PRIMARY KEY,
    payer_id UUID NOT NULL,
    wallet_origin_id UUID NOT NULL,
    payee_id UUID NOT NULL,
    wallet_destiny_id UUID NOT NULL,
    amount NUMERIC(15, 2) NOT NULL CHECK (amount >= 0.01),
    status VARCHAR(20) NOT NULL,
    enabled BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    version INT DEFAULT 1 NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_transfers_payer_id ON transfers (payer_id);

CREATE INDEX IF NOT EXISTS idx_transfers_wallet_origin_id ON transfers (wallet_origin_id);

CREATE INDEX IF NOT EXISTS idx_transfers_payee_id ON transfers (payee_id);

CREATE INDEX IF NOT EXISTS idx_transfers_wallet_destiny_id ON transfers (wallet_destiny_id);

CREATE INDEX IF NOT EXISTS idx_transfers_enabled ON transfers (enabled);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_transfers_enabled;

DROP INDEX IF EXISTS idx_transfers_wallet_destiny_id;

DROP INDEX IF EXISTS idx_transfers_payee_id;

DROP INDEX IF EXISTS idx_transfers_wallet_origin_id;

DROP INDEX IF EXISTS idx_transfers_payer_id;

DROP TABLE IF EXISTS transfers;
-- +goose StatementEnd
