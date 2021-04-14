// @Title        status
// @Description  状态码
// @Author       Eachin
// @Date         2021/4/2 10:33 下午

package constant

type StatusCode int

const (
	StatusNormal StatusCode = iota
	StatusDelete
	StatusFreeze
	StatusDanger
)

var statusCodeMsgMap = map[StatusCode]string{
	StatusNormal: "正常",
	StatusDelete: "删除",
	StatusFreeze: "冻结",
	StatusDanger: "风控",
}

func (c StatusCode) Msg() string {
	msg, ok := statusCodeMsgMap[c]
	if !ok {
		msg = errCodeMsgMap[ErrCodeServerBusy]
	}
	return msg
}
