package value

// IUserAPICaller is 用户api调用者
type IUserAPICaller interface {
	CommitUpdate(route string) error
	Update(userID string, rewards ...Reward)
}
