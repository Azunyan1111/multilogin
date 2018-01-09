CREATE TABLE `users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(255) NOT NULL DEFAULT '',
  `user` char(255) NOT NULL DEFAULT '',
  `image` varchar(8190) DEFAULT 'http://noimage.com',
  `age` int(11) DEFAULT '0',
  `birthday` varchar(255) DEFAULT '1900-1-1',
  `email` char(255) NOT NULL DEFAULT '',
  `email_ok` tinyint(1) DEFAULT '0',
  `phone` char(255) DEFAULT '090-000-000',
  `phone_ok` tinyint(1) DEFAULT '0',
  `address` varchar(8190) DEFAULT '住所不明',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `service` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(255) NOT NULL DEFAULT '',
  `name` char(255) NOT NULL DEFAULT '',
  `email` char(255) NOT NULL DEFAULT '',
  `url` varchar(8190) NOT NULL DEFAULT 'no Url',
  `url_callback` varchar(8190) NOT NULL DEFAULT 'no Callback',
  `token` char(255) NOT NULL DEFAULT 'no token',
  `secret` char(255) NOT NULL DEFAULT 'no secret',
  `p_name` tinyint(1) NOT NULL DEFAULT '0',
  `p_image` tinyint(1) NOT NULL DEFAULT '0',
  `p_age` tinyint(1) NOT NULL DEFAULT '0',
  `p_birthday` tinyint(1) NOT NULL DEFAULT '0',
  `p_email` tinyint(1) NOT NULL DEFAULT '0',
  `p_phone` tinyint(1) NOT NULL DEFAULT '0',
  `p_address` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `confirmed_service` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_uuid` char(255) DEFAULT NULL,
  `service_uuid` char(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `users` (`id`, `uuid`, `user`, `image`, `age`, `birthday`, `email`, `email_ok`, `phone`, `phone_ok`, `address`, `created_at`, `updated_at`, `deleted_at`)
VALUES
  (1, '26d2983e-3d5a-421c-bf6f-d4608025e555', 'Azunyan1111', 'http://noimage.com/azunyan', 76, '2017-08-01', 'azunyan1111@azunyan.me', 1, '090-1145-1419', 1, 'fukuoka', '2017-12-31 13:37:06', '2018-01-05 21:04:23', NULL),
  (2, 'uuid2', 'hoge', 'https://11neko.com/wp-content/uploads/2014/10/hoge.jpg', 16, '1990-11-17', 'hoge@hoge.com', 1, '090-4545-1919', 1, 'tokyo hoge city', '2017-12-11 13:16:28', '2018-01-03 16:31:37', NULL),
  (3, 'uuid3', 'foo', 'https://i.stack.imgur.com/sViXp.png', 16, '1980-1-1', 'foo@foo.com', 0, '090-4545-1919', 0, 'tokyo', '2017-12-11 13:16:30', '2018-01-03 16:30:12', NULL);

INSERT INTO `confirmed_service` (`id`, `user_uuid`, `service_uuid`, `created_at`, `updated_at`, `deleted_at`)
VALUES
  (1, 'daa8123d-be45-4574-88d7-339b145396fc', '025ad602-7dba-4c08-8226-704b65f2873c', NULL, '2018-01-04 21:38:44', NULL),
  (2, 'daa8123d-be45-4574-88d7-339b145396fc', '124ah368-1eha-7h81-2345-365a24h6522y', NULL, '2018-01-03 16:31:07', NULL),
  (3, '26d2983e-3d5a-421c-bf6f-d4608025e555', '124ah368-1eha-7h81-2345-365a24h6522y', NULL, '2018-01-04 21:40:35', NULL),
  (4, '26d2983e-3d5a-421c-bf6f-d4608025e555', '025ad602-7dba-4c08-8226-704b65f2873c', '2018-01-03 16:31:17', '2018-01-03 16:31:16', '2018-01-03 16:31:17');

INSERT INTO `service` (`id`, `uuid`, `name`, `email`, `url`, `url_callback`, `token`, `secret`, `p_name`, `p_image`, `p_age`, `p_birthday`, `p_email`, `p_phone`, `p_address`, `created_at`, `updated_at`, `deleted_at`)
VALUES
  (42, '025ad602-7dba-4c08-8226-704b65f2873c', 'GodService', 'god@god.com', 'http://bar.com', 'http://bar.com/callback36', '684797be-0e32-4072-8d8c-fa753bf59b03', 'c1db2a73-782a-43f3-9973-c95e8abb7191', 1, 1, 1, 1, 1, 1, 1, '2018-01-02 11:00:02', '2018-01-05 21:04:23', NULL),
  (45, '124ah368-1eha-7h81-2345-365a24h6522y', 'CoolService', 'bar@bar.com', 'http://bar.com', 'http://bar.com/callback', '684797be-0e32-4072-8d8c-fa753bf59b02', 'c1db2a73-782a-43f3-9973-c95e8abb7190', 1, 1, 1, 1, 1, 1, 1, '2017-12-12 13:20:22', '2018-01-05 14:22:24', NULL),
  (188, '5cca5c22-8cfc-4539-89f5-bc7833546d92', 'a', 'a@a.com', '', '', '2f0a751d-fbf3-4ff2-bb6a-5f1b46f9d65b', '2eab2e1b-2bfb-45d1-a558-1909259b41ea', 0, 0, 1, 0, 1, 0, 0, '2018-01-02 10:53:20', '2018-01-02 10:53:20', NULL);
