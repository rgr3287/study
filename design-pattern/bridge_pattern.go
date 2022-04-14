package main

import "fmt"

// Bridge 패턴은 두 구현체 간의 결합을 제거하기 위해 (decoupling) 사용되는 패턴입니다.

// 이 패턴의 예제로는 컴퓨터와 프린트기 예제가 많이 사용됩니다.

// 이렇게 맥과 윈도우 컴퓨터가 있고 삼성프린터기와 HP프린터기 구조체를 만들어서 예제를 만들어 보겠습니다.

type MacWithSamsungPrinter struct{}

type MacWithHpPrinter struct{}

type WindowWithSamsungPrinter struct{}

type WindowWithHpPrinter struct{}

// 이렇게 인터페이스 2가지를 만들고
type printer interface {
	printFile()
}

type computer interface {
	print()
	setPrinter(printer)
}

// 삼성 프린터기와 HP 프린터기는 아래처럼 나타낼 수 있습니다.
type samsungPrinter struct{}

func (p *samsungPrinter) printFile() {
	fmt.Println("Printing by a samsung..")
}

type hpPrinter struct{}

func (p *hpPrinter) printFile() {
	fmt.Println("Printing by a hp..")
}

// 맥
type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print in mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

// 윈도우
type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print in windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

// mac과 windows 구조체는 필드에 printer 인터페이스 타입을 가지고 있으므로
// 그렇기 때문에 프린트기를 동작시킬 수 있으며 printer의 인터페이스 타입만 만족시키면 어떤 프린트기도 동작 시킬수 있다.
// 상당히 유연한 구조로 바뀌었습니다.
