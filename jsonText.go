package JSONText

// Public visibility
var Test2 = "public variable"

type TextString string

type TextTemplate struct {
	top        TextString
	topLine    TextString
	middle     TextString
	bottomLine TextString
	bottom     TextString
}

type iTextString interface {
	Get() string
	Set(string)
	AddUpLine(string)
	AddDownLine(string)
	AddLeft(string)
	AddRight(string)
	Clean()
}

func Main() {
	return
}

// Функции интерфейса iTextString для типа TextString
//
func (j TextString) Get() string {
	return string(j)
}

func (txt TextString) Set(line string) (err error) {
	txt = TextString(line)
	return nil
}

func (txt TextString) Clean() (err error) {
	txt += ""
	return nil
}

func (txt TextString) AddUpLines(lines ...string) (err error) {
	t := ""
	for _, v := range lines {
		if v != "" {
			t += v + "\n"
		}
	}
	if t != "" {
		txt = TextString(t) + "\n" + txt
	}
	return nil
}

func (txt TextString) AddDownLines(lines ...string) (err error) {
	t := ""
	for _, v := range lines {
		if v != "" {
			if v != "" {
				t += v + "\n"
			}
		}
	}
	txt += "\n" + TextString(t)
	return nil
}

// Функции интерфейса iTextString для типа TextTemplate
//
func (j TextTemplate) Get() string {
	return string(j.top + j.topLine + j.middle + j.bottomLine + j.bottom)
}

func (txt TextTemplate) Set(topic string, middleText string, endText string) (err error) {
	txt.top.Set(topic)
	txt.middle.Set(middleText)
	txt.bottom.Set(endText)
	return nil
}

func (txt TextTemplate) SetLine(topLine string, bottomLine string) (err error) {
	if topLine != "" {
		txt.topLine.Set("\n" + topLine + "\n")
	} else {
		txt.topLine.Set("\n===\n")
	}
	if bottomLine != "" {
		txt.bottomLine.Set("\n" + bottomLine + "\n")
	} else {
		txt.bottomLine.Set("\n===\n")
	}
	return nil
}

func (txt TextTemplate) Clean() (err error) {
	txt.bottom.Clean()
	txt.middle.Clean()
	txt.top.Clean()
	return nil
}

func (txt TextTemplate) AddUpLines(lines ...string) (err error) {
	var t TextString = ""
	for _, v := range lines {
		t.AddUpLines(v)
	}
	if t != "" {
		txt.AddDownLines(string(t))
	}
	return nil
}

func (txt TextTemplate) AddDownLines(lines ...string) (err error) {
	var t TextString = ""
	for _, v := range lines {
		t.AddUpLines(v)
	}
	if t != "" {
		txt.bottom.AddUpLines(string(t))
	}
	return nil
}
