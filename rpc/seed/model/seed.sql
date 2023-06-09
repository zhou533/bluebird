CREATE TABLE IF NOT EXISTS "seed"(
    "id" bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL, 
    "twitter_id" bigint,
    "name" varchar, 
    "screen_name" varchar UNIQUE NOT NULL, 
    "location" varchar, 
    "url" varchar, 
    "description" varchar, 
    "profile_image_url" varchar,
    "latest_tweet_id" bigint,
    "latest_tweet_at" timestamp,
    "create_at" timestamp,
    "update_at" timestamp,
    "status" int,
    PRIMARY KEY("id")
);