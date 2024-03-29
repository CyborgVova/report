--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1
-- Dumped by pg_dump version 15.0

-- Started on 2024-03-29 02:35:46

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'WIN1251';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 17834)
-- Name: reports; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.reports (
    report_id bigint NOT NULL,
    creation_time date,
    report_info character(1),
    model_id bigint
);


ALTER TABLE public.reports OWNER TO "user";

--
-- TOC entry 214 (class 1259 OID 17833)
-- Name: reports_report_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.reports_report_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.reports_report_id_seq OWNER TO "user";

--
-- TOC entry 3325 (class 0 OID 0)
-- Dependencies: 214
-- Name: reports_report_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.reports_report_id_seq OWNED BY public.reports.report_id;


--
-- TOC entry 3173 (class 2604 OID 17837)
-- Name: reports report_id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.reports ALTER COLUMN report_id SET DEFAULT nextval('public.reports_report_id_seq'::regclass);


--
-- TOC entry 3319 (class 0 OID 17834)
-- Dependencies: 215
-- Data for Name: reports; Type: TABLE DATA; Schema: public; Owner: user
--

INSERT INTO public.reports (report_id, creation_time, report_info, model_id) VALUES (25, '2023-03-02', 'A', 10);
INSERT INTO public.reports (report_id, creation_time, report_info, model_id) VALUES (26, '2023-03-04', 'B', 20);
INSERT INTO public.reports (report_id, creation_time, report_info, model_id) VALUES (27, '2023-03-05', 'C', 30);
INSERT INTO public.reports (report_id, creation_time, report_info, model_id) VALUES (28, '2023-03-09', 'D', 40);
INSERT INTO public.reports (report_id, creation_time, report_info, model_id) VALUES (29, '2023-03-13', 'E', 50);
INSERT INTO public.reports (report_id, creation_time, report_info, model_id) VALUES (30, '2023-03-07', 'Z', 30);
INSERT INTO public.reports (report_id, creation_time, report_info, model_id) VALUES (31, '2023-03-14', 'F', 60);
INSERT INTO public.reports (report_id, creation_time, report_info, model_id) VALUES (32, '2023-03-20', 'G', 70);
INSERT INTO public.reports (report_id, creation_time, report_info, model_id) VALUES (33, '2023-03-25', 'J', 80);
INSERT INTO public.reports (report_id, creation_time, report_info, model_id) VALUES (34, '2023-03-25', 'O', 30);


--
-- TOC entry 3326 (class 0 OID 0)
-- Dependencies: 214
-- Name: reports_report_id_seq; Type: SEQUENCE SET; Schema: public; Owner: user
--

SELECT pg_catalog.setval('public.reports_report_id_seq', 34, true);


--
-- TOC entry 3175 (class 2606 OID 17839)
-- Name: reports reports_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.reports
    ADD CONSTRAINT reports_pkey PRIMARY KEY (report_id);


-- Completed on 2024-03-29 02:35:46

--
-- PostgreSQL database dump complete
--

