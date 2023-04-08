-- CreateTable
CREATE TABLE IF NOT EXISTS channel (
    id TEXT NOT NULL,
    table_id INTEGER NOT NULL,
    status TEXT NOT NULL,
    time_start TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    time_end TIMESTAMP(3) NOT NULL,
    course_id INTEGER NOT NULL,

    CONSTRAINT channel_pkey PRIMARY KEY (id)
);

-- CreateTable
CREATE TABLE IF NOT EXISTS  corporation (
    id TEXT NOT NULL,
    name TEXT,

    CONSTRAINT corporation_pkey PRIMARY KEY (id)
);

-- CreateTable
CREATE TABLE IF NOT EXISTS  course (
    id SERIAL NOT NULL,
    course_name TEXT NOT NULL,
    course_price INTEGER NOT NULL,
    course_timelimit INTEGER DEFAULT 90,
    course_priority INTEGER,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3),
    corporation_id TEXT NOT NULL,

    CONSTRAINT course_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS  course_on_menu (
    course_id INTEGER NOT NULL,
    menu_id INTEGER NOT NULL,
    corporation_id TEXT NOT NULL,

    CONSTRAINT course_on_menu_pkey PRIMARY KEY (course_id,menu_id)
);

-- CreateTable
CREATE TABLE IF NOT EXISTS  desk (
    id SERIAL NOT NULL,
    table_name TEXT NOT NULL,
    is_occupied BOOLEAN NOT NULL,
    corporation_id TEXT NOT NULL,

    CONSTRAINT desk_pkey PRIMARY KEY (id)
);

-- CreateTable
CREATE TABLE IF NOT EXISTS  menu (
    id SERIAL NOT NULL,
    menu_type TEXT NOT NULL,
    price INTEGER NOT NULL DEFAULT 0,
    available BOOLEAN NOT NULL DEFAULT true,
    has_image BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3),
    corporation_id TEXT NOT NULL,
    menu_name_en TEXT,
    menu_name_th TEXT,
    priority INTEGER,

    CONSTRAINT menu_pkey PRIMARY KEY (id)
);

-- CreateTable
CREATE TABLE IF NOT EXISTS  app_order (
    id SERIAL NOT NULL,
    order_amount INTEGER NOT NULL,
    total_price INTEGER NOT NULL,
    process_type TEXT NOT NULL,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3),
    channel_id TEXT NOT NULL,
    menu_id INTEGER NOT NULL,

    CONSTRAINT order_pkey PRIMARY KEY (id)
);

-- CreateIndex
CREATE UNIQUE INDEX user_email_key ON app_user(email);

-- CreateIndex
CREATE UNIQUE INDEX channel_id_key ON channel(id);


-- AddForeignKey
ALTER TABLE app_user ADD CONSTRAINT user_corporation_id_fkey FOREIGN KEY (corporation_id) REFERENCES corporation(id) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE channel ADD CONSTRAINT channel_course_id_fkey FOREIGN KEY (course_id) REFERENCES course(id) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE channel ADD CONSTRAINT channel_table_id_fkey FOREIGN KEY (table_id) REFERENCES desk(id) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE course ADD CONSTRAINT course_corporation_id_fkey FOREIGN KEY (corporation_id) REFERENCES Corporation(id) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE course_on_menu ADD CONSTRAINT courseOnMenu_corporation_id_fkey FOREIGN KEY (corporation_id) REFERENCES corporation(id) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE course_on_menu ADD CONSTRAINT courseOnMenu_course_id_fkey FOREIGN KEY (course_id) REFERENCES course(id) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE course_on_menu ADD CONSTRAINT courseOnMenu_menu_id_fkey FOREIGN KEY (menu_id) REFERENCES menu(id) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE desk ADD CONSTRAINT desk_corporation_id_fkey FOREIGN KEY (corporation_id) REFERENCES corporation(id) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE menu ADD CONSTRAINT menu_corporation_id_fkey FOREIGN KEY (corporation_id) REFERENCES corporation(id) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE app_order ADD CONSTRAINT order_channel_id_fkey FOREIGN KEY (channel_id) REFERENCES channel(id) ON DELETE CASCADE ON UPDATE CASCADE;

