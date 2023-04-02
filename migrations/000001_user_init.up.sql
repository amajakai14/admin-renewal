CREATE TABLE IF NOT EXISTS app_user (
    id SERIAL NOT NULL,
    name TEXT,
    email TEXT,
    password TEXT,
    email_verified BOOLEAN,
    role TEXT,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    corporation_id TEXT,

    CONSTRAINT user_pkey PRIMARY KEY (id)
);

