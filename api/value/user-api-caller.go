package value

import (
	"fmt"

	"github.com/ahl5esoft/lite-go/api"
)

const (
	commitUpdateRouteFormat = "%s/server/update-users"
)

type commitUpdateRequest struct {
	Messages []updateRequest
}

type updateRequest struct {
	Rewards []Reward
	Route   string
	UserID  string
}

type userAPICaller struct {
	apiCaller         api.ICaller
	commitUpdateRoute string
	updateRequests    []updateRequest
}

func (m *userAPICaller) CommitUpdate(route string) error {
	defer func() {
		m.updateRequests = make([]updateRequest, 0)
	}()

	for i := 0; i < len(m.updateRequests); i++ {
		m.updateRequests[i].Route = route
	}
	return m.apiCaller.VoidCall(m.commitUpdateRoute, commitUpdateRequest{
		Messages: m.updateRequests,
	})
}

func (m *userAPICaller) Update(userID string, rewards ...Reward) {
	m.updateRequests = append(m.updateRequests, updateRequest{
		Rewards: rewards,
		UserID:  userID,
	})
}

// NewUserAPICall is IUserAPICaller实例
func NewUserAPICall(apiCaller api.ICaller, app string) IUserAPICaller {
	return &userAPICaller{
		apiCaller:         apiCaller,
		commitUpdateRoute: fmt.Sprintf(commitUpdateRouteFormat, app),
		updateRequests:    make([]updateRequest, 0),
	}
}
