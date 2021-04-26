package TextFormer

/*
import ("fmt")
func main() {
	var tmp TextString = ""
	var nn TextTemplate = TextTemplate{
		Top:         tmp,
		TopSplit:    tmp,
		Middle:      tmp,
		BottomSplit: tmp,
		Bottom:      tmp}

	var tt IText
	tt = &nn
	var ts IText
	ts = &tmp
	ts.SetSplit("", "")
	tt.Set("Верх", "Низ", "Текст посередке")
	tt.AddTopUD("Заголовки", "подзаголовки")
	tt.AddBottomUD("нижние строки страницы", "последние строчечки")
	tt.AddDU("текст1 порядок чтения снизу вверх")
	tt.AddDU("текст2 для логов, блогов")
	tt.AddDU("текст3 и для устаревающей информации")
	tt.AddUD("txt1 для чтения сверху вниз")
	tt.AddUD("txt2 для чтения по порядку")
	tt.AddUD("txt3 как в книгах")
	tt.SetSplit("~~Строка разделителя верхних заголовков и текста~~", "~~Строка разделителя текста и нижних строк страницы ~~")
	fmt.Printf(string(tt.Get())) // Для вывода итогового текста в консоль
	return
}
*/

type TextString string

type TextTemplate struct {
	Top         TextString
	TopSplit    TextString
	Middle      TextString
	BottomSplit TextString
	Bottom      TextString
}

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
