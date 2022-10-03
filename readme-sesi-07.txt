Belajar:
- GO API
   + Make web service CRU\D

create table employees before you run project in sesi-07, because we need that table.
follow this SQL:
   create table employees(
	id SERIAL primary key,
	full_name varchar(50) not null,
	email text unique not null,
	age int not null,
	division varchar(20) not null
);

- Gorm Framework
   + Make web service CRUD with ORM
   + Eager Load
   + Hook

