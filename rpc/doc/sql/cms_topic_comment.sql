create table cms_topic_comment
(
    id               bigint auto_increment
        primary key,
    member_nick_name varchar(255)  not null,
    topic_id         bigint        not null,
    member_icon      varchar(255)  not null,
    content          varchar(1000) not null,
    create_time      datetime      not null,
    show_status      int(1)        not null
)
    comment '专题评论表' charset = utf8;

