package response

const (
	SUCCESS                        = 200
	ERROR                          = 500
	INVALID_PAYLOAD                = 400
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PAYLOAD:                "invalid payload",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "check token fail",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "check token timeout",
	ERROR_AUTH_TOKEN:               "auth token",
	ERROR_AUTH:                     "auth",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
