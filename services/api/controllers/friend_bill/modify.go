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

func (s Service) ModifyFriendBill(ctx context.Context, req *monify.ModifyFriendBillRequest) (*emptypb.Empty, error) {
	logger := ctx.Value(lib.LoggerContextKey{}).(*zap.Logger)
	_, ok := ctx.Value(lib.UserIdContextKey{}).(uuid.UUID)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized.")
	}
	db := ctx.Value(lib.DatabaseContextKey{}).(*sql.DB)

	//START transaction
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		logger.Error("Failed to begin transaction", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	defer tx.Rollback()

	// Delete
	_, err = tx.ExecContext(ctx, `DELETE FROM friend_bill WHERE friend_bill_id = $1`, req.FriendBillId)
	if err != nil {
		logger.Error("Delete friend bill error. (Modify)", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	// Insert
	_, err = tx.ExecContext(ctx, `
		INSERT INTO friend_bill (friend_bill_id, relation_id, amount, title, description)
		VALUES ($1, $2, $3, $4, $5)
	`, req.FriendBillId, req.RelationId, req.Amount, req.Title, req.Description)
	if err != nil {
		logger.Error("Insert values into friend_bill error. (Modify)", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	//Commit
	if err = tx.Commit(); err != nil {
		logger.Error("Failed to commit transaction", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal")
	}
	return &emptypb.Empty{}, nil
}
