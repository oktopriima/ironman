-- +migrate Up
CREATE TABLE IF NOT EXISTS `mkproject`.`user_socials`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     INT UNSIGNED NOT NULL,
    `social_name` VARCHAR(255) NOT NULL,
    `social_id`   VARCHAR(255) NOT NULL,
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `fk_user_social_users1_idx` (`user_id` ASC),
    CONSTRAINT `fk_user_social_users1`
        FOREIGN KEY (`user_id`)
            REFERENCES `mkproject`.`users` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `mkproject`.`user_socials`;