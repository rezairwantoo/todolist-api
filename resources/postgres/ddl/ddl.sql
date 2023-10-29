-- public.tasks definition

-- Drop table

-- DROP TABLE public.tasks;

CREATE TABLE public.tasks (
	id serial4 NOT NULL,
	todo_id int8 NULL,
	title varchar(255) NOT NULL,
	description text NOT NULL,
	file text NULL,
	deleted_at timestamp NULL,
	created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT todo_pk PRIMARY KEY (id)
);