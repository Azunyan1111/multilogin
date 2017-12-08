CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(255) DEFAULT '',
  `user` char(255) DEFAULT '',
  `image` varchar(8190) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `birthday` varchar(255) DEFAULT NULL,
  `email` char(255) DEFAULT '',
  `email_ok` tinyint(1) DEFAULT NULL,
  `phone` char(255) DEFAULT NULL,
  `phone_ok` tinyint(1) DEFAULT NULL,
  `address` varchar(8190) DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

CREATE TABLE `service` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(255) DEFAULT NULL,
  `name` char(255) DEFAULT NULL,
  `url` varchar(8190) DEFAULT NULL,
  `url_handler` varchar(8190) DEFAULT NULL,
  `token` char(255) DEFAULT NULL,
  `secret` char(255) DEFAULT NULL,
  `p_name` tinyint(1) DEFAULT NULL,
  `p_image` tinyint(1) DEFAULT NULL,
  `p_age` tinyint(1) DEFAULT NULL,
  `p_birthday` tinyint(1) DEFAULT NULL,
  `p_email` tinyint(1) DEFAULT NULL,
  `p_phone` tinyint(1) DEFAULT NULL,
  `p_address` tinyint(1) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `confirmed_service` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_uuid` char(255) DEFAULT NULL,
  `service_uuid` char(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;