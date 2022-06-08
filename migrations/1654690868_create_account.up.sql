CREATE TABLE "account" (
  "id" varchar(64) PRIMARY KEY,
  "opening_date" timestamp,
  "user_id" varchar(64),
  "created_at" timestamp,
  CONSTRAINT "fk_user" FOREIGN KEY("user_id") REFERENCES "user"("id") ON DELETE CASCADE
);