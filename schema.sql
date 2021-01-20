create table users
(
    id         varchar(50) not null
        constraint users_pkey primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username   varchar(100)
        constraint users_username_key
            unique,
    password   text        not null,
    email      varchar(50) not null
        constraint users_email_key
            unique,
    profile    varchar(255)
);

create index idx_users_deleted_at
    on users (deleted_at);

create table departments
(
    id          varchar(50)  not null
        constraint departments_pkey
            primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    name        varchar(255) not null,
    description text
);

create index idx_departments_name
    on departments (name);

create index idx_departments_deleted_at
    on departments (deleted_at);

create table employees
(
    id         varchar(50)  not null
        constraint employees_pkey
            primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    full_name  varchar(255) not null,
    phone      varchar(255) not null,
    age        smallint,
    status     varchar(20)  not null,
    user_id    varchar(50)
        constraint fk_employees_user
            references users
            on update cascade on delete set null
);

create index idx_employees_full_name
    on employees (full_name);

create index idx_employees_deleted_at
    on employees (deleted_at);

create table employee_departments
(
    employee_id   varchar(50) not null
        constraint fk_employee_departments_employee
            references employees,
    department_id varchar(50) not null
        constraint fk_employee_departments_department
            references departments,
    constraint employee_departments_pkey
        primary key (employee_id, department_id)
);

create table projects
(
    id            varchar(50)  not null
        constraint projects_pkey
            primary key,
    created_at    timestamp with time zone,
    updated_at    timestamp with time zone,
    deleted_at    timestamp with time zone,
    name          varchar(255) not null,
    description   text,
    department_id varchar(50)
        constraint fk_projects_department
            references departments
            on update cascade on delete set null
);

create index idx_projects_deleted_at
    on projects (deleted_at);

create table teams
(
    id          varchar(50)  not null
        constraint teams_pkey
            primary key,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    deleted_at  timestamp with time zone,
    name        varchar(255) not null,
    employee_id varchar(50)
        constraint fk_teams_employee
            references employees
            on update cascade on delete set null,
    project_id  varchar(50)
        constraint fk_teams_project
            references projects
            on update cascade on delete set null
);

create index idx_teams_deleted_at
    on teams (deleted_at);

create table employee_works
(
    team_id     varchar(50) not null
        constraint fk_employee_works_team
            references teams,
    employee_id varchar(50) not null
        constraint fk_employee_works_employee
            references employees,
    constraint employee_works_pkey
        primary key (team_id, employee_id)
);


