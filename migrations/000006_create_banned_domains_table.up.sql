CREATE TABLE "banned_domains"(
    "id" UUID NOT NULL,
    "url" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
ALTER TABLE
    "banned_domains" ADD PRIMARY KEY("id");