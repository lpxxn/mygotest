package dig1

import (
	"bytes"
	"reflect"
	"runtime"
	"testing"

	"go.uber.org/dig"
)

func TestDigErr1(t *testing.T) {
	d := dig.New()
	f1 := func() int {
		return 1
	}

	f2 := func() int {
		return 2
	}
	err := d.Provide(f1)
	if err != nil {
		t.Fatal(err)
	}
	err = d.Provide(f2)
	if err != nil {
		t.Error(err)
	}
	/*
		cannot provide function "github.com/mygotest/fx_demo/dig1".TestDigErr1.func2 (/Users/li/go/src/github.com/mygotest/fx_demo/dig1/dig_err_test.go:15):
		cannot provide int from [0]: already provided by "github.com/mygotest/fx_demo/dig1".TestDigErr1.func1 (/Users/li/go/src/github.com/mygotest/fx_demo/dig1/dig_err_test.go:11)
	*/
	fl := func(constructor interface{}) {
		fptr := reflect.ValueOf(constructor).Pointer()
		f := runtime.FuncForPC(fptr)
		fileName, lineNum := f.FileLine(fptr)
		t.Logf("fileName: %s, line: %d funcName: %s\n", fileName, lineNum, f.Name())
	}
	fl(f1)
	fl(f2)
}

func TestDitErr2(t *testing.T) {
	d := dig.New()
	type out1 struct {
		dig.Out
		Value []int `group:"val,flatten"`
	}

	provide := func(i []int) {
		if err := d.Provide(func() out1 {
			return out1{Value: i}
		}); err != nil {
			t.Fatal(err)
		}
	}

	provide([]int{1, 2})
	provide([]int{3, 4})

	type in struct {
		dig.In

		Values []int `group:"val"`
	}
	type out2 struct {
		dig.Out

		Value []int `group:"val,flatten"`
	}
	// valcy
	if err := d.Provide(func(i in) out2 {
		return out2{Value: append(i.Values, 7, 8)}
	}); err != nil {
		//t.Log(err.Error())
		t.Fatal(err)
	}
	/*
		cannot provide function "github.com/mygotest/fx_demo/dig1".TestDitErr2.func2 (/Users/li/go/src/github.com/mygotest/fx_demo/dig1/dig_err_test.go:60):
		this function introduces a cycle: int[group="val"] provided by "github.com/mygotest/fx_demo/dig1".TestDitErr2.func2 (/Users/li/go/src/github.com/mygotest/fx_demo/dig1/dig_err_test.go:60)
		depends on int[group="val"] provided by "github.com/mygotest/fx_demo/dig1".TestDitErr2.func2 (/Users/li/go/src/github.com/mygotest/fx_demo/dig1/dig_err_test.go:60)
	*/

	err := d.Invoke(func(i in) {
		t.Log(i.Values)
	})
	if err != nil {
		t.Fatal(err)

	}
}

func TestDigErr3(t *testing.T) {
	d := dig.New()
	f1 := func() int {
		return 1
	}

	err := d.Provide(f1)
	if err != nil {
		t.Fatal(err)
	}
	// 有返回参数不报错
	if err := d.Invoke(func(i int) (string, error) {
		t.Log(i)
		return "", nil
	}); err != nil {
		t.Fatal(err)
	}

	//  missing type: string
	// Invoke 返回的 string 参数，不被记录到provider里
	// 这里也返回string 不报错
	if err := d.Invoke(func(s string) (string, error) {
		t.Log(s)
		return "", nil
	}); err != nil {
		t.Fatal(err)
	}
}

func TestFuncLine(t *testing.T) {
	// go test -v -gcflags all="-N -l"  -run TestFuncLine 禁止内联
	f1 := func() int {
		return 1
	}

	f2 := func() int {
		return 2
	}
	fl := func(constructor interface{}) {
		fptr := reflect.ValueOf(constructor).Pointer()
		f := runtime.FuncForPC(fptr)
		fileName, lineNum := f.FileLine(fptr)
		t.Logf("fileName: %s, line: %d funcName: %s\n", fileName, lineNum, f.Name())
	}
	fl(f1)
	fl(f2)
}

func TestDigErr4(t *testing.T) {
	d := dig.New()
	f1 := func() int {
		return 1
	}
	val1 := 0
	f2 := func(i int) (string, error) {
		val1 += i
		return "ok", nil
	}
	// 顺序无关
	if err := d.Provide(f2); err != nil {
		t.Fatal(err)
	}

	if err := d.Provide(f1); err != nil {
		t.Fatal(err)
	}
	type Resp struct {
		dig.Out
		Name string
	}
	type Param struct {
		dig.In
		Age int
	}
	// invoke 里没有用float32的参数，这个方法也不执行
	// 其他被调用到的 provide没有用float32，所以这个也不执行
	f3 := func(s string, p Param) (int, error) {
		t.Log("func3")
		t.Log(s)
		//return Resp{Name: "aaa"}, nil
		return 0, nil
	}
	if err := d.Provide(f3); err != nil {
		t.Fatal(err)
	}

	t.Log(d.String())

	b := &bytes.Buffer{}
	if err := dig.Visualize(d, b); err != nil {
		t.Fatal(err)
	}
	t.Log(b.String())
}
