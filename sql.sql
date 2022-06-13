/*
SQLyog Ultimate v12.4.3 (64 bit)
MySQL - 5.7.25-log : Database - douyin_db
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`douyin_db` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `douyin_db`;

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id(主键)',
  `name` varchar(50) DEFAULT NULL COMMENT '用户名称',
  `password` varchar(80) NOT NULL COMMENT '用户密码',
  `phone` varchar(20) NOT NULL COMMENT '手机号',
  `follow_count` bigint(20) unsigned DEFAULT NULL COMMENT '关注数',
  `follower_count` bigint(20) unsigned DEFAULT NULL COMMENT '粉丝数',
  `is_follow` tinyint(1) DEFAULT NULL COMMENT '是否关注',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

/*Data for the table `user` */

insert  into `user`(`id`,`name`,`password`,`phone`,`follow_count`,`follower_count`,`is_follow`,`create_at`,`update_at`) values 
(1,'kkite','abc123','12345678910',NULL,NULL,NULL,NULL,NULL),
(2,'zhangsan','abc123','12345678911',NULL,NULL,NULL,NULL,NULL),
(3,'daisy','123456','232212312321',100,132303,NULL,NULL,NULL),
(5,'admin','123456','',100,0,0,'2022-06-12 19:02:26','2022-06-12 19:02:26'),
(6,'admin1','123456','',0,0,0,'2022-06-12 19:43:50','2022-06-12 19:43:50'),
(7,'admin11','123456','',0,0,0,'2022-06-13 15:19:00','2022-06-13 15:19:00');

/*Table structure for table `video` */

DROP TABLE IF EXISTS `video`;

CREATE TABLE `video` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '视频id（主键）',
  `author_id` bigint(20) unsigned NOT NULL COMMENT '作者id（外键）',
  `play_url` varchar(100) NOT NULL COMMENT '视频url',
  `cover_url` varchar(100) NOT NULL COMMENT '视频封面url',
  `favorite_count` bigint(20) unsigned DEFAULT '0' COMMENT '点赞数',
  `comment_count` bigint(20) unsigned DEFAULT '0' COMMENT '评论数',
  `is_favorite` tinyint(1) DEFAULT '0' COMMENT '是否点赞',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_video_fk` (`author_id`),
  CONSTRAINT `user_video_fk` FOREIGN KEY (`author_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

/*Data for the table `video` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
