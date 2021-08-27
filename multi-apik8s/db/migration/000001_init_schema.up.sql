CREATE TABLE "fibs" (
    "id" bigserial PRIMARY KEY,
    "index" bigint UNIQUE NOT NULL,
    "value"  bigint NOT NULL,
    "create_at" timestamptz NOT NULL DEFAULT (now())
)