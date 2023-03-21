package model

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
)

type (
	Seed struct {
		Id              int64          `gorm:"column:id;primaryKey"`
		TwitterId       sql.NullInt64  `gorm:"column:twitter_id"`
		Name            sql.NullString `gorm:"column:name"`
		ScreenName      string         `gorm:"column:screen_name;index;not mull"`
		Location        sql.NullString `gorm:"column:location"`
		Url             sql.NullString `gorm:"column:url"`
		Description     sql.NullString `gorm:"column:description"`
		ProfileImageUrl sql.NullString `gorm:"column:profile_image_url"`
		LatestTweetId   sql.NullInt64  `gorm:"column:latest_tweet_id"`
		LatestTweetAt   sql.NullTime   `gorm:"column:latest_tweet_at"`
		CreateAt        int64          `gorm:"column:create_at;autoCreateTime:milli"`
		UpdateAt        int64          `gorm:"column:update_at;autoUpdateTime:milli"`
		Status          int64          `gorm:"column:status;default:0"`
	}

	SeedModel interface {
		Insert(ctx context.Context, data *Seed) error
		InsertScreenName(ctx context.Context, screenName string) error

		FindOne(ctx context.Context, id int64) (*Seed, error)
		FindOneByScreenName(ctx context.Context, screenName string) (*Seed, error)
		Update(ctx context.Context, data *Seed) error
	}

	SeedRepo struct {
		repo *Repository
		log  logx.Logger
	}
)

func (Seed) TableName() string {
	return "seed"
}

func NewSeedModel(repo *Repository, log logx.Logger) SeedModel {
	return &SeedRepo{
		repo: repo,
		log:  log,
	}
}

func (sr *SeedRepo) Insert(ctx context.Context, seed *Seed) error {
	result := sr.repo.DB(ctx).Create(seed)
	return result.Error
}

func (sr *SeedRepo) InsertScreenName(ctx context.Context, screenName string) error {
	result := sr.repo.DB(ctx).Create(&Seed{ScreenName: screenName})
	return result.Error
}

func (sr *SeedRepo) FindOne(ctx context.Context, id int64) (*Seed, error) {
	var seed Seed
	result := sr.repo.DB(ctx).First(&seed, id)
	return &seed, result.Error
}

func (sr *SeedRepo) FindOneByScreenName(ctx context.Context, screenName string) (*Seed, error) {
	var seed Seed
	result := sr.repo.DB(ctx).Where("screen_name = ?", screenName).First(&seed)
	return &seed, result.Error
}

func (sr *SeedRepo) Update(ctx context.Context, seed *Seed) error {
	result := sr.repo.DB(ctx).Save(seed)
	return result.Error
}
