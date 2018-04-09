package main

// Personer 人
type Personer interface {
	LanguageManager
	Name() string
}

// CrePerson 创建一个人
func CrePerson(name string) Personer {
	return &person{
		name: name,
	}
}

type person struct {
	baseLanguageManag
	name string
}

func (p person) Name() string {
	return p.name
}
