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

	
	// 配列からスライスへの変換
	sliceArX := [...]int{5, 6, 7, 8}
	sliceArY := sliceArX[:2]
	sliceArZ := sliceArX[2:]
	sliceArD := sliceArX[:]
	sliceArX[0] = 10
	
	fmt.Println(sliceArX)
	fmt.Println(sliceArY)
	fmt.Println(sliceArZ)
	fmt.Println(sliceArD)

	
	// メモリを共有しない独立したスライスの作成
	memoX := []int{1,2,3,4}
	memoY := make([]int, 4)
	num := copy(memoY, memoX)
	fmt.Println(memoY, num)

	x := []int{1,2,3,4,5}
	num := copy(x[:3], x[1:])
	

	// マップリテラル
	totalWins := map[string]int{}
	// 空のマップリテラルで初期化すると、そのマップに対して読み書きの両方が可能になる
	fmt.Println(totalWins == nul)
	fmt.Println(totalWins["abc"]) // 0
	totalWins["abc"] = 3;
	fmt.Println(totalWins["abc"]);
	
	teams := map[string][]string {
		"ライターズ": []string{"夏目", "森", "国木田"},
		"ナイツ": []string{"武田", "徳川", "明智"},
		"ミュージシャンズ": []string{"ラベル", "ベートーベン", "リスト"},
	}
	

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok) // 5 true
	v, ok = m["world"]
	fmt.Println(v, ok) // 0 true
	v, ok = m["goodbye"]
	fmt.Println(v, ok) // 0 false
	
	var maps := map[int]bool{}
	var vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, val := range vals {
		maps[val] = true
	}

	// mapは「APIを定義しない」。特定のキーだけを含むような「限定されたマップ」の定義ができないということ

	// struct
	type person struct {
		name string
		age int
		pet string
	}
	var fred person
	bob := person{} // 構造体リテラル。全フィールドがゼロ値で初期化される
	julia := person{
		"ジュリア"
		40,
		"cat"
	}
	beth := person{
		age: 30,
		name: "ベス",
	}
}