package main

import "fmt"

func main() {
	start()
}

func start() {
	from := inputValue("摂氏[°C] -> 華氏[°F]", "°C")
	to := from*1.8 + 32
	fmt.Printf("結果: %.2f[°F]\n", to)
}

// TODO: 値を入力する関数を定義する
func inputValue(name, fromUnit string) float64 {
	fmt.Println(name)
	var v float64
	fmt.Print("変換する値[°C]>")
	fmt.Scanln(&v)
	return v
}
