package modelErr

// model 层错误封装，只允许在model中使用
// rpc logic 中捕获到 model 中返回的错误，已在model中封装，直接抛出即可

var msg map[int]string

func init() {
	msg = make(map[int]string)
	msg[ARTICLE_NOT_FOUND] = "article not found"
}
