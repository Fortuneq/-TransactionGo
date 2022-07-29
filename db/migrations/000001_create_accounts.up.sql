create table "accounts"(
	id serial primary key not null,
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
insert into "accounts"(id,user_name,balance) values
(1,'Vladislav',1550.22),
(2,'Alina',20231.1232),
(3,'Vladimir',68149.94),
(4,'Alexey',87517.25),
(5,'Maria',41314.59),
(6,'Daria',3066.74),
(7,'Natalya',97231.05),
(8,'Kirill',51270.212),
(9,'Ohmat',2019.96 ),
(10,'Ozamat',32285.96),
(11,'Roman',19756.01),
(12,'Maxim', 72323.2),
(13,'Anatoly', 123323.21);

alter table "transfer" add foreign key("from_user_id") references "accounts"("id");
alter table "transfer" add foreign key("to_user_id") references "accounts"("id");