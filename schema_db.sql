create table users
(
    id         int auto_increment
        primary key,
    email      varchar(50)  not null,
    username   varchar(50)  not null,
    password   varchar(255) not null,
    created_at timestamp    null,
    updated_at timestamp    null,
    constraint email
        unique (email)
);

create table matches
(
    id         int auto_increment
        primary key,
    user_id    int       not null,
    person_id  int       not null,
    status     tinyint   not null,
    created_at timestamp null,
    updated_at timestamp null,
    constraint matches_ibfk_1
        foreign key (user_id) references users (id),
    constraint matches_ibfk_2
        foreign key (person_id) references users (id)
);

create index person_id
    on matches (person_id);

create index user_id
    on matches (user_id);

create table profiles
(
    id         int auto_increment
        primary key,
    user_id    int          not null,
    photo      varchar(255) not null,
    birth_date date         not null,
    gender     tinyint      not null,
    bio        text         null,
    created_at timestamp    null,
    updated_at timestamp    null,
    constraint profiles_ibfk_1
        foreign key (user_id) references users (id)
);

create index user_id
    on profiles (user_id);

create table subscriptions
(
    id         int auto_increment
        primary key,
    user_id    int       not null,
    start_date timestamp null,
    end_date   timestamp null,
    status     tinyint   not null,
    created_at timestamp null,
    updated_at timestamp null,
    constraint subscriptions_ibfk_1
        foreign key (user_id) references users (id)
);

create table payments
(
    id              int auto_increment
        primary key,
    user_id         int                           not null,
    subscription_id int                           not null,
    amount          int                           not null,
    payment_method  varchar(50)                   not null,
    status          varchar(20) default 'PENDING' not null,
    payment_date    timestamp                     null,
    created_at      timestamp                     null,
    updated_at      timestamp                     null,
    constraint payments_ibfk_1
        foreign key (user_id) references users (id),
    constraint payments_ibfk_2
        foreign key (subscription_id) references subscriptions (id)
);

create index subscription_id
    on payments (subscription_id);

create index user_id
    on payments (user_id);

create index user_id
    on subscriptions (user_id);

