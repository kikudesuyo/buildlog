package service

import (
	"context"
	"encoding/json"

	"github.com/kikudesuyo/buildlog/api/entity"
	"github.com/kikudesuyo/buildlog/api/repository"
)

func GetProfile(ctx context.Context) (*entity.ProfileResponse, error) {
	if cached, ok := contentCache.Get("profile"); ok {
		return cloneProfileResponse(cached.(*entity.ProfileResponse)), nil
	}

	dbProfile, err := repository.GetProfile(ctx, database)
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

	profile := &entity.ProfileResponse{
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
	}
	contentCache.Set("profile", cloneProfileResponse(profile))
	return profile, nil
}

func cloneProfileResponse(profile *entity.ProfileResponse) *entity.ProfileResponse {
	if profile == nil {
		return nil
	}

	clone := *profile
	clone.Bio = append([]string(nil), profile.Bio...)
	clone.Highlights = append([]entity.ProfileHighlight(nil), profile.Highlights...)
	clone.Expertise = append([]string(nil), profile.Expertise...)
	return &clone
}

func UpdateProfile(ctx context.Context, req entity.UpdateProfileRequest) (*entity.ProfileResponse, error) {
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

	if err := repository.UpdateProfile(ctx, database, &dbProfile); err != nil {
		return nil, err
	}
	contentCache.Delete("profile")

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
