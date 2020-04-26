CREATE TABLE IF NOT EXISTS data_api.budget_mitaka (
    id int(11) AUTO_INCREMENT NOT NULL primary key,
    category VARCHAR(45) NOT NULL,
    tax_type1 VARCHAR(255) NOT NULL,
    tax_type2 VARCHAR(255) NOT NULL,
    tax_type3 VARCHAR(255) NOT NULL,
    tax_type4 VARCHAR(255) NOT NULL,
    tax_type5 VARCHAR(255) NOT NULL,
    tax_type6 VARCHAR(255) NOT NULL,
    amount int(11) NOT NULL,
    INDEX(id)
);

LOAD DATA LOCAL INFILE "./data/000_令和二年度予算データ（歳入）" INTO TABLE budget_mitaka FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"';

CREATE TABLE IF NOT EXISTS data_api.town_list (
    id int(11) AUTO_INCREMENT NOT NULL primary key,
    town_name VARCHAR(45) NOT NULL,
    town_ruby VARCHAR(45) NOT NULL,
    INDEX(id)
);

INSERT INTO data_api.town_list(town, town_ruby) VALUES ('三鷹', 'ﾐﾀｶ');