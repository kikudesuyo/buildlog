package service

import (
	"context"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

func GetProfile(ctx context.Context, db *gorm.DB) (*entity.DBTableProfile, error) {
	return repository.GetProfile(ctx, db)
}

func UpdateProfile(ctx context.Context, db *gorm.DB, req entity.UpdateProfileRequest) (*entity.DBTableProfile, error) {
	profile := entity.DBTableProfile{
		ID:           1,
		Name:         req.Name,
		Subtitle:     req.Subtitle,
		Title:        req.Title,
		AvatarURL:    req.AvatarURL,
		Quote:        req.Quote,
		Bio:          req.Bio,
		Highlights:   req.Highlights,
		Award:        req.Award,
		Expertise:    req.Expertise,
		ContactEmail: req.ContactEmail,
		FinalQuote:   req.FinalQuote,
	}

	if err := repository.UpdateProfile(ctx, db, &profile); err != nil {
		return nil, err
	}

	return &profile, nil
}
