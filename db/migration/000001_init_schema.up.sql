CREATE TABLE "hotel" (
                         "id" bigserial PRIMARY KEY,
                         "name" varchar NOT NULL,
                         "address" varchar NOT NULL,
                         "location" varchar NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "room" (
                        "id" bigserial PRIMARY KEY,
                        "name" varchar NOT NULL,
                        "room_type_id" bigserial NOT NULL,
                        "hotel_id" bigserial NOT NULL,
                        "is_available" bigint NOT NULL,
                        "status" bigint NOT NULL,
                        "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "room_inventory" (
                                  "hotel_id" bigint NOT NULL,
                                  "room_id" bigserial NOT NULL,
                                  "room_type_id" bigserial NOT NULL,
                                  "date" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
                                  "total_inventory" int NOT NULL,
                                  "total_reserved" int NOT NULL,
                                  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "rate" (
                        "hotel_id" bigint,
                        "room_id" bigserial NOT NULL,
                        "rate" int NOT NULL,
                        "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "reservation" (
                               "id" bigserial PRIMARY KEY,
                               "hotel_id" bigint,
                               "room_id" bigserial NOT NULL,
                               "start_date" timestamptz NOT NULL,
                               "end_date" timestamptz NOT NULL,
                               "status" int NOT NULL,
                               "user_id" bigserial,
                               "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "username" varchar NOT NULL,
                         "hashed_password" varchar NOT NULL,
                         "full_name" varchar NOT NULL,
                         "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "room_type" (
                             "id" bigserial PRIMARY KEY,
                             "name" varchar NOT NULL,
                             "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "room_inventory" ("hotel_id", "room_type_id");

CREATE UNIQUE INDEX ON "room_inventory" ("hotel_id", "room_id");

CREATE UNIQUE INDEX ON "rate" ("hotel_id", "room_id");

ALTER TABLE "room" ADD FOREIGN KEY ("room_type_id") REFERENCES "room_type" ("id");

ALTER TABLE "room" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotel" ("id");

ALTER TABLE "room_inventory" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotel" ("id");

ALTER TABLE "room_inventory" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("id");

ALTER TABLE "rate" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotel" ("id");

ALTER TABLE "rate" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("id");

ALTER TABLE "reservation" ADD FOREIGN KEY ("hotel_id") REFERENCES "hotel" ("id");

ALTER TABLE "reservation" ADD FOREIGN KEY ("room_id") REFERENCES "room" ("id");

ALTER TABLE "reservation" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
