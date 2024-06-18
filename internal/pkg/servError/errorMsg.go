package servError

var msg = map[int]string{}

func Msg(code int) string {
	return msg[code]
}

func init() {
	msg[Common] = "common error"
	msg[UNKNOWN] = "unknown error"

	msg[InvalidParam] = "invalid param"
	msg[InvalidAction] = "invalid action"

	msg[ArticleNotFound] = "article not found"
}
