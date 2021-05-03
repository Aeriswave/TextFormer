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
	SetParent(IText, Hash) bool
	checkChild(IText, int) bool
	getHash() Hash

	SetText(IText, ...IText) IText
	GetText() TextString
	GetFullText() TextString
	GetType() string

	Clean()
	Destroy() // уничтожить этот элемент текста
	//Detonate() // уничтожить всю связанную цепочку элементов текста
	delete(IText, Hash) bool
}

type TextBlock struct {
	address TextAddress
	text    IText

	top      IText
	topSplit IText
	subSplit IText
	sub      IText
}

type TextAddress struct {
	hash     Hash
	nextHash int
	prevHash int
	parent   IText
	chain    IChain
}

type Hash struct {
	prev int
	this int
	next int
}

type TextModule struct {
	address TextAddress
	text    IText

	subRise map[int]IText
	maxRise IText
	subFall map[int]IText
	minFall IText
}

type TextChain struct {
	prev IText
	this IText
	next IText
}

type IChain interface {
	Insert(TextModule)
	Split() []TextChain
	// 	InCapsulate // в блок, в модуль или в цепочку
	// 	ReCapsulate // в текст, в блок, в модуль или в цепочку
	Detonate()
}
