-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS wallets (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    amount NUMERIC(15, 2) NOT NULL CHECK (amount >= 0.01),
    enabled BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    version INT DEFAULT 1 NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_wallets_user_id ON wallets (user_id);

CREATE INDEX IF NOT EXISTS idx_wallets_enabled ON wallets (enabled);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_wallets_enabled;

DROP INDEX IF EXISTS idx_wallets_user_id;

DROP TABLE IF EXISTS wallets;
-- +goose StatementEnd
