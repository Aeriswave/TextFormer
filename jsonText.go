package jsonText

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

type ITextString interface {
	Get() string
	Set(string)
	AddUpperLine(string)
	AddLowerLine(string)
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

func (txt TextString) AddUpperLines(lines ...string) (err error) {
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

func (txt TextString) AddLowerLines(lines ...string) (err error) {
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

func (txt TextTemplate) AddUpperLines(lines ...string) (err error) {
	var t TextString = ""
	for _, v := range lines {
		t.AddUpperLines(v)
	}
	if t != "" {
		txt.AddLowerLines(string(t))
	}
	return nil
}

func (txt TextTemplate) AddLowerLines(lines ...string) (err error) {
	var t TextString = ""
	for _, v := range lines {
		t.AddUpperLines(v)
	}
	if t != "" {
		txt.bottom.AddUpperLines(string(t))
	}
	return nil
}
