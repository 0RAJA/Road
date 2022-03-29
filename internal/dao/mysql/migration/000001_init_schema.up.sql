create table if not exists manager
(
    username   varchar(50) not null
        primary key,
    password   text        not null,
    avatar_url text        not null,
    constraint manager_username_uindex
        unique (username)
);

create table if not exists post
(
    id          bigint                               not null
        primary key,
    cover       mediumtext                           not null,
    title       mediumtext                           not null,
    abstract    mediumtext                           not null,
    content     mediumtext                           not null,
    public      tinyint(1) default 1                 not null,
    deleted     tinyint(1) default 0                 not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    modify_time timestamp  default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint post_id_uindex
        unique (id)
)
    charset = utf8;

create table if not exists comment
(
    id            bigint                              not null
        primary key,
    post_id       bigint                              not null,
    username      varchar(10)                         not null,
    content       mediumtext                          not null,
    to_comment_id bigint    default 0                 not null,
    create_time   timestamp default CURRENT_TIMESTAMP not null,
    modify_time   timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint comment_id_uindex
        unique (id),
    constraint comment_post_id_fk
        foreign key (post_id) references road.post (id)
            on update cascade on delete cascade
)
    charset = utf8;

create table if not exists post_num
(
    post_id     bigint           not null
        primary key,
    star_num    bigint default 0 not null,
    visited_num bigint default 0 not null,
    constraint post_num_post_id_uindex
        unique (post_id),
    constraint post_num_post_id_fk
        foreign key (post_id) references road.post (id)
            on update cascade on delete cascade
);

create table if not exists tags
(
    id          bigint                              not null
        primary key,
    tag_name    varchar(10)                         not null,
    create_time timestamp default CURRENT_TIMESTAMP not null,
    constraint tags_id_uindex
        unique (id),
    constraint tags_tag_name_uindex
        unique (tag_name)
)
    charset = utf8;

create table if not exists post_tag
(
    id      bigint auto_increment
        primary key,
    post_id bigint not null,
    tag_id  bigint not null,
    constraint post_tag_id_uindex
        unique (id),
    constraint post_tag_post_id_tag_id_uindex
        unique (post_id, tag_id),
    constraint post_tag_post_id_fk
        foreign key (post_id) references road.post (id)
            on update cascade on delete cascade,
    constraint post_tag_tags_id_fk
        foreign key (tag_id) references road.tags (id)
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

create table if not exists user
(
    username       varchar(50)                         not null
        primary key,
    avatar_url     text                                not null,
    depository_url text                                not null,
    address        varchar(20)                         not null,
    create_time    timestamp default CURRENT_TIMESTAMP not null,
    modify_time    timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint user_username_uindex
        unique (username)
);

create table if not exists user_star
(
    id       int auto_increment
        primary key,
    username varchar(50) not null,
    post_id  bigint      not null,
    constraint user_star_id_uindex
        unique (id),
    constraint user_star_post_id_fk
        foreign key (post_id) references road.post (id)
            on update cascade on delete cascade,
    constraint user_star_user_username_fk
        foreign key (username) references road.user (username)
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

# 插入默认管理员--用户名:0RAJA 密码:12345
insert manager (username, password, avatar_url)
values ('0RAJA', '$2a$10$UhLdG/wYglcjhZ3LahHwjekbLqHJixsGF2SOi28IrpubvpkoPz53W', 'https://avatars.githubusercontent.com/u/76676061?v=4')
