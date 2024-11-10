--
-- PostgreSQL database dump
--

-- Dumped from database version 17.0
-- Dumped by pg_dump version 17.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: movies; Type: TABLE; Schema: public; Owner: greenlight
--

CREATE TABLE public.movies (
    id bigint NOT NULL,
    title text NOT NULL,
    year integer NOT NULL,
    runtime integer NOT NULL,
    genres text[] NOT NULL,
    version integer DEFAULT 1 NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    CONSTRAINT genres_length_check CHECK (((array_length(genres, 1) >= 1) AND (array_length(genres, 1) <= 5))),
    CONSTRAINT movies_runtime_check CHECK ((runtime >= 0)),
    CONSTRAINT movies_year_check CHECK (((year >= 1888) AND ((year)::double precision <= date_part('year'::text, now()))))
);


ALTER TABLE public.movies OWNER TO greenlight;

--
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: greenlight
--

CREATE SEQUENCE public.movies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.movies_id_seq OWNER TO greenlight;

--
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: greenlight
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: greenlight
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO greenlight;

--
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: pg_database_owner
--

GRANT ALL ON SCHEMA public TO greenlight;


--
-- PostgreSQL database dump complete
--

