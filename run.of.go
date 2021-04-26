package main

type StrText struct {
	header IText
	index  int
	text   string
}

type StrBlock struct { // переделать на строчные начальные буквы в названиях переменных (приватные)
	header   IText
	index    int
	Top      IText
	TopSplit IText
	Mid      IText
	SubSplit IText
	Sub      IText
	text     string
}

type StrModule struct {
	header  IText
	index   int
	subText []IText
	text    string
}
