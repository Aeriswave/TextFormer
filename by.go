package main

type IText interface {
	Get() string
	Set(...string)
	AddUD(...string)
	AddDU(...string)
	SetSplit(string, string)
	AddTopUD(...string)
	AddBottomUD(...string)
	Clean()
}
