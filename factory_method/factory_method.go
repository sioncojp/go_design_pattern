package factory_method

import (
	"fmt"
)

type User struct {
	Name string
	Buy  string
}

type kit interface {
	build()
}

type Gundam struct {
	kit
}

func (s *Gundam) Create() {
	s.kit.build()
}

type WGundam struct{}

func (s *WGundam) build() {
	println("ビームライフル組み立てる")
}

func NewWGundam() *Gundam {
	return &Gundam{new(WGundam)}
}

type GodGundam struct{}

func (s *GodGundam) build() {
	println("ゴッドフィンガーを組み立てる")
}

func NewGodGundam() *Gundam {
	return &Gundam{new(GodGundam)}
}

// factoryクラス
func NewGundam(s string) *Gundam {
	switch s {
	case "GodGundam":
		return NewGodGundam()
	case "WGundam":
		return NewWGundam()
	default:
		fmt.Println("作りたいガンダム選んでね")
	}
	return nil
}
