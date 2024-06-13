create table passport(
    passport_id int generated always as identity primary key,
    type        varchar(32) not null,
    number      varchar(32) not null unique
);

create table company(
    company_id int generated always as identity primary key,
    name       varchar(100) not null unique
);

create table department(
    department_id int generated always as identity primary key,
    company_id    int references company (company_id),
    name          text unique not null,
    phone         text unique not null
);

create table employee(
    employee_id   bigint primary key,
    first_name    varchar(32) not null,
    second_name   varchar(32) not null,
    phone         varchar(20) not null unique,
    company_id    int references company (company_id),
    passport_id   int references passport (passport_id),
    department_id int references department (department_id)
);