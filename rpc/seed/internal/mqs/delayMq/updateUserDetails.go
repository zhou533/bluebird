package delayMq

import (
	"bluebird/rpc/seed/internal/svc"
	"context"
	"database/sql"

	"github.com/hibiken/asynq"
)

type UpdateUserDetailsHandler struct {
	svcCtx *svc.ServiceContext
}

func NewUpdateUserDetailsHandler(svcCtx *svc.ServiceContext) *UpdateUserDetailsHandler {
	return &UpdateUserDetailsHandler{
		svcCtx: svcCtx,
	}
}

func (h *UpdateUserDetailsHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	// implement the business logic of the task
	seeds, err := h.svcCtx.SeedModel.FindUsersByStatus(ctx, 0)
	if err != nil {
		// TODO log
		return err
	}

	if len(seeds) == 0 {
		return nil
	}

	if len(seeds) > 100 {
		seeds = seeds[:100]
	}

	var usernames []string
	for _, seed := range seeds {
		usernames = append(usernames, seed.UserName)
	}

	users, err := h.svcCtx.TwitterClient.LookupUser(usernames)
	if err != nil {
		// TODO log
		return err
	}

	for _, user := range users {
		seed, err := h.svcCtx.SeedModel.FindOneByUserName(ctx, user.UserName)
		if err != nil {
			// TODO log
			continue
		}

		seed.TwitterId = sql.NullString{String: user.ID, Valid: len(user.ID) > 0}
		seed.Name = sql.NullString{String: user.Name, Valid: len(user.Name) > 0}
		seed.Location = sql.NullString{String: user.Location, Valid: len(user.Location) > 0}
		seed.Url = sql.NullString{String: user.URL, Valid: len(user.URL) > 0}
		seed.Description = sql.NullString{String: user.Description, Valid: len(user.Description) > 0}
		seed.ProfileImageUrl = sql.NullString{String: user.ProfileImageURL, Valid: len(user.ProfileImageURL) > 0}

		seed.Status = 1

		err = h.svcCtx.SeedModel.Update(ctx, seed)
		if err != nil {
			// TODO log
			continue
		}
	}
	return nil
}
