package main

import "fmt"

/**
抽象工厂
*/

/*
 在小明的学校，每一年开学都会发教材，
主要包括语文书、数学书、英语书，还有各种练习试卷。
这一天，小明去领了三本教材，分别是语文书、数学书和英语书，老师忙不过来，指定某个同学去发书，
同学们都去这个同学这里去领书。这个同学就是工厂。
*/

type Book interface {
	Name() string
}

type Paper interface {
	Name() string
}

type chineseBook struct {
	name string
}

type chinesePaper struct {
	name string
}

func (cb *chineseBook) Name() string {
	return cb.name
}

type mathBook struct {
	name string
}

func (mb *mathBook) Name() string {
	return mb.name
}

type englishBook struct {
	name string
}

func (eb *englishBook) Name() string {
	return eb.name
}

// 发书人
type Assigner interface {
	GetBook(name string) Book
	GetPaper(string) Paper
}

type assigner struct{}

func (a *assigner) GetBook(name string) Book {
	if name == "语文书" {
		return &chineseBook{name: "语文书"}
	} else if name == "数学书" {
		return &mathBook{name: "数学书"}
	} else if name == "英语书" {
		return &englishBook{name: "英语书"}
	}
	return nil
}

type chineseBookAssigner struct{}

func (cba *chineseBookAssigner) GetBook(name string) Book {
	if name == "语文书" {
		return &chineseBook{
			name: "语文书",
		}
	}
	return nil
}

func main() {
	var a Assigner
	fmt.Println(a.GetBook("语文书").Name())
}
