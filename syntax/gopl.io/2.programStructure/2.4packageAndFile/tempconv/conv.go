package tempconv

// 转换函数

// 摄氏度转华氏度
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// 开尔文转摄氏度
func KToC(k Kelvin) Celsius {
	return Celsius(k) + AbsoluteZeroC
}

// 摄氏度转开尔文
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}
