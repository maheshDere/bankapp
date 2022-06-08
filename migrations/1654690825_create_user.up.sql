CREATE TABLE "user" (
  "id" varchar(64) PRIMARY KEY,
  "name" varchar(200) NOT NULL DEFAULT '',    
  "email" varchar(512) UNIQUE NOT NULL DEFAULT '',
  "password" varchar(200),
  "role_type" int, 
  "created_at" timestamp,
  "updated_at" timestamp
);