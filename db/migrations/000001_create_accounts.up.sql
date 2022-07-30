create table "accounts"(
	id bigserial primary key not null,
	user_name varchar(250) not null,
	balance real not null
);

create table "transfer"(
	id serial primary key not null,
	from_user_id bigint not null,
	to_user_id bigint not null,
	amount real not null,
	created_at timestamp not null default(now())
);

alter table "transfer" add foreign key("from_user_id") references "accounts"("id");
alter table "transfer" add foreign key("to_user_id") references "accounts"("id");