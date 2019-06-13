package constant

const (
	// UserVerifyCheckAll 全部检查
	UserVerifyCheckAll byte = 255
	// UserVerifyCheckPwd 检查密码
	UserVerifyCheckPwd byte = 1
	// UserVerifyCheckStat 检查状态
	UserVerifyCheckStat byte = UserVerifyCheckPwd << 1
)
