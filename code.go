package TextFormer

// Функции интерфейса IText для типа TextString
//
func (j *TextString) Get() string {
	return string(*j)
}

func (txt *TextString) Set(lines ...string) {
	txt.Clean()
	txt.AddTopUD(lines...)
	return
}

func (txt *TextString) SetSplit(topSplit string, bottomSplit string) {
	if topSplit != "" {
		txt.AddDU(topSplit)
	}
	if bottomSplit != "" {
		txt.AddUD(bottomSplit)
	}
	return
}

func (txt *TextString) Clean() {
	*txt = ""
	return
}

func (txt *TextString) AddDU(lines ...string) {
	for _, v := range lines {
		if v != "" {
			*txt = TextString(v + "\n" + txt.Get())
		}
	}
	return
}

func (txt *TextString) AddUD(lines ...string) {
	for _, v := range lines {
		if v != "" {
			*txt += TextString(v + "\n")
		}
	}
	return
}

func (txt *TextString) AddTopUD(lines ...string) {
	for _, v := range lines {
		txt.AddUD(v)
	}
	return
}

func (txt *TextString) AddBottomUD(lines ...string) {
	t := ""
	for _, v := range lines {
		if v != "" {
			t += v + "\n"
		}
	}
	*txt += TextString(t)
	return
}

// Функции интерфейса IText для типа TextTemplate
//
func (j *TextTemplate) Get() string {
	//	var t string
	if (j.TopSplit != "") && (j.BottomSplit != "") {
		return string(j.Top + j.TopSplit + j.Middle + j.BottomSplit + j.Bottom)
	}
	if j.TopSplit != "" {
		return string(j.Top + j.TopSplit + j.Middle + j.Bottom)

	}
	return string(j.Top + j.Middle + j.Bottom)
}

// Задает значения текстовых блоков в следующем порядке: верхний, нижний, средние
func (txt *TextTemplate) Set(lines ...string) {
	txt.Clean()
	for i, v := range lines {
		if v != "" {
			switch i {
			case 0:
				txt.Top.Set(v)
			case 1:
				txt.Bottom.Set(v)
			case 2:
				txt.Middle.Set(v)
			default:
				txt.Middle.AddTopUD(v)
			}
		}
	}
	return
}

func (txt *TextTemplate) SetSplit(topSplit string, bottomSplit string) {
	if topSplit != "" {
		txt.TopSplit.Set(topSplit)
	} else {
		txt.TopSplit.Set("")
	}
	if bottomSplit != "" {
		txt.BottomSplit.Set(bottomSplit)
	} else {
		txt.BottomSplit.Set("")
	}
	return
}

func (txt *TextTemplate) Clean() {
	txt.Bottom.Clean()
	txt.Middle.Clean()
	txt.Top.Clean()
	txt.SetSplit("", "")

	return
}

func (txt *TextTemplate) AddUD(lines ...string) {
	for _, v := range lines {
		txt.Middle.AddUD(v)
	}
	return
}

func (txt *TextTemplate) AddDU(lines ...string) {
	for _, v := range lines {
		txt.Middle.AddDU(v)
	}
	return
}

func (txt *TextTemplate) AddTopUD(lines ...string) {
	for _, v := range lines {
		txt.Top.AddBottomUD(v)
	}
	return
}

func (txt *TextTemplate) AddBottomUD(lines ...string) {
	for _, v := range lines {
		txt.Bottom.AddBottomUD(v)
	}
}
