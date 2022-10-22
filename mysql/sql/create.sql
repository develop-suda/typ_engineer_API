
SET NAMES utf8mb4;
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL';

DROP SCHEMA IF EXISTS db_typ_engineer;
CREATE SCHEMA db_typ_engineer;
USE db_typ_engineer;

CREATE TABLE `word_types` (
  `word_type_id` int(255) NOT NULL UNIQUE,
  `word_type` varchar(256) NOT NULL UNIQUE,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT false,
  PRIMARY KEY (`word_type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8;

CREATE TABLE `parts_of_speeches` (
  `parts_of_speech_id` int(255) NOT NULL UNIQUE,
  `parts_of_speech` varchar(256) NOT NULL UNIQUE,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT false,
  PRIMARY KEY (`parts_of_speech_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8;

CREATE TABLE `words` (
  `word` varchar(256) NOT NULL UNIQUE,
  `description` varchar(256) NOT NULL,
  `parts_of_speech_id` int(255) DEFAULT NULL,
  `word_type_id` int(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT false,
  PRIMARY KEY (`word`),
  FOREIGN KEY (`parts_of_speech_id`) REFERENCES `parts_of_speeches` (`parts_of_speech_id`),
  FOREIGN KEY (`word_type_id`) REFERENCES `word_types` (`word_type_id`),
  INDEX (`word`)
) ENGINE=InnoDB AUTO_INCREMENT=8;

CREATE TABLE `users` (
  `user_id` INT(8) UNSIGNED NOT NULL UNIQUE AUTO_INCREMENT,
  `last_name` varchar(256) NOT NULL,
  `first_name` varchar(256) NOT NULL,
  `email` varchar(256) NOT NULL UNIQUE,
  `password` varchar(256) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_id`),
  INDEX (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8;

CREATE TABLE `login_history` (
  `user_id`int(8) UNSIGNED NOT NULL,
  `login_date` datetime NOT NULL,
  `logout_date` datetime,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT false,
  PRIMARY KEY (`user_id`,`login_date`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8;

CREATE TABLE `typing_alphabet_informations` (
  `user_id`int(8) UNSIGNED NOT NULL,
  `alphabet` varchar(1) NOT NULL,
  `typing_count` int(255) NOT NULL,
  `typing_miss_count` int(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_id`,`alphabet`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
  INDEX (`user_id`,`alphabet`)
) ENGINE=InnoDB AUTO_INCREMENT=8;

CREATE TABLE `typing_word_informations` (
  `user_id`int(8) UNSIGNED NOT NULL,
  `word` varchar(256) NOT NULL,
  `typing_count` int(255) NOT NULL,
  `typing_miss_count` int(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_id`,`word`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
  FOREIGN KEY (`word`) REFERENCES `words` (`word`),
  INDEX (`user_id`,`word`)
) ENGINE=InnoDB AUTO_INCREMENT=8;

CREATE TABLE `my_options` (
  `option_id` int(255) NOT NULL AUTO_INCREMENT UNIQUE,
  `user_id`int(8) UNSIGNED NOT NULL,
  `parts_of_speech_id` int(255) NOT NULL,
  `word_type_id` int(255) NOT NULL,
  `alphabet` varchar(1) NOT NULL,
  `quantity` int(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`option_id`),
  FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8;

CREATE VIEW v_words AS SELECT * FROM words WHERE words.is_deleted = false;
CREATE VIEW v_parts_of_speeches AS SELECT * FROM parts_of_speeches WHERE parts_of_speeches.is_deleted = false;
CREATE VIEW v_word_types AS SELECT * FROM word_types WHERE word_types.is_deleted = false;
CREATE VIEW v_typing_word_informations AS SELECT * FROM typing_word_informations WHERE typing_word_informations.is_deleted = false;
CREATE VIEW v_typing_alphabet_informations AS SELECT * FROM typing_alphabet_informations WHERE typing_alphabet_informations.is_deleted = false;
CREATE VIEW v_users AS SELECT * FROM users WHERE users.is_deleted = false;
CREATE VIEW v_login_history AS SELECT * FROM login_history WHERE login_history.is_deleted = false;
CREATE VIEW v_my_options AS SELECT * FROM my_options WHERE my_options.is_deleted = false;