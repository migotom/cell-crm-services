package event

import (
	"context"

	"github.com/migotom/cell-centre-services/pkg/entities"
)

// Repository of events.
type Repository interface {
	New(ctx context.Context, event *entities.Event) error
}
