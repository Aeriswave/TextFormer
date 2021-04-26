package main

type IText interface {
	Get() string
	Set(...string)
	Split(string, string) // установка разделителей текста
	AddFall(...string)    // для чтения сверху вниз
	AddRise(...string)    // чтения снизу вверх
	TopAddFall(...string) // для чтения сверху вниз
	SubAddFall(...string) // для чтения сверху вниз
	Clean()
}
