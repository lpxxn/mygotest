CREATE DATABASE IF NOT EXISTS learn_en CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

use learn_en;

DROP TABLE IF EXISTS cet_6;
create TABLE cet_6 (
    id int NOT NULL AUTO_INCREMENT,
    word varchar(30) not null default '',
    meaning varchar(200) NOT null default '',
    eg varchar(1000) not null default '',
    PRIMARY KEY (id)
    );
CREATE INDEX idx_cet_6_word ON cet_6 (word);