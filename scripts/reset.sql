drop table if exists affinities;
drop table if exists flavors;

create table flavors (
  id serial primary key,
  name varchar(100) not null
);
create table affinities (
  flavorOne int not null,
  flavorTwo int not null
);

insert into flavors (name) values ('garlic');
insert into flavors (name) values ('tomato');
insert into flavors (name) values ('thyme');
insert into flavors (name) values ('vanilla');
insert into flavors (name) values ('lemon');
insert into flavors (name) values ('rosemary');
