package main

type TextString string
type TextBlock []TextString

type TextTemplate struct {
	Top      IText
	TopSplit IText
	Mid      []IText
	NizSplit IText
	Niz      IText
}
