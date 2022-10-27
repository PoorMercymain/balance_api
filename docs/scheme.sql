create table if not exists "user"(
    id serial constraint user_pk PRIMARY KEY,
    username text not null,
    balance integer not null
);

create table if not exists service(
    id  serial constraint service_pk PRIMARY KEY,
    service_name text not null,
    price integer not null
);

create table if not exists "order"(
    id  serial constraint order_pk PRIMARY KEY,
    user_id  serial constraint order_fk_user_id REFERENCES "user"
);

create table if not exists reserve(
    id  serial constraint reserve_pk PRIMARY KEY,
    user_id integer constraint reserve_fk_user_id REFERENCES "user",
    order_id integer constraint reserve_fk_order_id REFERENCES "order",
    money integer not null
);

create table if not exists order_service(
    order_id integer constraint order_service_fk_order_id REFERENCES "order",
    service_id integer constraint order_service_fk_service_id REFERENCES service
);

create table if not exists report(
    user_id integer constraint report_pk REFERENCES "user",
    service_id integer constraint report_fk REFERENCES service,
    revenue integer not null
);
