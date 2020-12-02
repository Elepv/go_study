package observer

import "fmt"

//Subject 定义一个目标的结构体对象
type Subject struct {
	observers []Observer
	context   string
}

//NewSubject 新建Subject的方法
func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

//Attach 订阅目标的方法
func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

//通知订阅者的方法
func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

//UpdateContext 更新内容的方法
func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

//Observer 建立一个观察者对象的接口
//该接口规定了一个更新方法来更新目标
type Observer interface {
	Update(*Subject)
}

//Reader 建立一个订阅者的对象
type Reader struct {
	name string
}

//NewReader 建立一个实例化结构体的方法
func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

//Update 具体实现订阅者更新信息的方法
func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}
