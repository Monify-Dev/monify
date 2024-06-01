package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	monify "monify/protobuf/gen/go"
	"testing"
)

func TestCreateAndGetGroupBill(t *testing.T) {
	client := GetTestClient(t)
	_ = client.CreateTestUser()

	group, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test"})
	assert.NoError(t, err)
	inviteCode, err := client.GenerateInviteCode(context.Background(), &monify.GenerateInviteCodeRequest{GroupId: group.GroupId})
	assert.NoError(t, err)
	_ = client.CreateTestUser()
	memberId2, err := client.JoinGroup(context.Background(), &monify.JoinGroupRequest{InviteCode: inviteCode.InviteCode})
	assert.NoError(t, err)

	response1, err := client.CreateGroupBill(context.TODO(), &monify.CreateGroupBillRequest{
		GroupId:     group.GroupId,
		TotalMoney:  250,
		Title:       "test",
		Description: "test",
		SplitPeople: []*monify.SplitPerson{
			{
				MemberId: memberId2.MemberId,
				Amount:   100,
			},
			{
				MemberId: group.MemberId,
				Amount:   150,
			},
		},
		PrepaidPeople: []*monify.PrepaidPerson{
			{
				MemberId: group.MemberId,
				Amount:   250,
			},
		},
	})

	assert.NoError(t, err)

	response2, err := client.GetGroupBills(context.TODO(), &monify.GetGroupBillsRequest{GroupId: group.GroupId})
	assert.NoError(t, err)
	assert.Len(t, response2.GroupBills[0].SplitPeople, 2)
	assert.Len(t, response2.GroupBills[0].PrepaidPeople, 1)
	assert.Equal(t, response1.BillId, response2.GroupBills[0].BillId)
	assert.Equal(t, group.MemberId, response2.GroupBills[0].PrepaidPeople[0].MemberId)
	assert.Equal(t, 250.0, response2.GroupBills[0].PrepaidPeople[0].Amount)
}