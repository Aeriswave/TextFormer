package main

func aparser(c chan<- string) {
	//	несколько строчек
	var q1 TextItem = TextItem{text: "txt1 для чтения сверху вниз"}
	var q2 TextItem = TextItem{text: "txt2 написано для чтения сверху вниз"}
	var q3 TextItem = TextItem{text: "txt3 как в книгах"}
	q11 := []TextItem{
		{text: "текст1 порядок написания снизу вверх"},
		{text: "текст2 для логов, блогов"},
		{text: "текст3 и для устаревающей информации"},
	}

	var qq TextBlock = TextBlock{}
	qq.init()
	qq.top.NewText("верх")
	qq.text.NewText("середина")
	qq.sub.NewText("низ")
	qq.topSplit.NewText("+++верхний разделитель+++")
	qq.subSplit.NewText("---нижний разделитель---")
	qq.AddFall(nil, nil, nil, nil, &q1, &q2, &q3)
	var w IText = qq.AddRise(nil, nil, nil, nil, &q11[0], &q11[1], &q11[2])
	//	w.NewText("1111", "3", "4", "2222")

	c <- string(w.GetFullText())
	w.Clean()
	c <- "--\n"

	var ee TextModule = TextModule{}
	ee.init()
	w = ee.NewText("top", "---", "Второй модуль текста", "===", "sub")
	//	ee.subRise[4].Destroy()
	//	fmt.Println("=-=-=-")
	c <- string(w.GetFullText()) // вывод итогового текста в консоль с блокировкой канала до чтения данных из него
	//	fmt.Println("_+_+_+")
	//	w.Destroy()
	w.Clean()
}
