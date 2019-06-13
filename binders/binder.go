package binders

// User bind struct
type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Age      int    `form:"age"`
}
