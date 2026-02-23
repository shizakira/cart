BEGIN;

CREATE TABLE IF NOT EXISTS carts
(
    user_id    BIGINT PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS cart_items
(
    user_id    BIGINT      NOT NULL,
    sku_id     BIGINT      NOT NULL,
    count      SMALLINT    NOT NULL CHECK (count > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (user_id, sku_id),
    FOREIGN KEY (user_id) REFERENCES carts (user_id) ON DELETE CASCADE
);

COMMIT;