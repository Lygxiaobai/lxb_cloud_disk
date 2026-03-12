// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"cloud_disk/core/internal/config"
	"cloud_disk/core/internal/models"
	"github.com/redis/go-redis/v9"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.Mysql.DataSource),
		RDB:    models.InitRedis(c.Redis.Addr),
	}
}
