CREATE TABLE
  `students` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `fname` varchar(50) NOT NULL,
    `lname` varchar(50) NOT NULL,
    `date_of_birth` datetime NOT NULL,
    `email` varchar(50) NOT NULL,
    `address` varchar(50) NOT NULL,
    `gender` varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
  ) ENGINE = InnoDB AUTO_INCREMENT = 7 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci