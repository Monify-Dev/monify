package friend_bill

import (
	"context"
	"database/sql"
	"monify/lib"
	monify "monify/protobuf/gen/go"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s Service) DeleteFriendBill(ctx context.Context, req *monify.DeleteFriendBillRequest) (*emptypb.Empty, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	_, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	_, err := db.ExecContext(ctx, `DELETE FROM friend_bill WHERE friend_bill_id = $1`, req.FriendBillId)
	if err != nil {
		logger.Error("Delete friend bill error.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &emptypb.Empty{}, nil
}
