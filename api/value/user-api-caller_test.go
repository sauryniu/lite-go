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
			rewards: []userReward{
				{
					Reward: Reward{
						Count:       1,
						TargetIndex: 2,
						TargetType:  3,
						ValueType:   4,
					},
					UserID: "one",
				},
			},
		}

		commitUpdateRoute := "route"
		mockAPICaller.EXPECT().VoidCall(self.commitUpdateRoute, commitUpdateRequest{
			Rewards: []userReward{
				{
					Reward: Reward{
						Count:       1,
						Route:       commitUpdateRoute,
						TargetIndex: 2,
						TargetType:  3,
						ValueType:   4,
					},
					UserID: self.rewards[0].UserID,
				},
			},
		}).Return(nil)

		err := self.CommitUpdate(commitUpdateRoute)
		assert.NoError(t, err)
		assert.Len(t, self.rewards, 0)
	})
}

func Test_userAPICaller_Update(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		self := userAPICaller{
			rewards: make([]userReward, 0),
		}
		userID := "one"
		reward := Reward{
			Count:       1,
			TargetIndex: 2,
			TargetType:  3,
			ValueType:   4,
		}
		self.Update(userID, reward)
		assert.EqualValues(t, self.rewards, []userReward{
			{
				Reward: reward,
				UserID: userID,
			},
		})
	})
}
