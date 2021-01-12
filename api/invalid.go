package api

type invalidAPI struct{}

func (m invalidAPI) Call() (interface{}, error) {
	return nil, NewError(APIErrorCode, "")
}

func (m invalidAPI) GetScope() int {
	return 0
}
