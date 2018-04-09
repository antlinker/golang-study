package main

func main() {
	// 创建语言
	en, cn := enLanguage{}, cnLanguage{}
	r603, r601, rzh := CreRoom("603"), CreRoom("601"), CreRoom("只说中文")
	p, p1 := CrePerson("我"), CrePerson("你")
	r603.IntoSay(p)
	p.SetLanguage(en)
	r601.IntoSay(p)
	r603.IntoSay(p1)
	p1.SetLanguage(cn)
	r601.IntoSay(p1)
	r601.ShowSayNum()
	r603.ShowSayNum()

	rzh.SetLanguage(cn)
	rzh.IntoSay(p)
	rzh.IntoSay(p1)
	rzh.ShowSayNum()
}
