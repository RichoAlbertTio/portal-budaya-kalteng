-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- USERS
CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	username VARCHAR(50) UNIQUE NOT NULL,
	password_hash VARCHAR(255) NOT NULL,
	full_name VARCHAR(100),
	role VARCHAR(20) NOT NULL DEFAULT 'admin', -- admin|superadmin
	bio TEXT,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- ABOUT (single page, bisa multi-entry untuk versi/riwayat)
CREATE TABLE IF NOT EXISTS about (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
title VARCHAR(150) NOT NULL,
description TEXT NOT NULL,
updated_by UUID REFERENCES users(id) ON DELETE SET NULL,
created_at TIMESTAMP NOT NULL DEFAULT NOW(),
updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- CATEGORIES
CREATE TABLE IF NOT EXISTS categories (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
name VARCHAR(80) NOT NULL,
slug VARCHAR(100) UNIQUE NOT NULL,
description TEXT,
created_at TIMESTAMP NOT NULL DEFAULT NOW(),
updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_categories_name ON categories(name);


-- TRIBES (Suku di Kalteng)
CREATE TABLE IF NOT EXISTS tribes (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
name VARCHAR(100) NOT NULL,
slug VARCHAR(120) UNIQUE NOT NULL,
description TEXT
);


-- REGIONS (Kab/Kota atau kawasan budaya)
CREATE TABLE IF NOT EXISTS regions (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
name VARCHAR(100) NOT NULL,
slug VARCHAR(120) UNIQUE NOT NULL,
description TEXT
);


-- CONTENTS (Artikel/Objek budaya)
CREATE TABLE IF NOT EXISTS contents (
id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
title VARCHAR(180) NOT NULL,
slug VARCHAR(200) UNIQUE NOT NULL,
image_url TEXT,
summary TEXT,
body TEXT NOT NULL,
status VARCHAR(20) NOT NULL DEFAULT 'draft', -- draft|published
published_at TIMESTAMP,
category_id UUID REFERENCES categories(id) ON DELETE SET NULL,
author_id UUID REFERENCES users(id) ON DELETE SET NULL,
created_at TIMESTAMP NOT NULL DEFAULT NOW(),
updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- PIVOTS
CREATE TABLE IF NOT EXISTS content_tribes (
content_id UUID REFERENCES contents(id) ON DELETE CASCADE,
tribe_id UUID REFERENCES tribes(id) ON DELETE CASCADE,
PRIMARY KEY(content_id, tribe_id)
);


CREATE TABLE IF NOT EXISTS content_regions (
content_id UUID REFERENCES contents(id) ON DELETE CASCADE,
region_id UUID REFERENCES regions(id) ON DELETE CASCADE,
PRIMARY KEY(content_id, region_id)
);


