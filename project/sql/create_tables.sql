-- Database export via SQLPro (https://www.sqlprostudio.com/allapps.html)
-- Exported by jhonisoares at 02-11-2023 09:04.
-- WARNING: This file may contain descructive statements such as DROPs.
-- Please ensure that you are running the script at the proper location.


-- BEGIN TABLE public.products
BEGIN;

CREATE TABLE IF NOT EXISTS public.products (
	id bigint DEFAULT nextval('products_sampleid_seq'::regclass) NOT NULL,
	name character(1) NOT NULL,
	description character(1) NOT NULL,
	photos text NOT NULL,
	created_at timestamp without time zone NOT NULL,
	updated_at timestamp without time zone NOT NULL,
	PRIMARY KEY(id)
);

COMMIT;

-- Table public.products contains no data. No inserts have been genrated.
-- Inserting 0 rows into public.products


-- END TABLE public.products

-- BEGIN TABLE public.users
BEGIN;

CREATE TABLE IF NOT EXISTS public.users (
	id integer DEFAULT nextval('user_id_seq'::regclass) NOT NULL,
	email character varying(255),
	first_name character varying(255),
	last_name character varying(255),
	"password" character varying(60),
	user_active integer DEFAULT 0,
	created_at timestamp without time zone,
	updated_at timestamp without time zone,
	PRIMARY KEY(id)
);

COMMIT;

-- Inserting 1 row into public.users
-- Insert batch #1
INSERT INTO public.users (id, email, first_name, last_name, "password", user_active, created_at, updated_at) VALUES
(2, 'admin@example.com', 'Admin', 'User', '$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe', 1, '2022-03-14 00:00:00', '2022-03-14 00:00:00');

-- END TABLE public.users

