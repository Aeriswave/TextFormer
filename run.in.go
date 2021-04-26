package main

type IText interface {
	GetText() string                 // Получить текст
	RiseList(h int, size int) string // Получить текст
	FallList(h int, size int) string // Получить текст
	Set(IText, int, ...string)       // задать значение текста
	CheckChild(IText, int) bool      // проверить ссылку на дочерний текст по индексу
	Split(string, string)            // установка разделителей текста
	AddFall(IText, ...string)        // для чтения сверху вниз
	AddRise(IText, ...string)        // чтения снизу вверх
	TopAddFall(...string)            // для чтения сверху вниз
	SubAddFall(...string)            // для чтения сверху вниз
	Clean()
}
