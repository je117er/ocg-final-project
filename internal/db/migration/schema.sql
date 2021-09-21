CREATE TABLE `product`
(
    `id`                      varchar(50) PRIMARY KEY,
    `name`                    varchar(255),
    `price`                   decimal,
    `vendor`                  varchar(255),
    `vaccine_type`            varchar(100),
    `authorized_ages`         int,
    `dose`                    varchar(255),
    `antigen_nature`          varchar(255),
    `route_of_administration` varchar(255),
    `storage_requirements`    varchar(255),
    `available_formats`       varchar(255),
    `diluent`                 varchar(255),
    `adjuvant`                varchar(255),
    `alternate_name`          varchar(100),
    `minimum_interval`        int,
    `immunization_schedule`   int,
    `authorized_interval`     int,
    `extended_interval`       int,
    `background`              text,
    `regulatory_actions`      text,
    `safety_status`           text,
    `authorization_status`    tinyint(1),
    `trials`                  text,
    `distribution`            text,
    `funding`                 text,
    `slug`                    varchar(100),
    `image`                   varchar(255),
    `lot_number`              varchar(10),
    `expiry_date`             date
);

CREATE TABLE `customer`
(
    `id`               int PRIMARY KEY AUTO_INCREMENT,
    `email`            varchar(100) unique ,
    `name`             varchar(100),
    `address`          varchar(255),
    `gender`           varchar(10),
    `dob`              date,
    `phone_number`     varchar(15),
    `insurance_number` varchar(15),
    `city`             varchar(255),
    `district`         varchar(255),
    `commune`          varchar(60),
    `ethnicity`        varchar(60),
    `nationality`      varchar(25)
);

CREATE TABLE `medical_condition`
(
    `id`               int AUTO_INCREMENT,
    `code`             varchar(255) UNIQUE NOT NULL,
    `description`      text,
    `condition_status` tinyint(1),
    `customer_id`      int                 NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `constraint`
(
    `id`               int PRIMARY KEY AUTO_INCREMENT,
    `code`             varchar(255) UNIQUE NOT NULL,
    `description`      text,
    `vaccine_eligible` tinyint(1),
    `need_precaution`  tinyint(1)
);

CREATE TABLE `product_constraint`
(
    `product_id`    varchar(50) NOT NULL,
    `constraint_id` int NOT NULL,
    PRIMARY KEY (`product_id`, `constraint_id`)
);

CREATE TABLE `clinic`
(
    `id`       varchar(50) PRIMARY KEY,
    `name`     varchar(255),
    `address`  varchar(255),
    `contact`  varchar(255),
    `location` varchar(255),
    `status`   tinyint(1)
);

CREATE TABLE `session_capacity`
(
    `id`                 int PRIMARY KEY AUTO_INCREMENT,
    `clinic_id`          varchar(50),
    `capacity`           int,
    `type`              tinyint,
    `status` tinyint,
    `current_date`       date
);

CREATE TABLE `stock_item`
(
    `id`          varchar(50) PRIMARY KEY NOT NULL,
    `clinic_id`   varchar(50),
    `quantity`    int,
    `name`        varchar(255),
    `price`       decimal,
    `lot_number`  varchar(10),
    `expiry_date` date,
    `product_id`  varchar(50)
);

CREATE TABLE `booking`
(
    `id`                  int AUTO_INCREMENT,
    `customer_id`         int NOT NULL,
    `date_registered`     datetime,
    `date_booked`         date,
    `time_period`         tinyint,
    `doses_completed`     tinyint,
    `vaccine_name`        varchar(255),
    `extended_interval`   int,
    `lot_number`          varchar(10),
    `clinic_name`         varchar(255),
    `price`               decimal,
    `sent_reminder_email` tinyint,
    `daily_capacity_id`   int,
    `total_bill`          decimal,
    `payment_status`      tinyint(1),
    `stock_item_id`       varchar(50),
    `discount_id`         int NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `price_rule`
(
    `id`                int AUTO_INCREMENT,
    `product_id`        varchar(50) NOT NULL,
    `value`             decimal,
    `value_type`        tinyint,
    `once_per_customer` tinyint(1),
    `usage_limit`       int,
    `starts_at`         datetime,
    `ends_at`           datetime,
    `title`             varchar(255),
    PRIMARY KEY (`id`)
);

CREATE TABLE `discount`
(
    `id`            int AUTO_INCREMENT,
    `code`          varchar(255),
    `usage_count`   int,
    `price_rule_id` int NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `admin`
(
    `id`       int PRIMARY KEY AUTO_INCREMENT,
    `username` varchar(25),
    `password` varchar(100)
);

ALTER TABLE `medical_condition`
    ADD FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`);

ALTER TABLE `product_constraint`
    ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);

ALTER TABLE `product_constraint`
    ADD FOREIGN KEY (`constraint_id`) REFERENCES `constraint` (`id`);

ALTER TABLE `booking`
    ADD FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`);

ALTER TABLE `booking`
    ADD FOREIGN KEY (`discount_id`) REFERENCES `discount` (`id`);

ALTER TABLE `price_rule`
    ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);

ALTER TABLE `discount`
    ADD FOREIGN KEY (`price_rule_id`) REFERENCES `price_rule` (`id`);

ALTER TABLE session_capacity
    ADD FOREIGN KEY (`clinic_id`) REFERENCES `clinic` (`id`);

ALTER TABLE `booking`
    ADD FOREIGN KEY (`daily_capacity_id`) REFERENCES session_capacity (`id`);

ALTER TABLE `stock_item`
    ADD FOREIGN KEY (`clinic_id`) REFERENCES `clinic` (`id`);

ALTER TABLE `stock_item`
    ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`id`);

ALTER TABLE `booking`
    ADD FOREIGN KEY (`stock_item_id`) REFERENCES `stock_item` (`id`);

