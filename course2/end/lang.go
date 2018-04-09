package main

// LanguageManager 语言管理器
type LanguageManager interface {
	Language() Language
	SetLanguage(Language)
}

// Language 语言
type Language interface {
	// 返回hello world
	HelloWorld() string
}

type enLanguage struct {
}

func (enLanguage) HelloWorld() string {
	return " Hello world!"
}

type cnLanguage struct {
}

func (cnLanguage) HelloWorld() string {
	return "你好，世界!"
}

type baseLanguageManag struct {
	lang Language
}

func (l baseLanguageManag) Language() Language {
	return l.lang
}
func (l *baseLanguageManag) SetLanguage(lang Language) {
	l.lang = lang
}
