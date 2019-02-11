package main

<<<<<<< HEAD
import "fmt"
=======
// Person 请假人
type Person struct {
	name   string
	dayoff int
	reason string
	leader Manager
}

func (c Person) ReuqestD() {

}

// ManagerInterface 管理人员接口
type ManagerInterface interface {
	HandleRequest()
}

// Manager 管理人员基类
type Manager struct {
	Name        string
	Title       string
	NextHandler ManagerInterface
}

func (c Manager) HandleRequest() string {
	return "hello"
}

// Supervisor 直接上级
type Supervisor struct {
	Manager
}

func (c Supervisor) HandleRequest() {
	return
}

// DepartmentManager 部门总监
type DepartmentManager struct {
	Manager
}

func (c DepartmentManager) HandleRequest() {
	return
}

// CEO ceo
type CEO struct {
	Manager
}

func (c CEO) HandleRequest() {
	return
}

// Administrator 行政
type Administrator struct {
	Manager
}

func (c Administrator) HandleRequest() {
	return
}
>>>>>>> 8cbb495e8422cd47aaa577f325d29064b5255ce8

func main() {
	fmt.Println("hello")
}
