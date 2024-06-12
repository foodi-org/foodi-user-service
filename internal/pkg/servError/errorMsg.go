package servError

var msg = map[int]string{}

func Msg(code int) string {
	return msg[code]
}

func init() {
	msg[COMMON] = "common error"
	msg[UNKONW] = "unkonw error"

	msg[InvalidParam] = "invalid param"
	msg[InvalidAction] = "invalid action"

	msg[ARTICLE_NOT_FOUND] = "article not found"
}
