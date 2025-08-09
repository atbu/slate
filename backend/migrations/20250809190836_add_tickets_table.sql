-- +goose Up
-- +goose StatementBegin
CREATE TYPE ticket_state AS ENUM (
    'New',
    'Proposed',
    'Active',
    'In Development',
    'Developed',
    'Review',
    'In Test',
    'Tested',
    'Closed'
);

CREATE TABLE IF NOT EXISTS tickets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    current_state TICKET_STATE NOT NULL DEFAULT 'New',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by UUID NOT NULL REFERENCES users(id),
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by UUID NOT NULL REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tickets;
DROP TYPE IF EXISTS ticket_state;
-- +goose StatementEnd
