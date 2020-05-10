package main

import "testing"

func TestExample(t *testing.T) {

	input := [][]byte{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	}

	square := maximalSquare(input)

	t.Log(square)

	if square == 4 {
		t.Logf("success")
	} else {
		t.Fatal("squaer should be 4")
	}
}

func TestExample2(t *testing.T) {

	//[["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
	//[["0","0","0"],["0","0","0"],["0","0","0"],["0","0","0"]]

	input := [][]byte{
		{0x30, 0x30, 0x30},
		{0x30, 0x30, 0x30},
		{0x30, 0x30, 0x30},
		{0x30, 0x30, 0x30},
	}

	square := maximalSquare(input)

	t.Log(square)

	if square == 0 {
		t.Logf("success")
	} else {
		t.Fatal("squaer should be 0")
	}
}

func TestExample65(t *testing.T) {

	//["0","0","1","0"],
	//["1","1","1","1"],
	//["1","1","1","1"],
	//["1","1","1","0"],
	//["1","1","0","0"],
	//["1","1","1","1"],
	//["1","1","1","0"]]

	input := [][]byte{
		{48, 48, 49, 48},
		{49, 49, 49, 49},
		{49, 49, 49, 49},
		{49, 49, 49, 48},
		{49, 49, 48, 48},
		{49, 49, 49, 49},
		{49, 49, 49, 48},
	}

	square := maximalSquare(input)

	t.Log(square)

	if square == 9 {
		t.Logf("success")
	} else {
		t.Fatal("squaer should be 9")
	}
}
