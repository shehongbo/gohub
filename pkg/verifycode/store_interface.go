package verifycode

type Store interface {
	// 保存验证码
	Set(id string, value string) bool

	// 获取验证码
	Get(id string, clear bool) string

	//验证验证码
	Verify(id, answer string, cleat bool) bool
}
