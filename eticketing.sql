--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: terminals; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.terminals (
    id bigint NOT NULL,
    name text,
    location text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.terminals OWNER TO postgres;

--
-- Name: terminals_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.terminals_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.terminals_id_seq OWNER TO postgres;

--
-- Name: terminals_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.terminals_id_seq OWNED BY public.terminals.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    username text,
    name text,
    email text,
    role text,
    password text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: terminals id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.terminals ALTER COLUMN id SET DEFAULT nextval('public.terminals_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: terminals; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.terminals (id, name, location, created_at, updated_at) FROM stdin;
1			2025-07-08 15:10:05.348049+07	2025-07-08 15:10:05.348049+07
2	Terminal Terboyo	Semarang	2025-07-08 15:11:18.932947+07	2025-07-08 15:11:18.932947+07
3	Terminal Terboyo	Semarang	2025-07-08 15:11:21.42182+07	2025-07-08 15:11:21.42182+07
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, username, name, email, role, password, created_at, updated_at) FROM stdin;
1	ryan123	ryan	admin@test.com	admin	admin123	\N	\N
\.


--
-- Name: terminals_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.terminals_id_seq', 3, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: terminals terminals_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.terminals
    ADD CONSTRAINT terminals_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

