create table g_id
(
	biz_tag varchar(128) not null,
	max_id bigint null,
	step int null,
	remark varchar(200) null,
	create_time bigint null,
	update_time bigint null,
	constraint g_id_pk
		primary key (biz_tag)
);

