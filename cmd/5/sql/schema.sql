create table if not exists users (
  id varchar(50) primary key,
  first_name varchar(20) not null,
  last_name varchar(20) not null
);