package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

var justString string

func createHugeString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz"

	var sb strings.Builder

	for i := 0; i < length; i++ {
		randVal := rand.Intn(len(charset))
		sb.WriteByte(charset[randVal])
	}

	return sb.String()
}

func someFunc() {
	v := createHugeString(1 << 20)

	// Сохраняем слайс глобально, сохранив отрезок из первых 100 символов
	// но остальная часть (1024-100) никуда не делась, и так же существует в памяти

	// Вывод MemStats alloc в данном случае
	// 148536 - изначальная
	// 1278296 - итоговая
	//justString = v[:100]

	// Вывод MemStats alloc в данном случае
	// 148536 - изначальная
	// 164200 - итоговая
	var sb strings.Builder

	for i := 0; i < 100; i++ {
		sb.WriteByte(v[i])
	}

	justString = sb.String()
}

func main() {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc)

	someFunc()

	// После вызова GC, в варианте с justString = v[:100] в памяти остается весь слайс
	// А если мы создаем свою строку, то большая строка очищается, так как не остается ссылающих элементов
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Println(m.Alloc)

	fmt.Println(justString)
}
