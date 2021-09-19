CREATE TABLE `product` (
                           `id` binary(16) PRIMARY KEY,
                           `price` decimal,
                           `sku` varchar(10),
                           `vaccine_type` tinyint(1),
                           `authorized_ages` int,
                           `dose` varchar(255),
                           `antigen_nature` varchar(255),
                           `route_of_administration` varchar(255),
                           `storage_requirements` varchar(255),
                           `available_formats` varchar(255),
                           `diluent` varchar(255),
                           `adjuvant` varchar(255),
                           `alternate_name` varchar(100),
                           `minimum_interval` varchar(255),
                           `authorized_interval` varchar(255),
                           `extended_interval` varchar(255),
                           `background` text,
                           `regulatory_actions` text,
                           `safety_status` text,
                           `authorization_status` tinyint(1),
                           `trials` text,
                           `distribution` text,
                           `funding` text,
                           `slug` varchar(100),
                           `image` varchar(255),
                           `lot_number` varchar(10),
                           `expiry_date` date
);

CREATE TABLE `vendor` (
                          `id` int PRIMARY KEY AUTO_INCREMENT,
                          `name` varchar(255),
                          `contact` varchar(255),
                          `address` varchar(255),
                          `country` varchar(255)
);

CREATE TABLE `product_vendor` (
                                  `vendor_id` int NOT NULL,
                                  `product_id` binary(16) NOT NULL,
                                  PRIMARY KEY (`vendor_id`, `product_id`)
);

CREATE TABLE `customer` (
                            `id` int PRIMARY KEY AUTO_INCREMENT,
                            `email` varchar(100),
                            `name` varchar(100),
                            `address` varchar(255),
                            `gender` varchar(10),
                            `dob` date,
                            `phone_number` varchar(15),
                            `insurance_number` varchar(15),
                            `city` varchar(255),
                            `district` varchar(255),
                            `commune` varchar(60),
                            `ethnicity` varchar(60),
                            `nationality` varchar(25)
);

CREATE TABLE `medical_condition` (
                                     `id` int PRIMARY KEY AUTO_INCREMENT,
                                     `code` varchar(255),
                                     `description` text,
                                     `condition_status` tinyint(1),
                                     `customer_id` int NOT NULL ,
                                     `constraint_id` int NOT NULL
);

CREATE TABLE `constraint` (
                              `id` int PRIMARY KEY AUTO_INCREMENT,
                              `code` varchar(255),
                              `vaccine_eligible` tinyint(1),
                              `need_precaution` tinyint(1)
);

CREATE TABLE `product_constraint` (
                                      `product_id` binary(16) NOT NULL ,
                                      `constraint_id` int NOT NULL ,
                                      `quantity` int,
                                      PRIMARY KEY (`product_id`, `constraint_id`)
);

CREATE TABLE `clinic` (
                          `id` binary(16) PRIMARY KEY,
                          `name` varchar(255),
                          `address` varchar(255),
                          `contact` varchar(255),
                          `location` varchar(255),
                          `status` tinyint(1)
);

CREATE TABLE `daily_capacity` (
                                  `id` int PRIMARY KEY AUTO_INCREMENT,
                                  `clinic_id` binary(16),
                                  `morning_capacity` int,
                                  `afternoon_capacity` int,
                                  `evening_capacity` int,
                                  `current_date` date
);

CREATE TABLE `clinic_product` (
                                  `product_id` binary(16) NOT NULL ,
                                  `clinic_id` binary(16) NOT NULL ,
                                  PRIMARY KEY (`product_id`, `clinic_id`)
);

CREATE TABLE `booking` (
                           `id` int AUTO_INCREMENT,
                           `customer_id` int NOT NULL,
                           `date_registered` datetime,
                           `date_booked` date,
                           `time_period` tinyint,
                           `doses_completed` tinyint,
                           `vaccine_picked` tinyint,
                           `clinic_name` varchar(255),
                           `sent_reminder_email` tinyint,
                           `daily_capacity_id` int,
                           `total_bill` decimal,
                           `payment_status` tinyint(1),
                           `product_id` binary(16),
                           `discount_id` int NOT NULL,
                           PRIMARY KEY (`id`, `customer_id`, `discount_id`)
);

CREATE TABLE `price_rule` (
                              `id` int AUTO_INCREMENT,
                              `product_id` binary(16) NOT NULL,
                              `value` decimal,
                              `value_type` tinyint,
                              `once_per_customer` tinyint(1),
                              `usage_limit` int,
                              `starts_at` datetime,
                              `ends_at` datetime,
                              `title` varchar(255),
                              PRIMARY KEY (`id`, `product_id`)
);

CREATE TABLE `discount` (
                            `id` int AUTO_INCREMENT,
                            `code` varchar(255),
                            `usage_count` int,
                            `price_rule_id` int NOT NULL,
                            PRIMARY KEY (`id`, `price_rule_id`)
);

CREATE TABLE `admin` (
                         `id` int PRIMARY KEY AUTO_INCREMENT,
                         `username` varchar(25),
                         `password_hash` varchar(100)
);

ALTER TABLE `product_vendor` ADD FOREIGN KEY (`vendor_id`) REFERENCES `vendor` (`id`);

ALTER TABLE `product_vendor` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);

ALTER TABLE `medical_condition` ADD FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`);

ALTER TABLE `medical_condition` ADD FOREIGN KEY (`constraint_id`) REFERENCES `constraint` (`id`);

ALTER TABLE `product_constraint` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);

ALTER TABLE `product_constraint` ADD FOREIGN KEY (`constraint_id`) REFERENCES `constraint` (`id`);

ALTER TABLE `clinic_product` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);

ALTER TABLE `clinic_product` ADD FOREIGN KEY (`clinic_id`) REFERENCES `clinic` (`id`);

ALTER TABLE `booking` ADD FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`);

ALTER TABLE `booking` ADD FOREIGN KEY (`discount_id`) REFERENCES `discount` (`id`);

ALTER TABLE `price_rule` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);

ALTER TABLE `discount` ADD FOREIGN KEY (`price_rule_id`) REFERENCES `price_rule` (`id`);

ALTER TABLE `daily_capacity` ADD FOREIGN KEY (`clinic_id`) REFERENCES `clinic` (`id`);

ALTER TABLE `booking` ADD FOREIGN KEY (`daily_capacity_id`) REFERENCES `daily_capacity` (`id`);

ALTER TABLE `booking` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);
