package demo2

import "testing"

func say(s string) string {
	return "hello " + s
}

func TestHelp1(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper() //t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	assert2 := func(t *testing.T, got, want string) {
		// 这里没有调用 t.Helper()，报错的时候行号是
		// TestHelp1/test2: demo_test.go:21: got 'hello lili' want 'hello li'
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("test1", func(t *testing.T) {
		got := say("li")
		want := "hello zhang"
		assert(t, got, want)
	})

	t.Run("test2", func(t *testing.T) {
		got := say("lili")
		want := "hello li"
		assert2(t, got, want)
	})
}
