package friend_bill

import (
	monify "monify/protobuf/gen/go"
)

type Service struct {
	monify.UnimplementedFriendBillServiceServer
}
