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
)

func (s Service) CreateFriendBill(ctx context.Context, req *monify.CreateFriendBillRequest) (*monify.CreateFriendBillResponse, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	if req.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "Title is required")
	}
	if req.Amount == 0 {
		return nil, status.Error(codes.InvalidArgument, "Amount should not be zero")
	}
	_, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}

	//Insert
	friend_billId := uuid.New()
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	_, err := db.ExecContext(ctx, `
		INSERT INTO friend_bill (friend_bill_id, relation_id, amount, title, description)
		VALUES ($1, $2, $3, $4, $5)
	`, friend_billId, req.RelationId, req.Amount, req.Title, req.Description)
	if err != nil {
		logger.Error("Insert values into friend_bill error.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &monify.CreateFriendBillResponse{
		FriendBillId: friend_billId.String(),
	}, nil

}
