CREATE TABLE chats
(
    id           UUID PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    is_group     BOOLEAN      NOT NULL DEFAULT FALSE,
    employee_ids UUID[] NOT NULL,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE employee_chat_settings
(
    id           UUID PRIMARY KEY,
    chat_id      UUID        NOT NULL REFERENCES chats (id) ON DELETE CASCADE,
    employee_id  UUID        NOT NULL,
    display_name VARCHAR(255),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (chat_id, employee_id)
);

CREATE TABLE messages
(
    id           UUID PRIMARY KEY,
    chat_id      UUID        NOT NULL REFERENCES chats (id) ON DELETE CASCADE,
    employee_id  UUID        NOT NULL,
    colleague_id UUID,
    text         TEXT        NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);