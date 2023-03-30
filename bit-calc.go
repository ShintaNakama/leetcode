package main

import "fmt"

// ANDビット演算
// 下記の場合、10101010と11110000の比較し、それぞの数字の各ビットを比較して両方1なら1それ以外は0にする
// どちらかというと図で考える。a,bの左端は両方1なので演算結果のcの左端も1となり、右端はa,bが0なのでcも0となる
// a: 10101010
// b: 11110000
// c: 10100000
func andCalc() {
	var a uint8 = 0b10101010 // 170
	var b uint8 = 0b11110000 // 240

	c := a & b // AND演算

	fmt.Println(a & b)
	fmt.Println(c)
	fmt.Printf("a & b = %08b\n", c) // 出力: a & b = 10100000
}

// ORビット演算
// 下記の場合、10101010と11110000の比較し、それぞの数字の各ビットを比較して片方または両方1なら1それ以外は0にする
// どちらかというと図で考える。a,bの左端は両方1なので演算結果のcの左端も1となり、右端はa,bが0なのでcも0となる
// a: 10101010
// b: 11110000
// c: 11111010
func orCalc() {
	var a uint8 = 0b10101010 // 170
	var b uint8 = 0b11110000 // 240

	c := a | b // OR演算

	fmt.Println(a | b)
	fmt.Println(c)
	fmt.Printf("a | b = %08b\n", c) // 出力: a | b = 11111010
}

// XORビット演算
// 下記の場合、10101010と11110000の比較し、それぞの数字の各ビットを比較して片方のみ1なら1それ以外は0にする
// どちらかというと図で考える。a,bの左端は両方1なので演算結果のcの左端は0となり、右端はa,bが0なのでcも0となる
// a: 10101010
// b: 11110000
// c: 01011010
func xorCalc() {
	var a uint8 = 0b10101010 // 170
	var b uint8 = 0b11110000 // 240

	c := a ^ b // XOR演算

	fmt.Println(a ^ b)
	fmt.Println(c)
	fmt.Printf("a ^ b = %08b\n", c) // 出力: a ^ b = 01011010
}

// ビットクリア
// https://qiita.com/yamoridon/items/568c2c680e56b24394e6
// a: 10101010
// b: 11110000
// c: 00001010
func bitclearCalc() {
	var a uint8 = 0b10101010 // 170
	var b uint8 = 0b11110000 // 240

	c := a &^ b // ビットクリア

	fmt.Println(a &^ b)
	fmt.Println(c)
	fmt.Printf("a &^ b = %08b\n", c) // 出力: a &^ b = 00001010
}
