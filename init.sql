create database oauth2;

use oauth2;

drop table if exists users;
create table users
(
    id          int unsigned primary key auto_increment,
    user_id     bigint unsigned unique                      not null comment '用户ID',
    phone       varchar(32) unique                          null comment '手机号码',
    country     smallint unsigned default 86                not null comment '国家号（中国86）',
    username    varchar(32)                                 not null comment '用户姓名',
    avatar      varchar(256)                                not null comment '用户头像',
    status      tinyint unsigned  default 0                 not null comment '状态',
    json_extent json                                        null comment 'json拓展',
    created_at  datetime          default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at  datetime          default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    deleted_at  datetime                                    null comment '删除时间',
    index username (username),
    index create_at (created_at)
) comment '用户表';


drop table if exists wechat_users;
create table wechat_users
(
    id          int unsigned primary key auto_increment,
    user_id     bigint unsigned unique                     not null comment '用户ID',
    open_id     varchar(32) unique                         null comment '微信 openid',
    union_id    varchar(32) unique                         null comment '微信 unionid',
    status      tinyint unsigned default 0                 not null comment '状态',
    json_extent json                                       null comment 'json拓展',
    created_at  datetime         default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at  datetime         default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    deleted_at  datetime                                   null comment '删除时间'
) comment '用户微信表';


drop table if exists configs;
create table configs
(
    id         int unsigned primary key auto_increment,
    config_key varchar(32)                                not null comment '配置名称',
    version    tinyint unsigned                           not null comment '版本',
    config     json                                       not null comment '配置',
    status     tinyint unsigned default 0                 not null comment '状态',
    created_at datetime         default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at datetime         default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    deleted_at datetime                                   null comment '删除时间',
    unique index key_version (config_key, version)
) comment '配置表';


INSERT INTO oauth2.configs (config_key, version, config)
VALUES ('oauth2_system', 1, '{
  "oauth2": "http://127.0.0.1:3000/oauth2"
}');