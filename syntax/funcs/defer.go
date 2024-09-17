package main

// 是一个后进先出的特性
func Defer() {
	// 允许在返回的前一刻 执行一段逻辑，起到延迟调用的逻辑

	defer func() {
		println("第一个 defer")
	}()

	defer func() {
		println("第二个 defer")
	}()
}

func DeferClosure() {
	i := 0
	defer func() {
		println(i) // 1
	}()
	i = 1
}

func DeferClosureV1() {
	i := 0
	defer func(val int) {
		println(val) // 0
	}(i)
	i = 1
}
