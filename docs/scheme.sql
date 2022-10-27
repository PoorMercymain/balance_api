create table if not exists task(
    id serial constraint task_pk PRIMARY KEY,
    order_name text not null,
    start_date timestamptz
);

create table if not exists work(
    id  serial constraint work_pk PRIMARY KEY,
    task_id integer constraint work_fk REFERENCES task ON DELETE CASCADE,
    duration integer not null,
    resource integer not null,
    parent_id integer
);

create table if not exists reserve(
    reserve_id  serial constraint reserve_pk PRIMARY KEY,
    user_id integer constraint reserve_fk_user_id REFERENCES user,
    order_id integer constraint reserve_fk_order_id REFERENCES order,
    service_id integer constraint reserve_fk_service_id REFERENCES service,
    money integer not null
);

create table if not exists service(
    service_id  serial constraint service_pk PRIMARY KEY,
    service_name text not null,
    price integer not null
);

create table if not exists user(
    user_id  serial constraint user_pk PRIMARY KEY,
    username text not null,
    balance integer not null
);

create table if not exists order(
    order_id  serial constraint order_pk PRIMARY KEY,
    user_id  serial constraint order_fk_user_id REFERENCES user,
    service_ids integer array,
    total integer not null
);
