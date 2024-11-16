CREATE TABLE Chats
(
    id           UUID PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    is_group     BOOLEAN      NOT NULL DEFAULT FALSE,
    employee_ids UUID[]       NOT NULL DEFAULT ARRAY[]::UUID[],
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE EmployeeChatSettings
(
    id           UUID PRIMARY KEY,
    chat_id      UUID        NOT NULL REFERENCES Chats (id) ON DELETE CASCADE,
    employee_id  UUID        NOT NULL,
    display_name VARCHAR(255),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (chat_id, employee_id)
);

CREATE TABLE Messages
(
    id           UUID PRIMARY KEY,
    chat_id      UUID        NOT NULL REFERENCES chats (id) ON DELETE CASCADE,
    employee_id  UUID        NOT NULL,
    colleague_id UUID        NOT NULL,
    text         TEXT        NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);