package background_jobs

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"time"
)

const TypeCheckExpiredCampaignInvitations = "campaign:check_expired_invitations"
const CheckExpiredInvitationsInterval = 1 * time.Hour

type CheckExpiredCampaignInvitationsTask struct {
}

func newCheckExpiredCampaignInvitationsTask() (*asynq.Task, error) {
	return asynq.NewTask(
		TypeCheckExpiredCampaignInvitations,
		nil,
		asynq.Unique(1*time.Hour),
		asynq.TaskID(TypeCheckExpiredCampaignInvitations),
	), nil
}

func HandleCheckExpiredCampaignInvitationsTaskWrapper(repo sqlc.Querier) func(ctx context.Context, t *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {
		_, err := repo.Worker_UpdateStatusOfExpiredCampaignInvitation(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}

func RegisterCheckExpiredCampaignInvitationsTask(
	scheduler *asynq.Scheduler,
	mux *asynq.ServeMux,
	repo sqlc.Querier,
) error {
	task, err := newCheckExpiredCampaignInvitationsTask()
	if err != nil {
		return err
	}

	spec := fmt.Sprintf("@every %v", CheckExpiredInvitationsInterval)

	mux.HandleFunc(
		TypeCheckExpiredCampaignInvitations,
		HandleCheckExpiredCampaignInvitationsTaskWrapper(repo),
	)

	_, err = scheduler.Register(
		spec,
		task,
		asynq.Unique(1*time.Hour),
		asynq.TaskID(TypeCheckExpiredCampaignInvitations),
	)
	return err
}
