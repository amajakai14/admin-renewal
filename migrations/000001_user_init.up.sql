CREATE TABLE IF NOT EXISTS app_user (
    id TEXT NOT NULL,
    name TEXT,
    email TEXT,
    password TEXT,
	salt TEXT,
    email_verified TIMESTAMP(3),
    image TEXT,
    role TEXT,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    corporation_id TEXT,

    CONSTRAINT user_pkey PRIMARY KEY (id)
);

