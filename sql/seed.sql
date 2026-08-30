CREATE TABLE Games (
    guild_id TEXT,
    creator_id TEXT,
    created_at BIGINT,
    imposters TEXT[],
    innocents TEXT[],
    PRIMARY KEY (guild_id, creator_id)
);