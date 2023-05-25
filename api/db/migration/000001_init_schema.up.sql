CREATE TABLE "employees" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "performance" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX on "employees" ("name");
