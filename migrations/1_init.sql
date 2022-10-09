-- +migrate Up
CREATE TABLE "users" (
    "user_id" PRIMARY KEY,
    "full_name" text,
    "email" UNIQUE,
    "password" text, 
    "role" text,
    "created_at" TIMESTAMPTZ NOT NULL, 
    "updated_ate" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "repos" (
    "name" PRIMARY KEY,
    "description" text,
    "url" text,
    "color" text,
    "lang" text,
    "fork" text,
    "stars" text,
    "starts_today" text,
    "build_by" text,
    "created_at" TIMESTAMPTZ NOT NULL, 
    "updated_ate" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "bookmarks" (
    "bid" PRIMARY KEY,
    "user_id" text,
    "repo_name" text,
    "created_at" TIMESTAMPTZ NOT NULL, 
    "updated_ate" TIMESTAMPTZ NOT NULL
);

ALTER TABLE "bookmarks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
ALTER TABLE "bookmarks" ADD FOREIGN KEY ("repo_name") REFERENCES "repos" ("name");

-- +migrate Down
DROP TABLE users;
DROP TABLE repos;
DROP TABLE bookmarks;

