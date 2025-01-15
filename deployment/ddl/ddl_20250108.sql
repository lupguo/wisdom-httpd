create table if not exists wisdom_db.t_wisdoms
(
    id         int auto_increment comment '主键ID'
        primary key,
    wisdom_no  varchar(255)                       not null comment '名言16进制编码（对外展示）',
    sentence   text                               null comment '名言句子',
    speaker    varchar(255)                       null comment '名人',
    refer_url  varchar(255)                       null comment '出处URL, URL Reference',
    image      varchar(255)                       null comment '名言关联图片（由AI生成）',
    created_at datetime default CURRENT_TIMESTAMP null comment '创建时间',
    updated_at datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    deleted_at datetime                           null comment '删除时间',
    constraint wisdom_no
        unique (wisdom_no)
)
    comment '名言金句存储表';

create index idx_created_at
    on wisdom_db.t_wisdoms (created_at);

create index idx_deleted_at
    on wisdom_db.t_wisdoms (deleted_at);

create index idx_speaker
    on wisdom_db.t_wisdoms (speaker);

create index idx_updated_at
    on wisdom_db.t_wisdoms (updated_at);

create index idx_wisdom_no
    on wisdom_db.t_wisdoms (wisdom_no);

