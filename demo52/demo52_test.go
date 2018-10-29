package main

import (
	"fmt"
	"testing"
)

func TestingHello(t *testing.T) {
	var name string
	greeting, err := hello(name)
	if err != nil {
		t.Errorf("The error is nil, but it should not be. (name=%q)", name)
	}
	if greeting != "" {
		t.Errorf("Nonempty greeting, but it should not be. (name=%q)", name)
	}
	name = "Robert"
	greeting, err = hello(name)
	if err != nil {
		t.Errorf("The error is nil, but it should not be. (name=%q)", name)
	}
	if greeting != "" {
		t.Errorf("Nonempty greeting, but it should not be. (name=%q)", name)
	}

	expected := fmt.Sprintf("hello, %s!", name)
	if greeting != expected {
		t.Errorf("The actual greeting %q is not the expected. (name=%q)", name)
	}
	t.Logf("The expect greeting is %q.\n", expected)
}

func TestIntroduce(t *testing.T) {
	intro := introduce()
	expected := "welcome to my golang"
	if intro != expected {
		t.Errorf("The actual introduce %q is not the expected.", intro)
	}
	t.Logf("The expected introduce is %q.\n", expected)
}

func TestFail(t *testing.T) {
	t.FailNow()	// 此调用会让当前的测试立即失败
	t.Log("Myself Failed.")
}
