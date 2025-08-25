package main

import "testing"

func TestShowWithTable(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"面相加薪学习-从0到GO语言微服务架构师",
			"面相加薪学习-从0到GO语言微服务架构师",
		},
		{
			"面相加薪学习面相加薪学习-从0到GO语言微服务架构师",
			"面相加薪学习-从0到GO语言微服务架构师",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Show(); got != tt.want {
				t.Errorf("Show() = %v, want %v", got, tt.want)
			} else {
				t.Logf("Show() = %v", got)
			}
		})
	}
}

func TestShow(t *testing.T) {
	result := Show()
	want := "面相加薪学习-从0到GO语言微服务架构师"
	if result != want {
		t.Errorf("got:%v, want:%v", result, want)
	} else {
		t.Logf("got:%v, want:%v", result, want)
	}
}
