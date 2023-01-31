package model

import "time"

type Article struct {
	ID            int       `json:"id" gorm:"primary_key"`
	Title         string    `json:"title"`         //文章标题
	Cover         string    `json:"cover"`         //文章封面
	Introduction  string    `json:"introduction"`  //文章简介
	Content       string    `json:"content"`       //文章内容
	Categories    string    `json:"categories"`    //文章分类
	Tags          string    `json:"tags"`          //文章标签
	Views         int       `json:"views"`         //文章浏览量
	Comment       int       `json:"comment"`       //文章评论数
	Like          int       `json:"like"`          //文章点赞数
	Collect       int       `json:"collect"`       //文章收藏数
	IsComment     int       `json:"isComment"`     //是否开启评论
	IsLike        int       `json:"isLike"`        //是否开启点赞
	IsCollect     int       `json:"isCollect"`     //是否开启收藏
	IsReward      int       `json:"isReward"`      //是否开启打赏
	Status        int       `json:"status"`        //文章状态
	PublishStatus int       `json:"publishStatus"` //文章发布状态
	Sort          int       `json:"sort"`          //文章排序
	CreateTime    time.Time `json:"createTime"`    //文章创建时间
	UpdateTime    time.Time `json:"updateTime"`    //文章更新时间
}

//创建文章表
//CREATE TABLE `articles` (
//	  `id` int(11) NOT NULL AUTO_INCREMENT,
//	  `title` varchar(255) NOT NULL comment '文章标题',
//	  `cover` varchar(255) NOT NULL comment '文章封面',
//	  `introduction` varchar(255) NOT NULL comment '文章简介',
//	  `content` text NOT NULL comment '文章内容',
//	  `categories` varchar(255) NOT NULL comment '文章分类',
//	  `tags` varchar(255) NOT NULL comment '文章标签',
//	  `views` int(11) NOT NULL DEFAULT '0' comment '文章浏览量',
//	  `comment` int(11) NOT NULL DEFAULT '0' comment '文章评论数',
//	  `like` int(11) NOT NULL DEFAULT '0' comment '文章点赞数',
//	  `collect` int(11) NOT NULL DEFAULT '0' comment '文章收藏数',
//	  `is_comment` tinyint(1) NOT NULL DEFAULT '1' comment '是否开启评论',
//	  `is_like` tinyint(1) NOT NULL DEFAULT '1' comment '是否开启点赞',
//	  `is_collect` tinyint(1) NOT NULL DEFAULT '1' comment '是否开启收藏',
//	  `is_reward` tinyint(1) NOT NULL DEFAULT '1' comment '是否开启打赏',
//	  `status` int(11) NOT NULL DEFAULT '0' comment '文章状态1=启用，2=禁用',
//	  `publish_status` int(11) NOT NULL DEFAULT '0' comment '文章发布状态',
//	  `sort` int(11) NOT NULL DEFAULT '0' comment '文章排序',
//	  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
//	  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
//	  PRIMARY KEY (`id`)
//	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='文章表';

//创建文章和标签关联表
//CREATE TABLE `article_tag` (
//	  `id` int(11) NOT NULL AUTO_INCREMENT,
//	  `article_id` int(11) NOT NULL DEFAULT '0' comment '文章id',
//	  `tag_id` int(11) NOT NULL DEFAULT '0' comment '标签id',
//	  PRIMARY KEY (`id`)
//	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='文章和标签关联表';
