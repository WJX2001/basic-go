package tempconv

import "fmt"

type Celsius float64    // 摄氏度
type Fahrenheit float64 // 华氏度
type Kelvin float64     // 开尔文（绝对温度）

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
	AbsoluteZeroK Kelvin  = 0       // Kelvin的绝对零度
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
