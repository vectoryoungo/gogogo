package main

import "fmt"

/*
Go 抽象工厂模式讲解和代码示例
please refer：https://refactoringguru.cn/design-patterns/abstract-factory/go/example#example-0
*/
type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}