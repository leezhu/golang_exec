### 新建表格
```sql
book表
create table book (
	`id` int(11) auto_increment,
	`name` varchar(100) not null default '' comment 'name',
	`author` varchar(100) not null default '' comment 'name',
	`image` varchar(100) not null default '' comment 'name',
	`created_at` TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP comment '创建时间',
	`updated_at` TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP comment '更新时间',
	 primary key(`id`),
	unique key(`name`)
);

chapter表
create table chapter (
	`id` int(11) auto_increment,
	`book_id` int(11) not null default 0 comment '',
	`title` varchar(100) not null default '' comment 'title',
	`content` text comment 'content',
	`sort` int not null default 0 comment '',
	`pre` int not null default 0 comment '',
	`next` int not null default 0 comment '',
	`created_at` TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP comment '创建时间',
	`updated_at` TIMESTAMP not null DEFAULT CURRENT_TIMESTAMP comment '更新时间',
	 primary key(`id`)
);
url表
create table url (
	`id` int(11) auto_increment primary key,
	`url` varchar(100) not null default '',
	`status` int not null default -1 ,
	unique key(url)
);
```

