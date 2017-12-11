CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(255) NOT NULL DEFAULT '',
  `user` char(255) NOT NULL DEFAULT '',
  `image` varchar(8190) DEFAULT 'No Image',
  `age` int(11) DEFAULT '0',
  `birthday` varchar(255) DEFAULT 'No Birthday',
  `email` char(255) NOT NULL DEFAULT '',
  `email_ok` tinyint(1) DEFAULT '0',
  `phone` char(255) DEFAULT 'No Phone',
  `phone_ok` tinyint(1) DEFAULT '0',
  `address` varchar(8190) DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `confirmed_service` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_uuid` char(255) DEFAULT NULL,
  `service_uuid` char(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `users` (`uuid`, `user`, `image`, `age`, `birthday`, `email`, `email_ok`, `phone`, `phone_ok`, `address`)
VALUES
  ('uuid', 'Azunyan1111', 'http://dic.nicovideo.jp/oekaki/185531.png', 16, '1996/11/11', 'azunyan1111@azunyan.me', True, '090-1145-1419', True, '東京都梓市');

INSERT INTO `users` (`uuid`, `user`, `image`, `age`, `birthday`, `email`, `email_ok`, `phone`, `phone_ok`, `address`)
VALUES
  ('uuid2', 'hoge', 'https://11neko.com/wp-content/uploads/2014/10/hoge.jpg', 16, '1990/11/17', 'hoge@hoge.com', True, '090-4545-1919', True, '東京都ほげ市');

INSERT INTO `users` (`uuid`, `user`, `image`, `age`, `birthday`, `email`, `email_ok`, `phone`, `phone_ok`, `address`)
VALUES
  ('uuid3', 'foo', 'https://i.stack.imgur.com/sViXp.png', 16, '1980/1/1', 'foo@foo.com', False, '090-4545-1919', False, '東京都ほげ市');
