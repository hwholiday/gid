create table study.segments
(
    biz_tag     varchar(128) not null,
    max_id      bigint       null,
    step        int          null,
    remark      varchar(200) null,
    create_time bigint       null,
    update_time bigint       null,
    constraint segments_pk
        primary key (biz_tag)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

INSERT INTO study.segments(`biz_tag`, `max_id`, `step`, `remark`, `create_time`, `update_time`)
VALUES ('test', 0, 100000, 'test', 1591706686, 1591706686);