package service

import (
	"context"
	"encoding/json"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
	"gorm.io/gorm"
)

// GetProfile はデータを取得します。
func GetProfile(ctx context.Context, db *gorm.DB) (*entity.ProfileResponse, error) {
	dbProfile, err := repository.GetProfile(ctx, db)
	if err != nil {
		return nil, err
	}

	var bio []string
	if dbProfile.Bio != "" {
		if err := json.Unmarshal([]byte(dbProfile.Bio), &bio); err != nil {
			return nil, err
		}
	}

	var highlights []entity.ProfileHighlight
	if dbProfile.Highlights != "" {
		if err := json.Unmarshal([]byte(dbProfile.Highlights), &highlights); err != nil {
			return nil, err
		}
	}

	var expertise []string
	if dbProfile.Expertise != "" {
		if err := json.Unmarshal([]byte(dbProfile.Expertise), &expertise); err != nil {
			return nil, err
		}
	}

	return &entity.ProfileResponse{
		Name:         dbProfile.Name,
		Subtitle:     dbProfile.Subtitle,
		Title:        dbProfile.Title,
		AvatarURL:    dbProfile.AvatarURL,
		Quote:        dbProfile.Quote,
		Bio:          bio,
		Highlights:   highlights,
		Award:        dbProfile.Award,
		Expertise:    expertise,
		ContactEmail: dbProfile.ContactEmail,
		FinalQuote:   dbProfile.FinalQuote,
	}, nil
}

// UpdateProfile はデータを更新します。
func UpdateProfile(ctx context.Context, db *gorm.DB, req entity.UpdateProfileRequest) (*entity.ProfileResponse, error) {
	bioJSON, err := json.Marshal(req.Bio)
	if err != nil {
		return nil, err
	}

	highlightsJSON, err := json.Marshal(req.Highlights)
	if err != nil {
		return nil, err
	}

	expertiseJSON, err := json.Marshal(req.Expertise)
	if err != nil {
		return nil, err
	}

	dbProfile := entity.DBTableProfile{
		ID:           1,
		Name:         req.Name,
		Subtitle:     req.Subtitle,
		Title:        req.Title,
		AvatarURL:    req.AvatarURL,
		Quote:        req.Quote,
		Bio:          string(bioJSON),
		Highlights:   string(highlightsJSON),
		Award:        req.Award,
		Expertise:    string(expertiseJSON),
		ContactEmail: req.ContactEmail,
		FinalQuote:   req.FinalQuote,
	}

	if err := repository.UpdateProfile(ctx, db, &dbProfile); err != nil {
		return nil, err
	}

	return &entity.ProfileResponse{
		Name:         dbProfile.Name,
		Subtitle:     dbProfile.Subtitle,
		Title:        dbProfile.Title,
		AvatarURL:    dbProfile.AvatarURL,
		Quote:        dbProfile.Quote,
		Bio:          req.Bio,
		Highlights:   req.Highlights,
		Award:        dbProfile.Award,
		Expertise:    req.Expertise,
		ContactEmail: dbProfile.ContactEmail,
		FinalQuote:   dbProfile.FinalQuote,
	}, nil
}
