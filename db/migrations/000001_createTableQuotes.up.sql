CREATE TABLE IF NOT EXISTS "quotes"
(
    "ID"     uuid PRIMARY KEY default gen_random_uuid(),
    "author" varchar(255) NOT NULL,
    "text"   varchar(255) NOT NULL
)