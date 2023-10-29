create table tasks (
	id serial CONSTRAINT todo_pk PRIMARY KEY,
	todo_id int8,
	title VARCHAR(255) not null,
	description text not null,
	file text null,
	deleted_at timestamp null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at timestamp default CURRENT_TIMESTAMP null
);