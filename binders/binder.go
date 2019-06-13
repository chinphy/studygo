package binders

// User bind struct
type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Age      uint8  `form:"age"`
	Status   uint8  `form:"status"`
}
