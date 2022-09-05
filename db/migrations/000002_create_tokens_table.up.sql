CREATE TABLE IF NOT EXISTS tokens(
    hashed_token bytea PRIMARY KEY,
    user_id bigint NOT NULL,
    scope text NOT NULL,
    expiry timestamp(0) with time zone NOT NULL
);

ALTER TABLE "tokens"
ADD CONSTRAINT "tokens_users_fk"
FOREIGN KEY ("user_id")
REFERENCES "users" ("id")
ON DELETE CASCADE;