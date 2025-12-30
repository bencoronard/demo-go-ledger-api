package config

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(lc fx.Lifecycle, props *Properties) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",
		props.Secret.DB.Host,
		props.Secret.DB.Port,
		props.Secret.DB.User,
		props.Secret.DB.Pass,
		props.Secret.DB.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(props.Env.CP.ConnectionPoolCap)
	sqlDB.SetMaxIdleConns(props.Env.CP.ConnectionPoolIdleMin)
	sqlDB.SetConnMaxLifetime(time.Duration(props.Env.CP.ConnectionTTL) * time.Millisecond)
	sqlDB.SetConnMaxIdleTime(time.Duration(props.Env.CP.ConnectionPoolIdleTimeout) * time.Millisecond)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return sqlDB.PingContext(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return sqlDB.Close()
		},
	})

	return db, nil
}
