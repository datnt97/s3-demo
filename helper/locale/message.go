package locale

var (
	SuccessCode   = "200"
	FailedCode    = "500"
	FailedMsg     = NewMessage(FailedCode)
	SuccessMsg    = NewWithMessage(SuccessCode, "Success")
	NoDataMsg     = NewWithMessage("204", "No Data")
	BadReqMsg     = NewWithMessage("400", "Bad Request")
	NoAuthMsg     = NewWithMessage("401", "Unauthorized")
	NoPerMsg      = NewMessage("missing_permission")
	DuplicateData = NewWithMessage("duplicate_data", "The data already exists")
)

func NewMessage(key string) *Message {
	return &Message{Key: key}
}

func NewWithMessage(key string, msg string) *Message {
	return &Message{Key: key, Message: msg}
}

type Message struct {
	Key     string `json:"key"`
	Message string `json:"msg"`
}
