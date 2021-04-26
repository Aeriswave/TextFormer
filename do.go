package main

import (
	//	_ "TextFormer"
	"fmt"
)

func main() {
	var nn StrBlock = StrBlock{
		Top:      new(StrText),
		TopSplit: new(StrText),
		Mid:      new(StrText),
		SubSplit: new(StrText),
		Sub:      new(StrText)}

	var tt IText = &nn

	//tt = &nn
	tt.Split("", "")
	tt.Set("Верх", "Низ", "Текст посередке")
	tt.TopAddFall("Заголовки", "подзаголовки")
	tt.SubAddFall("нижние строки страницы", "последние строчечки")
	tt.AddRise("текст1 порядок чтения снизу вверх")
	tt.AddRise("текст2 для логов, блогов")
	tt.AddRise("текст3 и для устаревающей информации")
	tt.AddFall("txt1 для чтения сверху вниз")
	tt.AddFall("txt2 для чтения по порядку")
	tt.AddFall("txt3 как в книгах")
	tt.Split("~~Строка разделителя верхнего заголовка~~", "~~Строка разделителя нижнего заголовка страницы ~~")
	fmt.Printf(string(tt.Get())) // Для вывода итогового текста в консоль

	return
}
