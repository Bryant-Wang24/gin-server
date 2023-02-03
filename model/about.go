package model

import "time"

type Image struct {
	ID     int    `json:"id,omitempty" gorm:"primary_key"`
	ImgUrl string `json:"imgUrl"`
	Link   string `json:"link"`
}

type About struct {
	ID         int       `json:"id" gorm:"primary_key"`
	Tags       string    `json:"tags"`
	Desc       string    `json:"desc"`
	ShowResume int       `json:"showResume"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Imgs       string    `json:"imgs"`
}

type AboutInfo struct {
	ID         int       `json:"id,omitempty"`
	Tags       []string  `json:"tags"`
	Desc       string    `json:"desc"`
	ShowResume int       `json:"showResume"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Imgs       []Image   `json:"imgs"`
}

// 创建关于表
// create table abouts(
// 	`id` int(11) not null auto_increment comment 'id',
// 	`tags` varchar(255) not null comment '标签',
// 	`desc` varchar(255) not null comment '描述',
// 	`show_Resume` tinyint(1) not null default 0 comment '是否显示简历',
// 	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
// 	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
// 	`imgs` varchar(255) not null comment '图片',
// 	primary key(`id`)
// )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='文章表';

// 创建图片表
// create table images(
// 	`id` int(11) not null auto_increment comment 'id',
// 	`img_url` varchar(255) not null comment '图片地址',
//  `link` varchar(255) not null comment '图片链接',
// 	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
// 	`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
// 	primary key(`id`)
// )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='图片表';
