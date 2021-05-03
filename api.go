package main

type TextString string

type TextItem struct {
	address TextAddress
	text    TextString
}

type IReader interface {
	GetType() string
	GetText() TextString
}

type IText interface {
	SetParent(IText, int) bool
	checkChild(IText, int) bool

	SetText(IText, ...IText) IText
	GetText() TextString
	GetType() string

	Clean()
	Destroy() // уничтожить этот элемент текста
	//Detonate() // уничтожить всю связанную цепочку элементов текста
	delete(IText, int) bool
}

type TextBlock struct {
	address  TextAddress
	top      IText
	topSplit IText
	mid      IText
	subSplit IText
	sub      IText
}

//type ITextBlock interface {
//	SetText(IText, ...IText) IText // Перезаписать текст этого блока/
//	GetText() TextString           // Получить текст этого блока//
//	GetType() string
//
//	AddFall(...IText) IText // для чтения сверху вниз
//	AddRise(...IText) IText // чтения снизу вверх
//
//	Clean()
//}

type TextAddress struct {
	index  int
	parent IText
	chain  IChain
}

type TextModule struct {
	address   TextAddress
	subRise   map[int]IText
	subCenter int
	//	addRise
	//	newSubCenter >0
	maxIndex int // центральная ячейка массива
	//	newSubCenter <0
	//	addFall
	subFall map[int]IText

	text IText
}

type TextChain struct {
	prev IText
	this TextModule
	next IText
}

type IChain interface {
	Insert(TextModule)
	Split() []TextChain
	// 	InCapsulate // в блок, в модуль или в цепочку
	// 	ReCapsulate // в текст, в блок, в модуль или в цепочку
	Detonate()
}
