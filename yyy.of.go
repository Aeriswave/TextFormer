package main

type StrText string
type StrArray []StrText
type StrBlock struct {
	Top      IText
	TopSplit IText
	Mid      []IText
	NizSplit IText
	Niz      IText
}