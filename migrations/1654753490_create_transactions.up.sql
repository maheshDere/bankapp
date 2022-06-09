CREATE TABLE "transactions" (
  "id" varchar(64) PRIMARY KEY,
  "tnx_type" int,
  "amount" float,
  "account_id" varchar(64),
  "created_at" timestamp,
  CONSTRAINT "fk_account" FOREIGN KEY("account_id") REFERENCES "account"("id") ON DELETE CASCADE
);