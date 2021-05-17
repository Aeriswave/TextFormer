package main

type Mainer string

type CW int

type CWCH struct {
	//	runIN   chan<- string // поток для приема сопрограммой сообщений синхронизации
	//	runOUT  <-chan string // поток для отправки сообщений синхронизации
	//	setIN   chan<- string // поток для управления этой сопрограммой
	//	setOUT  <-chan string // поток для отправки управляющих сообщений другим сопрограммам
	//	doIN    chan<- string // поток для получения команд для исполнения
	//	doOUT   <-chan string // поток для отправки команд для исполнения другим сопрограммам
	data chan string // поток для получения данных
	//	dataOUT <-chan string // поток для отправки данных
	logOUT <-chan string // поток для отправки сообщений в лог
}

type ICW interface {
	Init(func(CWCH)) (chs CWCH, log string)
}

type CWList struct {
	index       int
	coworks     int
	cowoIDsList map[int]string
	coworksList map[string]CW
}
type ICWList interface {
	Create(name string, f func(CWCH)) (icw CWCH, err string)
	Destroy(name string) (icw ICW, err string)
}

/*
type ISetter interface {
}

type IRunner interface {
}

type IDriver interface {
}

type IKeeper interface {
}
*/

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

type TextArray struct {
	address TextAddress
	text    IText
	array   map[int]IText
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
