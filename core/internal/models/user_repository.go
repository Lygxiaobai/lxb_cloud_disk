package models

import "time"

type UserRepository struct {
	Id                 int
	Identity           string
	ParentId           int64
	UserIdentity       string
	RepositoryIdentity string
	Name               string
	Ext                string    //文件扩展名
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
	DeletedAt          time.Time `xorm:"deleted"`
}

func (u UserRepository) TableName() string {
	return "user_repository"
}
