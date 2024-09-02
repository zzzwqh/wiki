package main

import "fmt"

type WhereToGo interface {
	WhereDinner()
}
type ChooseDinner interface {
	GoToDinner(dest string)
}

type BeiJingPerson struct {
}
type ShangHaiPerson struct {
}
type SiChanPerson struct {
}
type ChongQingPerson struct {
}

func (BeiJingPerson) WhereDinner() {
	fmt.Println("北京有什么好吃的？")
}
func (ShangHaiPerson) WhereDinner() {
	fmt.Println("上海有什么好吃的？")
}
func (ChongQingPerson) GoToDinner(dest string) {
	if dest == "" {
		fmt.Println("回家吃，我会做重庆小面")
		return
	}
	fmt.Println("吃火锅" + dest + "火锅店")
}
func (SiChanPerson) GoToDinner(dest string) {
	if dest == "" {
		fmt.Println("回家吃，我会做四川火锅")
		return
	}
	fmt.Println("吃火锅" + dest + "火锅店")
}

func main() {
	var personIns1 WhereToGo = ShangHaiPerson{}
	personIns1.WhereDinner()
	var personIns2 WhereToGo = BeiJingPerson{}
	personIns2.WhereDinner()

	var personIns3 ChooseDinner = ChongQingPerson{}
	personIns3.GoToDinner("")
	var personIns4 ChooseDinner = SiChanPerson{}
	personIns4.GoToDinner("蜀大侠")
}
