create table if not exists users (
    id varchar(40) not null,
    username varchar(150) not null,
    password varchar(200) not null,
    gender varchar(1) not null,
    created_at bigint,
    created_by varchar(40),
    updated_at bigint,
    updated_by varchar(40),
    deleted_at bigint,
    deleted_by varchar(40),
    primary key (id)
)