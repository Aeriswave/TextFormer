package jsonText

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
	Add(string)
	SetSplit(string, string)
	AddUpLines(...string)
	AddSubLines(...string)
	//	AddLeft(string)
	//	AddRight(string)
	Clean()
}

func Main() {
	var nn TextTemplate
	var tt IText = nn
	tt.Set("Верх", "Низ", "Середина")
	return
}

// Функции интерфейса iTextString для типа TextString
//
func (j TextString) Get() string {
	return j.Get()
}

func (txt TextString) Set(lines ...string) {
	txt.Clean()
	txt.AddUpLines(lines...)
	return
}

func (txt TextString) Clean() {
	txt += ""
	return
}

func (txt TextString) Add(line string) {
	if line != "" {
		txt = TextString(line) + "\n" + txt
	}
	return
}

func (txt TextString) AddUpLines(lines ...string) {
	t := ""
	for _, v := range lines {
		if v != "" {
			t += v + "\n"
		}
	}
	txt.Add(t)
	return
}

func (txt TextString) AddSubLines(lines ...string) {
	t := ""
	for _, v := range lines {
		if v != "" {
			if v != "" {
				t += v + "\n"
			}
		}
	}
	txt += "\n" + TextString(t)
	return
}

// Функции интерфейса iTextString для типа TextTemplate
//
func (j TextTemplate) Get() string {
	return string(j.Top + j.TopSplit + j.Middle + j.BottomSplit + j.Bottom)
}

// Задает значения текстовых блоков в следующем порядке: верхний, нижний, средние
func (txt TextTemplate) Set(lines ...string) {
	//	txt.sss
	for i, v := range lines {
		switch i {
		case 0:
			if v != "" {
				txt.Top.Set(v)
			}
		case 1:
			txt.Bottom.Set(v)
		case 2:
			txt.Middle.Set(v)
		default:
			txt.Middle.AddUpLines(v)
		}
	}
	return
}

func (txt TextTemplate) SetSplit(topSplit string, bottomSplit string) {
	if topSplit != "" {
		txt.TopSplit.Set("\n" + topSplit + "\n")
	} else {
		txt.TopSplit.Set("\n===\n")
	}
	if bottomSplit != "" {
		txt.BottomSplit.Set("\n" + bottomSplit + "\n")
	} else {
		txt.BottomSplit.Set("\n===\n")
	}
	return
}

func (txt TextTemplate) Clean() {
	txt.Bottom.Clean()
	txt.Middle.Clean()
	txt.Top.Clean()
	txt.SetSplit("", "")

	return
}

func (txt TextTemplate) Add(line string) {
	txt.Middle.Add(line)
	return
}

func (txt TextTemplate) AddUpLines(lines ...string) {
	for _, v := range lines {
		txt.Top.AddSubLines(v)
	}
	return
}

func (txt TextTemplate) AddSubLines(lines ...string) {
	for _, v := range lines {
		txt.Bottom.AddSubLines(v)
	}
}
