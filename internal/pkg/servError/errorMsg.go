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

	msg[MissPhoneNumber] = "miss phone number"

	msg[MissPassword] = "miss password"

	msg[LoginFailError] = "login fail"

	msg[LoginCodeError] = "login code error"

	msg[LoginPasswordError] = "login password error"

	msg[PhoneNumberNotFound] = "phone number not found"

	msg[AccountNotBindPassword] = "account not bind password"

	msg[AccountPasswordNotMatch] = "account password not match"

	msg[InvalidUserCoup] = "invalid user coup"

	msg[InvalidRegisterCoup] = "invalid register coup"

	msg[ArticleNotFound] = "article not found"

	msg[ArticleCreateFail] = "article create fail"

	msg[ArticleSaveFail] = "article save fail"

	msg[CommentNotFound] = "comment not found"

}
