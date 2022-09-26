CREATE TABLE IF NOT EXISTS public.phones
(
    id character varying(36) COLLATE pg_catalog."default" NOT NULL,
    "number" character varying(15) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT phones_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;