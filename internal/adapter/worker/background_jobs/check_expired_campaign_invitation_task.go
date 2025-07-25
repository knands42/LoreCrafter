package background_jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	sqlc "github.com/knands42/lorecrafter/pkg/sqlc/generated"
	"time"
)

const TypeCheckExpiredCampaignInvitations = "campaign:check_expired_invitations"
const CheckExpiredInvitationsInterval = 12 * time.Hour

type CheckExpiredCampaignInvitationsTask struct {
}

func newCheckExpiredCampaignInvitationsTask() (*asynq.Task, error) {
	payload, err := json.Marshal(CheckExpiredCampaignInvitationsTask{})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeCheckExpiredCampaignInvitations, payload), nil
}

func HandleCheckExpiredCampaignInvitationsTaskWrapper(repo sqlc.Querier) func(ctx context.Context, t *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {
		var task CheckExpiredCampaignInvitationsTask
		if err := json.Unmarshal(t.Payload(), &task); err != nil {
			return err
		}

		_, err := repo.Worker_UpdateStatusOfExpiredCampaignInvitation(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}

func RegisterCheckExpiredCampaignInvitationsTask(scheduler *asynq.Scheduler) error {
	task, err := newCheckExpiredCampaignInvitationsTask()
	if err != nil {
		return err
	}

	spec := fmt.Sprintf("@every %v", CheckExpiredInvitationsInterval)

	_, err = scheduler.Register(spec, task)
	return err
}
