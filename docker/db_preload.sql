-- public.jobs definition

-- Drop table

-- DROP TABLE public.jobs;

CREATE TABLE public.jobs (
                             id bigserial NOT NULL,
                             name text NOT NULL,
                             result text,
                             status varchar(10) NOT NULL,
                             data text NOT NULL,
                             CONSTRAINT jobs_pkey PRIMARY KEY (id)
);
CREATE INDEX jobs_queue_index ON public.jobs USING btree (name);

-- Permissions

ALTER TABLE public.jobs OWNER TO postgres;
GRANT ALL ON TABLE public.jobs TO postgres;

INSERT INTO public.jobs (name, "result", status, "data") VALUES('My Job 1', 'something', 'pending', '[1, 2, 3]');
INSERT INTO public.jobs (name, "result", status, "data") VALUES('My Job 2', null, 'done', '[1, 2, 3]');
INSERT INTO public.jobs (name, "result", status, "data") VALUES('My Job 3', null, 'pending', '[1, 2, 3]');
INSERT INTO public.jobs (name, "result", status, "data") VALUES('My Job 4', 'something something', 'error', '[1, 2, 3]');
