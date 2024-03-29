package role

import (
	"context"

	"github.com/migotom/cell-centre-services/pkg/entities"
	pb "github.com/migotom/cell-centre-services/pkg/pb"
)

type Repository interface {
	Get(ctx context.Context, filter *pb.RoleFilter) (*entities.Role, error)
}
