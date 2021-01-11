package api

var invalid = new(invalidAPI)

type invalidAPI struct{}

func (m invalidAPI) Auth() bool {
	return true
}

func (m invalidAPI) Call() (interface{}, error) {
	Throw(APIErrorCode, "")
	return nil, nil
}

func (m invalidAPI) SetRequest(_ interface{}) {

}

func (m invalidAPI) Valid() bool {
	return true
}
