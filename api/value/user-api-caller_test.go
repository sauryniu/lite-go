package value

import (
	"testing"

	"github.com/ahl5esoft/lite-go/api"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_userAPICaller_CommitUpdate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockAPICaller := api.NewMockICaller(ctrl)

		self := userAPICaller{
			apiCaller:         mockAPICaller,
			commitUpdateRoute: "caller-route",
			updateRequests: []updateRequest{
				{
					Rewards: []Reward{
						{
							Count:       1,
							TargetIndex: 2,
							TargetType:  3,
							ValueType:   4,
						},
					},
					UserID: "one",
				},
			},
		}

		commitUpdateRoute := "route"
		mockAPICaller.EXPECT().VoidCall(self.commitUpdateRoute, commitUpdateRequest{
			Messages: []updateRequest{
				{
					Rewards: self.updateRequests[0].Rewards,
					Route:   commitUpdateRoute,
					UserID:  self.updateRequests[0].UserID,
				},
			},
		}).Return(nil)

		err := self.CommitUpdate(commitUpdateRoute)
		assert.NoError(t, err)
		assert.Len(t, self.updateRequests, 0)
	})
}

func Test_userAPICaller_Update(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := userAPICaller{
			updateRequests: make([]updateRequest, 0),
		}
		userID := "one"
		reward := Reward{
			Count:       1,
			TargetIndex: 2,
			TargetType:  3,
			ValueType:   4,
		}
		self.Update(userID, reward)
		assert.EqualValues(t, self.updateRequests, []updateRequest{
			{
				Rewards: []Reward{reward},
				UserID:  userID,
			},
		})
	})
}
