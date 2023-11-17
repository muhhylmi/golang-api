ALTER TABLE IF EXISTS books
    ADD COLUMN price double precision NOT NULL;

create table carts (
    id varchar(40) not null,
    user_id varchar(40) not null,
    price double precision,

    created_at bigint,
    created_by varchar(40),
    updated_at bigint,
    updated_by varchar(40),
    deleted_at bigint,
    deleted_by varchar(40),
    primary key (id)
);

create table cart_details (
    id varchar(40) not null,
    cart_id varchar(40) not null,
    book_id varchar(40) not null,
    status varchar(10) not null,

    created_at bigint,
    created_by varchar(40),
    updated_at bigint,
    updated_by varchar(40),
    deleted_at bigint,
    deleted_by varchar(40),
    primary key (id)
);
