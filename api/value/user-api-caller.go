package value

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/api"
)

const (
	commitUpdateRouteFormat = "%s/server/update-users"
)

type commitUpdateRequest struct {
	Rewards []userReward
}

type userReward struct {
	Reward

	UserID string
}

type userAPICaller struct {
	apiCaller         api.ICaller
	commitUpdateRoute string
	rewards           []userReward
}

func (m *userAPICaller) CommitUpdate(route string) error {
	defer func() {
		m.rewards = make([]userReward, 0)
	}()

	for i := 0; i < len(m.rewards); i++ {
		m.rewards[i].Route = route
	}
	return m.apiCaller.VoidCall(m.commitUpdateRoute, commitUpdateRequest{
		Rewards: m.rewards,
	})
}

func (m *userAPICaller) Update(userID string, rewards ...Reward) {
	for _, r := range rewards {
		m.rewards = append(m.rewards, userReward{
			Reward: r,
			UserID: userID,
		})
	}
}

// NewUserAPICall is IUserAPICaller实例
func NewUserAPICall(apiCaller api.ICaller, app string) IUserAPICaller {
	return &userAPICaller{
		apiCaller:         apiCaller,
		commitUpdateRoute: fmt.Sprintf(commitUpdateRouteFormat, app),
		rewards:           make([]userReward, 0),
	}
}
