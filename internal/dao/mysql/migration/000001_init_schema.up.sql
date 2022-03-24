create table manager
(
    username   varchar(10)                    not null
        primary key,
    password   varchar(32)                    not null,
    avatar_url varchar(20) default 'avl_test' not null,
    constraint manager_username_uindex
        unique (username)
);


create table post
(
    id          bigint                                not null
        primary key,
    cover       varchar(20) default '1111'            not null,
    title       text                                  not null,
    abstract    text                                  not null,
    content     text                                  not null,
    public      tinyint(1)  default 1                 not null,
    deleted     tinyint(1)  default 0                 not null,
    create_time timestamp   default CURRENT_TIMESTAMP not null,
    modify_time timestamp   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint post_id_uindex
        unique (id)
);

create table if not exists tags
(
    id          bigint                              not null
        primary key,
    tag_name    varchar(10)                         not null,
    create_time timestamp default CURRENT_TIMESTAMP not null,
    constraint tags_id_uindex
        unique (id)
);

create table post_tag
(
    id      bigint auto_increment
        primary key,
    post_id bigint not null,
    tag_id  bigint not null,
    constraint post_tag_id_uindex
        unique (id),
    constraint post_tag_post_id_fk
        foreign key (post_id) references post (id)
            on update cascade on delete cascade,
    constraint post_tag_tags_id_fk
        foreign key (tag_id) references tags (id)
            on update cascade on delete cascade
);

create table if not exists tops
(
    id      bigint not null
        primary key,
    post_id bigint not null,
    constraint tops_id_uindex
        unique (id),
    constraint tops_post_id_uindex
        unique (post_id),
    constraint tops_post_id_fk
        foreign key (post_id) references road.post (id)
            on update cascade on delete cascade
);

create table user
(
    username       varchar(10)                         not null
        primary key,
    avatar_url     varchar(20)                         not null,
    depository_url varchar(20)                         not null,
    address        varchar(20)                         not null,
    create_time    timestamp default CURRENT_TIMESTAMP not null,
    modify_time    timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint user_username_uindex
        unique (username)
);


create table if not exists comment
(
    id            bigint                              not null
        primary key,
    post_id       bigint                              not null,
    username      varchar(10)                         not null,
    content       text                                not null,
    to_comment_id bigint    default 0                 not null,
    create_time   timestamp default CURRENT_TIMESTAMP not null,
    modify_time   timestamp default CURRENT_TIMESTAMP not null,
    constraint comment_id_uindex
        unique (id),
    constraint comment_user_username_fk
        foreign key (username) references road.user (username)
            on update cascade
);

create index comment_post_id_fk
    on road.comment (post_id);

create table user_star
(
    id       int auto_increment
        primary key,
    username varchar(10) not null,
    post_id  bigint      not null,
    constraint user_star_id_uindex
        unique (id),
    constraint user_star_post_id_fk
        foreign key (post_id) references post (id)
            on update cascade on delete cascade,
    constraint user_star_user_username_fk
        foreign key (username) references user (username)
            on update cascade on delete cascade
);

create table if not exists views
(
    id          bigint auto_increment
        primary key,
    views_num   bigint    default 0                 not null,
    create_time timestamp default CURRENT_TIMESTAMP not null,
    constraint views_id_uindex
        unique (id)
);

create table post_num
(
    post_id     bigint           not null
        primary key,
    star_num    bigint default 0 not null,
    visited_num bigint default 0 not null,
    constraint post_num_post_id_uindex
        unique (post_id),
    constraint post_num_post_id_fk
        foreign key (post_id) references post (id)
            on update cascade on delete cascade
);

