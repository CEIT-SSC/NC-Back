CREATE TABLE "users"
(
    "id"             SERIAL PRIMARY KEY,
    "username"       varchar,
    "password"       varchar,
    "student_number" int
);

ALTER TABLE users
    OWNER TO test_user;

CREATE TABLE "rooms"
(
    "user_id"     int,
    "room_title"  varchar,
    "room_state"  jsonb,
    "is_complete" boolean,
    "score"       int,
    PRIMARY KEY ("user_id", "room_title")
);

ALTER TABLE rooms
    OWNER TO test_user;

CREATE TABLE "user_tokens"(
    "user_id" int,
    "token" varchar,
    PRIMARY KEY("user_id","token")
)