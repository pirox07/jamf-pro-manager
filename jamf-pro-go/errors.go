package jamf_pro_go

const (
	//XXXXXXXXXX      = "[error message]"
)

type UnauthorizedError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type Error struct {
	StatusCode              int
	RawError                string
	IsAuthorizationRequired bool
}

func (e *Error) Error() string {
	return e.RawError
}
