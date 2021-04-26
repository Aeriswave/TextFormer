package main

import (
	"fmt"
)

func main() {

	// Текстовый блок
	var nn StrBlock = StrBlock{
		Top:      new(StrText),
		TopSplit: new(StrText),
		Mid:      new(StrBlock),
		SubSplit: new(StrText),
		Sub:      new(StrModule)}

	// API для работы с тектовым блоком
	var tt IText = &nn

	tt.Split("", "")
	tt.Set(tt, -1, "Верх", "Низ", "Текст посередке")
	tt.TopAddFall("Заголовки", "подзаголовки")
	tt.SubAddFall("нижние строки страницы", "последние строчечки")
	tt.AddRise(nil, "текст1 порядок чтения снизу вверх")
	tt.AddRise(nil, "текст2 для логов, блогов")
	tt.AddRise(nil, "текст3 и для устаревающей информации")
	tt.AddFall(nil, "txt1 для чтения сверху вниз")
	tt.AddFall(nil, "txt2 для чтения по порядку")
	tt.AddFall(nil, "txt3 как в книгах")
	tt.Split("~~Строка разделителя верхнего заголовка~~", "~~Строка разделителя нижнего заголовка страницы ~~")
	fmt.Printf(string(tt.GetText()))           // Для вывода итогового текста в консоль
	fmt.Printf("Для завершения нажмите Enter") // Для вывода итогового текста в консоль
	esc := ""
	fmt.Scanf("%s", &esc)
	fmt.Printf("\nПока!\n") // Для вывода итогового текста в консоль
	return
}
