package main

type TextString string

type TextAddress struct {
	hash   Hash
	parent IText
	chain  IChain // Пока не реализовано
}

type TextItem struct {
	address TextAddress
	text    TextString
}

type TextBlock struct {
	address TextAddress
	text    IText

	top      IText
	topSplit IText
	subSplit IText
	sub      IText
}

type TextModule struct {
	address TextAddress
	text    IText

	subRise map[int]IText
	maxRise IText
	subFall map[int]IText
	minFall IText
}

type IUser interface {
	NewText(...string) IText
	GetText() TextString
	GetFullText() TextString
}

type IText interface {
	NewText(...string) IText

	SetParent(IText, Hash) bool
	checkChild(IText, int) bool

	SetText(IText, ...IText) IText
	GetText() TextString
	GetFullText() TextString
	GetType() string

	Clean()
	Destroy() // уничтожить этот элемент текста
	//Detonate() // уничтожить всю связанную цепочку элементов текста

	getHash() Hash
	delete(IText, Hash) bool
}

type Hash struct {
	prev int
	this int
	next int
}

type TextChain struct { // Пока не реализовано
	prev IText
	this IText
	next IText
}

type IChain interface { // Пока не реализовано
	Insert(TextModule)
	Split() []TextChain
	// 	InCapsulate // в блок, в модуль или в цепочку
	// 	ReCapsulate // в текст, в блок, в модуль или в цепочку
	Detonate()
}
