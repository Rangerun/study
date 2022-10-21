package test

type TestStu struct {
	TestService
}

func NewTestService (...interface{}) (interface{}, error) {
	return &TestStu{}, nil
}

func (TsetStu *TestStu)GetName() string {
	stu := Stu{name: "angel"}
	return stu.name
} 
