package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	//b = int64(MyInt)
	var c MyInt
	t.Log(a, b, c)
}

// go 不支持隐式类型转换，别名也不行

func TestPointt(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

// go 不支持指针运算

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
}

// go 会初始化所有变量为0值
