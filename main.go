package main

import (
	"demo/tree"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!" + "dadaist111")


	root := tree.NewPrefixTree()

	fmt.Println(root)

	root.Insert("abc")
	root.Insert("abcall")
	root.Insert("ab")
	
	fmt.Println(root.SearchPrefix("abcd"))

	fmt.Println(root.Search("abcd"))
	fmt.Println(root.Search("a"))
	fmt.Println(root.Search("abcal"))	
	fmt.Println(root.Search("abc"))
	fmt.Println(root.Search("abcall"))



	fmt.Println("--------------------")
	root.Insert("你好啊")
	root.Insert("你是谁")
	root.Insert("你是？")
	root.Insert("我的")
	root.Insert("我的n好")
	root.Insert("me and you")
	fmt.Println(root.Search("你好啊"))
	fmt.Println(root.Search("你好"))
	fmt.Println(root.Search("me and"))
	fmt.Println(root.SearchPrefix("me and y"))
	fmt.Println(root.SearchPrefix("me and you"))
	
}

