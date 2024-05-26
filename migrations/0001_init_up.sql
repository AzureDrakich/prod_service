BEGIN;


SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = public;
SET default_tablespace = '';
SET default_with_oids = FALSE;


-- EXTENSTIONS --
CREATE EXTENSION IF NOT EXISTS pgcrypto;


-- TABLES --

CREATE TABLE IF NOT EXISTS public.currency
(
    id SERIAL PRIMARY KEY,
    name TEXT,
    symbol TEXT
);

CREATE TABLE public.category
(
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE IF NOT EXISTS public.product
(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    image_id UUID,
    price INT,
    currency_id INT,
    rating INT,
    category_id INT NOT NULL,
    specification JSONB,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

-- DATA --

INSERT INTO public.currency (name, symbol) VALUES ('рубль','Р');
INSERT INTO public.currency (name, symbol) VALUES ('доллар','$');
COMMIT;