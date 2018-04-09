package main

import (
	"fmt"
	"sync/atomic"
)

// Room 房间
type Room interface {
	LanguageManager
	// 进入房间说话
	IntoSay(Personer)
	// 显示说话次数
	ShowSayNum()
}

type room struct {
	baseLanguageManag
	name string
	cnt  int64
}

func (r *room) IntoSay(p Personer) {
	curlang := r.Language()
	if curlang == nil {
		curlang = p.Language()
	}
	if curlang == nil {
		fmt.Printf("%s进入房间[%s]没有说话\n", color("31", p.Name()), color("32", r.name))
		return
	}
	fmt.Printf("%s进入房间[%s]说：%s\n", color("31", p.Name()), color("32", r.name), color("33", curlang.HelloWorld()))
	_ = atomic.AddInt64(&r.cnt, 1)
}
func (r *room) ShowSayNum() {
	fmt.Printf("房间[%s]:已经有人说过%d次\n", color("32", r.name), r.cnt)
}

// CreRoom 创建房间
func CreRoom(name string) Room {
	return &room{
		name: name,
		cnt:  0,
	}
}
func color(color string, txt interface{}) string {
	return fmt.Sprintf("\033[1m\033[%sm%v\033[0m\033[0m", color, txt)
}
