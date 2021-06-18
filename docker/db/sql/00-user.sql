create table practice.user
(
    id   bigint unsigned auto_increment,
    name varchar(60) charset utf8,
    constraint id
        unique (id)
);