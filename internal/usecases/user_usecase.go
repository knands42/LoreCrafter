package usecases

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/knands42/lorecrafter/internal/domain"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"log"
)

type UserUseCase struct {
	ctx  context.Context
	repo sqlc.Querier
}

func NewUserUseCase(ctx context.Context, repo sqlc.Querier) *UserUseCase {
	return &UserUseCase{
		ctx:  ctx,
		repo: repo,
	}
}

func (uc *UserUseCase) GetUserInfo(userID uuid.UUID) (domain.User, error) {
	userReturned, err := uc.repo.GetUserByID(uc.ctx, pgtype.UUID{
		Bytes: userID,
		Valid: true,
	})
	if err != nil && err.Error() == "no rows in result set" {
		log.Printf("userNotFound %v", err)
		return domain.User{}, ErrUserNotFound
	} else if err != nil {
		log.Printf("error creating invite %v", err)
		return domain.User{}, ErrCreatingTheCampaignInvitation
	}

	return domain.FromSqlcUserToDomain(userReturned), nil
}
