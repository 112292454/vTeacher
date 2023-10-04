-- auto-generated definition
create table user
(
    user_id                 int auto_increment comment '用户id'
        primary key,
    username                varchar(20)              not null comment '用户名',
    mail                    varchar(320)             not null comment '邮箱',
    avatar                  varchar(50)              null comment '头像token',
    password                varchar(255)             not null comment '密码',
    create_ip               varchar(80)              null comment '注册IP',
    last_login_ip           varchar(80)              null comment '最后登陆IP',
    bio                     varchar(160)             null comment '个人简介',
    blog                    varchar(255)             null comment '个人主页',
    follower_count          int unsigned default '0' not null comment '关注我的人数',
    followee_count          int unsigned default '0' not null comment '我关注的人数',
    following_article_count int unsigned default '0' not null comment '我关注的文章数',
    following_video_count   int unsigned default '0' not null comment '我关注的视频数',
    article_count           int unsigned default '0' not null comment '我发表的文章数量',
    video_count             int unsigned default '0' not null comment '我发表的视频数量',
    notification_unread     int unsigned default '0' not null comment '未读通知数',
    inbox_unread            int unsigned default '0' not null comment '未读私信数',
    create_time             int unsigned default '0' not null comment '注册时间',
    last_login_time         int unsigned default '0' not null comment '最后登录时间',
    disable_time            int unsigned default '0' not null comment '禁用时间'
)
    comment '用户表';

create index create_time
    on user (create_time);

create index email
    on user (mail);

create index follower_count
    on user (follower_count);

create index user_name
    on user (username);