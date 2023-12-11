```sql
汽车表
create table cars
(
    id             bigint
        primary key,
    make           varchar(20)              not null,
    model          varchar(20)              not null,
    price_per_hour real                     not null,
    price_per_day  real                     not null,
    capacity       integer                  not null,
    description    varchar(255)             not null,
    create_at      timestamp with time zone not null,
    update_at      timestamp with time zone not null,
    img            text                     not null
);

alter table cars
    owner to postgres;

枚举类型
create type enum_customers_stateissue as enum ('AL', 'AK', 'AZ', 'AR', 'CA', 'CO', 'CT', 'DE', 'FL', 'GA', 'HI', 'ID', 'IL', 'IN', 'IA', 'KS', 'KY', 'LA', 'ME', 'MD', 'MA', 'MI', 'MN', 'MS', 'MO', 'MT', 'NE', 'NV', 'NH', 'NJ', 'NM', 'NY', 'NC', 'ND', 'OH', 'OK', 'OR', 'PA', 'RI', 'SC', 'SD', 'TN', 'TX', 'UT', 'VT', 'VA', 'WA', 'WV', 'WI', 'WY');

alter type enum_customers_stateissue owner to postgres;

顾客表
create table customers
(
    uid         varchar(30)               not null
        primary key,
    firstname   varchar(100)              not null,
    lastname    varchar(100)              not null,
    address     varchar(100)              not null,
    email       varchar(30)               not null,
    credit_card varchar(20)               not null,
    is_student  boolean                   not null,
    telephone   varchar(30)               not null,
    phone       varchar(30)               not null,
    licence     integer                   not null,
    tickets     integer                   not null,
    state_issue enum_customers_stateissue not null,
    expiration  timestamp with time zone,
    create_at   timestamp with time zone  not null,
    update_at   timestamp with time zone  not null
);

alter table customers
    owner to postgres;

存取车地点表
create table locations
(
    id             bigint
        primary key,
    street_address varchar(100),
    telephone      varchar(30),
    create_at      timestamp with time zone not null,
    update_at      timestamp with time zone not null
);


alter table locations
    owner to postgres;

记录表
create table book_records
(
    reservation_num varchar(20)              not null
        primary key,
    price_per_hour  real                     not null,
    price_per_day   real                     not null,
    reserved_date   timestamp with time zone not null,
    pick_up_time    timestamp with time zone,
    drop_of_time    timestamp with time zone,
    create_at       timestamp with time zone not null,
    update_at       timestamp with time zone not null,
    car_id          bigint
        constraint fk_book_records_car
            references cars,
    location_id     bigint
        constraint fk_book_records_location
            references locations,
    customer_id     varchar(30)              not null
        constraint fk_book_records_customer
            references customers
);

alter table book_records
    owner to postgres;


```