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

func (s Service) ListFriendBill(ctx context.Context, req *monify.ListFriendBillRequest) (*monify.ListFriendBillResponse, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	_, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)
	query, err := db.QueryContext(ctx, `
		SELECT friend_bill_id, amount, title, description
		FROM friend_bill
		WHERE relation_id = $1`, req.RelationId)
	if err != nil {
		logger.Error("Select friend bill information error.", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	defer query.Close()

	var friend_bills []*monify.FriendBill
	for {
		if !query.Next() {
			break
		}
		var friend_bill monify.FriendBill
		if err = query.Scan(&friend_bill.FriendBillId, &friend_bill.Amount, &friend_bill.Title, &friend_bill.Description); err != nil {
			logger.Error("Scan friend bill information error.", zap.Error(err))
			return nil, status.Error(codes.Internal, "")
		}
		friend_bills = append(friend_bills, &friend_bill)
	}

	return &monify.ListFriendBillResponse{
		FriendBills: friend_bills,
	}, nil
}
