package api

var invalid = new(invalidAPI)

type invalidAPI struct{}

func (m invalidAPI) Auth() bool {
	return true
}

func (m invalidAPI) Call() interface{} {
	Throw(APIErrorCode, "")
	return nil
}

func (m invalidAPI) Valid(ctx interface{}) bool {
	return true
}
