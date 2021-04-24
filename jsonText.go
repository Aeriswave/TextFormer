package jsonText

type TextString string

type TextTemplate struct {
	top         TextString
	topSplit    TextString
	middle      TextString
	bottomSplit TextString
	bottom      TextString
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
	return string(j.top + j.topSplit + j.middle + j.bottomSplit + j.bottom)
}

// Задает значения текстовых блоков в следующем порядке: верхний, нижний, средние
func (txt TextTemplate) Set(lines ...string) {
	//	txt.sss
	for i, v := range lines {
		switch i {
		case 0:
			if v != "" {
				txt.top.Set(v)
			}
		case 1:
			txt.bottom.Set(v)
		case 2:
			txt.middle.Set(v)
		default:
			txt.middle.AddUpLines(v)
		}
	}
	return
}

func (txt TextTemplate) SetSplit(topSplit string, bottomSplit string) {
	if topSplit != "" {
		txt.topSplit.Set("\n" + topSplit + "\n")
	} else {
		txt.topSplit.Set("\n===\n")
	}
	if bottomSplit != "" {
		txt.bottomSplit.Set("\n" + bottomSplit + "\n")
	} else {
		txt.bottomSplit.Set("\n===\n")
	}
	return
}

func (txt TextTemplate) Clean() {
	txt.bottom.Clean()
	txt.middle.Clean()
	txt.top.Clean()
	txt.SetSplit("", "")

	return
}

func (txt TextTemplate) Add(line string) {
	txt.middle.Add(line)
	return
}

func (txt TextTemplate) AddUpLines(lines ...string) {
	for _, v := range lines {
		txt.top.AddSubLines(v)
	}
	return
}

func (txt TextTemplate) AddSubLines(lines ...string) {
	for _, v := range lines {
		txt.bottom.AddSubLines(v)
	}
}
