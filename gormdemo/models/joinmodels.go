package models

import "time"

type JoinModelBase struct {
	ID         string     `gorm:"type:varchar(36);primary_key"`
	CreatedAt  time.Time  `gorm:"type:datetime;"`
	UpdateTime *time.Time `gorm:"type:datetime;" sql:"DEFAULT:current_timestamp"`
	TestTime   time.Time  `gorm:"type:datetime;" sql:"DEFAULT:NOW()"`
}

type Language struct {
	JoinModelBase
	Name string
}

type Movie struct {
	JoinModelBase
	Title      string
	LanguageID string
	Language   Language
}

type Artist struct {
	JoinModelBase
	Name   string
	Movies []Movie `gorm:"many2many:artist_movies"`
}

// db
//CREATE TABLE `gotest_languages` (
//`id` varchar(36) NOT NULL,
//`created_at` datetime DEFAULT NULL,
//`update_time` datetime DEFAULT CURRENT_TIMESTAMP,
//`test_time` datetime DEFAULT CURRENT_TIMESTAMP,
//`name` varchar(255) DEFAULT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;
