package main

import (
	"fmt"
	"os"
	"testing"
)

// 测试文件名称： file_name_test.go
// 测试命令：go test clac.go clac_test.go -v
func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("test fail")
		return
	}
	t.Logf("test pass")
}

func TestAddBatch(t *testing.T) {
	// 匿名构造体
	var tests = []struct {
		a, b int
		want int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{-1, -2, -3},
		{2, 2, 4},
	}
	// 批量测试
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := Add(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func setUp() {
	fmt.Println("setUp running....")
}

func tearDown() {
	fmt.Println("tearDown running....")
}

// 测试的主函数,注意这里是反射的必须是公开方法
func TestMain(m *testing.M) {
	fmt.Println("testMain start....")
	setUp()
	code := m.Run()
	tearDown()
	fmt.Println("testMain end....")
	os.Exit(code)
}
