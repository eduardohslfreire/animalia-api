create table if not exists tab_citizen (
	id serial not null,
	name varchar(50) not null,
    species varchar(50) not null,
    description varchar(100) not null,
    weight numeric not null,
    height numeric not null,
    photo_url varchar(100) null,
	has_pet_human boolean NULL default false,
	constraint tab_citizen_pkey primary key (id)
);

create table if not exists tab_role (
    id bigint not null,
    name varchar(50) not null, 
    single boolean NULL default false,
    constraint tab_role_pkey primary key (id)
);

create table if not exists tab_citizen_role (
    citizen_id bigint constraint fk_citizen references tab_citizen(id),
    role_id bigint constraint fk_role references tab_role(id)
);