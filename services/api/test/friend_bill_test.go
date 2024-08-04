package test

import (
	"context"
	monify "monify/protobuf/gen/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFriendBill(t *testing.T) {
	// Set a friend relation
	client := GetTestClient(t)
	user1 := client.CreateTestUser()
	_, err := client.UpdateUserNickId(context.Background(), &monify.UpdateUserNickIdRequest{NickId: "friend_bill_nickId1"})
	assert.NoError(t, err)
	user2 := client.CreateTestUser()
	_, err = client.UpdateUserNickId(context.Background(), &monify.UpdateUserNickIdRequest{NickId: "friend_bill_nickId2"})
	_, err = client.InviteFriend(context.TODO(), &monify.InviteFriendRequest{ReceiverNickId: "friend_bill_nickId1"})
	assert.NoError(t, err)
	client.SetTestUser(user1)
	invitaions, err := client.ListFriendInvitation(context.TODO(), &monify.FriendEmpty{})
	assert.NoError(t, err)
	_, err = client.AcceptInvitation(context.TODO(), &monify.AcceptInvitationRequest{User1Id: user1, User2Id: user2, InviteId: invitaions.Invitation[0].InviteId})
	assert.NoError(t, err)
	friends, err := client.ListFriend(context.TODO(), &monify.FriendEmpty{})
	assert.NoError(t, err)
	var friendS_relationId []string
	for _, friend := range friends.GetFriends() {
		friendS_relationId = append(friendS_relationId, friend.RelationId)
	}

	// Test create friend bill
	_, err = client.CreateFriendBill(context.TODO(), &monify.CreateFriendBillRequest{
		RelationId:  friendS_relationId[0],
		Amount:      300,
		Title:       "test1",
		Description: "test1",
	})
	_, err = client.CreateFriendBill(context.TODO(), &monify.CreateFriendBillRequest{
		RelationId:  friendS_relationId[0],
		Amount:      500,
		Title:       "test2",
		Description: "test2",
	})
	assert.NoError(t, err)
	_, err = client.CreateFriendBill(context.TODO(), &monify.CreateFriendBillRequest{
		RelationId:  friendS_relationId[0],
		Amount:      0,
		Title:       "test2",
		Description: "test2",
	})
	assert.Error(t, err) // Amount must not be zero

	// Test list friend bill
	res1, err := client.ListFriendBill(context.TODO(), &monify.ListFriendBillRequest{RelationId: friendS_relationId[0]})
	assert.NoError(t, err)
	assert.Equal(t, len(res1.GetFriendBills()), 2)
	assert.Equal(t, 300.0, res1.FriendBills[0].Amount)
	assert.Equal(t, "test1", res1.FriendBills[0].Title)
	assert.Equal(t, "test1", res1.FriendBills[0].Description)
	assert.Equal(t, 500.0, res1.FriendBills[1].Amount)
	assert.Equal(t, "test2", res1.FriendBills[1].Title)
	assert.Equal(t, "test2", res1.FriendBills[1].Description)

	// Test delete friend bill
	_, err = client.DeleteFriendBill(context.TODO(), &monify.DeleteFriendBillRequest{FriendBillId: res1.FriendBills[0].FriendBillId})
	assert.NoError(t, err)
	res2, err := client.ListFriendBill(context.TODO(), &monify.ListFriendBillRequest{RelationId: friendS_relationId[0]})
	assert.Equal(t, len(res2.GetFriendBills()), 1)
	assert.Equal(t, 500.0, res2.FriendBills[0].Amount)
	assert.Equal(t, "test2", res2.FriendBills[0].Title)
	assert.Equal(t, "test2", res2.FriendBills[0].Description)

	// Test modify friend bill
	_, err = client.ModifyFriendBill(context.TODO(), &monify.ModifyFriendBillRequest{
		FriendBillId: res2.FriendBills[0].FriendBillId,
		RelationId:   friendS_relationId[0],
		Amount:       1000,
		Title:        "test3",
		Description:  "test3",
	})
	assert.NoError(t, err)
	res3, err := client.ListFriendBill(context.TODO(), &monify.ListFriendBillRequest{RelationId: friendS_relationId[0]})
	assert.Equal(t, 1000.0, res3.FriendBills[0].Amount)
	assert.Equal(t, "test3", res3.FriendBills[0].Title)
	assert.Equal(t, "test3", res3.FriendBills[0].Description)
}
