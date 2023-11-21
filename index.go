package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	// 1_1234と、アンダースコアを記載できる。

	// var flag bool // 値が代入されないとfalseになる
	// var isAwesome = true

	// 型変換は全て明示的に行われる
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)
	fmt.Println(z, d)

	var b int = 10
	a := 10
	fmt.Println(b)
	fmt.Println(a)

	// 配列
	var intArray [3]int
	fmt.Println(intArray)

	var riteral = [3]int{10, 20, 30}
	fmt.Println(riteral)

	var sumArray = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(sumArray)

	var rite = [...]int{10, 20, 30}
	fmt.Println(rite)

	// 多次元配列
	var big [2][3]int
	fmt.Println(big)

	// 配列に代入
	rite[0] = 40


	// スライス
	var slice = []int{10, 20, 30}
	fmt.Println(slice)

	var sliceX [][]int
	fmt.Println(sliceX)

	// スライスのゼロ値、すなわちnulが初期値になる
	var sliceNil []int
	fmt.Println(sliceNil) // Goのnilは、いくつかの型において「値がない」ことを示す識別子となる。

	// 配列の代入
	var ar []int
	ar = append(ar, 10)
	fmt.Println(ar)
	ar = append(ar, 5, 6, 7)
	fmt.Println(ar)

	apRite := []int{20, 30, 40}
	fmt.Println(apRite)
	ar = append(ar, apRite...)
	print(apRite)


	// 各スライスには、キャパシティを持っている。
	// cap関数で確認することができる。p45

	// makeを使えば、スライスのキャパシティを宣言することができる。
	ma := make([]int, 5) // 長さ5、キャパシティ5のintのスライスが生成される。
	fmt.Println(ma)

	ng := make([]int , 5)
	ng = append(ng, 10) // lenは6、キャパシティは10になってしまう
	fmt.Println(ng)

	// この考え方よく使いそう
	good := make([]int, 0, 10) // 長さ0でキャパシティ10のスライス
	good = append(good, 3, 4, 5) // 長さは3に変わり、キャパシティは10のまま
	fmt.Println(good)

	// どうすればスライスを大きくする回数を減らせるかを考えよう
	// スライスが全く大きくならない可能性があるのならば、次のようにvarとnilスライスを使うのがいい。
	var data []int
	fmt.Println(data == nil) // true

	// 実行してみればスライスのサイズがどのくらいになるか予想ができるが、プログラムを書いている最中にはわからない場合にはmakeを使う


}