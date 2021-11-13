package contextkey

type Value string

func (m Value) String() string {
	return string(m)
}

const (
	HttpBody = "http.body"
)
