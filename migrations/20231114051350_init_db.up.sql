create table books (
    id varchar(40) not null,
    title varchar(150) not null,
    year varchar(5),
    author varchar(100),
    created_at bigint,
    created_by varchar(40),
    updated_at bigint,
    updated_by varchar(40),
    deleted_at bigint,
    deleted_by varchar(40),
    primary key (id)
)