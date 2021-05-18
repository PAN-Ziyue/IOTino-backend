package e

// MsgFlags
// define the status message
var MsgFlags = map[int]string{
    SUCCESS: "ok",
    ERROR:   "fail",
}

// GetMsg
// interface for getting the message
func GetMsg(code int) string {
    msg, ok := MsgFlags[code]

    if ok {
        return msg
    }
    return MsgFlags[ERROR]
}
