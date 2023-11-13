CREATE TABLE `gpt_ask` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `timestamp` int DEFAULT NULL,
                           `ask_user` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `ask_bot_question` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `bot_answer` varchar(10000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `use_token` int DEFAULT NULL,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;