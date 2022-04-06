
SET NAMES utf8mb4;
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL';

DROP SCHEMA IF EXISTS db_typ_engineer;
CREATE SCHEMA db_typ_engineer;
USE db_typ_engineer;

CREATE TABLE `words` (
  `word` varchar(256) NOT NULL,
  `parts_of_speech` varchar(256),
  `discription` varchar(256) NOT NULL,
  `level_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT false,
  PRIMARY KEY (`word`)
) ENGINE=InnoDB AUTO_INCREMENT=8;

CREATE TABLE `levels` (
  `level_id` int(11) NOT NULL,
  `level` varchar(256) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT false,
  PRIMARY KEY (`level_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8;