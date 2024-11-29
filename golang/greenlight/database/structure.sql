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
-- Name: events; Type: TABLE; Schema: public; Owner: greenlight
--

CREATE TABLE public.events (
    aggregate_id uuid NOT NULL,
    version integer NOT NULL,
    parent_id uuid NOT NULL,
    event_type character varying(50),
    payload jsonb NOT NULL,
    created_at timestamp without time zone NOT NULL
);


ALTER TABLE public.events OWNER TO greenlight;

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
-- Name: permissions; Type: TABLE; Schema: public; Owner: greenlight
--

CREATE TABLE public.permissions (
    id bigint NOT NULL,
    code text NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.permissions OWNER TO greenlight;

--
-- Name: permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: greenlight
--

CREATE SEQUENCE public.permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.permissions_id_seq OWNER TO greenlight;

--
-- Name: permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: greenlight
--

ALTER SEQUENCE public.permissions_id_seq OWNED BY public.permissions.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: greenlight
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO greenlight;

--
-- Name: tokens; Type: TABLE; Schema: public; Owner: greenlight
--

CREATE TABLE public.tokens (
    hash bytea NOT NULL,
    user_id bigint NOT NULL,
    scope text NOT NULL,
    expire_at timestamp(0) with time zone NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.tokens OWNER TO greenlight;

--
-- Name: user_permissions; Type: TABLE; Schema: public; Owner: greenlight
--

CREATE TABLE public.user_permissions (
    user_id bigint NOT NULL,
    permission_id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.user_permissions OWNER TO greenlight;

--
-- Name: users; Type: TABLE; Schema: public; Owner: greenlight
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    name text NOT NULL,
    email public.citext NOT NULL,
    password_hash bytea NOT NULL,
    status character varying(50) NOT NULL,
    version integer DEFAULT 1 NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    updated_at timestamp(0) with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.users OWNER TO greenlight;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: greenlight
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO greenlight;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: greenlight
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- Name: permissions id; Type: DEFAULT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.permissions ALTER COLUMN id SET DEFAULT nextval('public.permissions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (aggregate_id, version);


--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: permissions permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: tokens tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_pkey PRIMARY KEY (hash);


--
-- Name: user_permissions user_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.user_permissions
    ADD CONSTRAINT user_permissions_pkey PRIMARY KEY (user_id, permission_id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: movies_genres_idx; Type: INDEX; Schema: public; Owner: greenlight
--

CREATE INDEX movies_genres_idx ON public.movies USING gin (genres);


--
-- Name: movies_title_idx; Type: INDEX; Schema: public; Owner: greenlight
--

CREATE INDEX movies_title_idx ON public.movies USING gin (to_tsvector('english'::regconfig, title));


--
-- Name: tokens tokens_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.tokens
    ADD CONSTRAINT tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: user_permissions user_permissions_permission_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.user_permissions
    ADD CONSTRAINT user_permissions_permission_id_fkey FOREIGN KEY (permission_id) REFERENCES public.permissions(id) ON DELETE CASCADE;


--
-- Name: user_permissions user_permissions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: greenlight
--

ALTER TABLE ONLY public.user_permissions
    ADD CONSTRAINT user_permissions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: pg_database_owner
--

GRANT ALL ON SCHEMA public TO greenlight;


--
-- PostgreSQL database dump complete
--

