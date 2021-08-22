create table if not exists item
(
    id        int auto_increment
        primary key,
    uuid      text null,
    name      text null,
    createdBy text null
);

