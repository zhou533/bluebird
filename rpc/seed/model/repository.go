package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	conn *gorm.DB
	log  logx.Logger
}

type contextTxKey struct{}

func (repo *Repository) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return repo.conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (repo *Repository) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return repo.conn
}

func NewRepository(datasource string, log logx.Logger) (*Repository, func(), error) {

	cleanup := func() {
		log.Info("closing the database connection")
	}

	conn, err := gorm.Open(postgres.Open(datasource), &gorm.Config{})
	if err != nil {
		log.Errorf("failed opening connection to postgresql: %v", err)
		return nil, nil, err
	}

	if err := conn.AutoMigrate(&Seed{}); err != nil {
		log.Errorf("failed migrating database: %v", err)
		return nil, nil, err
	}

	return &Repository{
		conn: conn,
		log:  log,
	}, cleanup, nil
}
