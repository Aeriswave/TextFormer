package main

// Функции интерфейса IText для типа TextString
//
func (j *StrText) Get() string {
	return string(*j)
}

func (txt *StrText) Set(lines ...string) {
	txt.Clean()
	txt.AddTopUD(lines...)
	return
}

func (txt *StrText) SetSplit(topSplit string, bottomSplit string) {
	if topSplit != "" {
		txt.AddDU(topSplit)
	}
	if bottomSplit != "" {
		txt.AddUD(bottomSplit)
	}
	return
}

func (txt *StrText) Clean() {
	*txt = ""
	return
}

func (txt *StrText) AddDU(lines ...string) {
	for _, v := range lines {
		if v != "" {
			*txt = StrText(v + "\n" + txt.Get())
		}
	}
	return
}

func (txt *StrText) AddUD(lines ...string) {
	for _, v := range lines {
		if v != "" {
			*txt += StrText(v + "\n")
		}
	}
	return
}

func (txt *StrText) AddTopUD(lines ...string) {
	for _, v := range lines {
		txt.AddUD(v)
	}
	return
}

func (txt *StrText) AddBottomUD(lines ...string) {
	t := ""
	for _, v := range lines {
		if v != "" {
			t += v + "\n"
		}
	}
	*txt += StrText(t)
	return
}

// Функции интерфейса IText для типа TextTemplate
//
func (j *StrBlock) Get() string {
	return j.Top.Get() + j.TopSplit.Get() + j.Mid[0].Get() + j.NizSplit.Get() + j.Niz.Get()
}

// Задает значения текстовых блоков в следующем порядке: верхний, нижний, средние
func (txt *StrBlock) Set(lines ...string) {
	txt.Clean()
	for i, v := range lines {
		if v != "" {
			switch i {
			case 0:
				txt.Top.Set(v)
			case 1:
				txt.Niz.Set(v)
			case 2:
				txt.Mid[0].Set(v)
			default:
				txt.Mid[0].AddTopUD(v)
			}
		}
	}
	return
}

func (txt *StrBlock) SetSplit(topSplit string, bottomSplit string) {
	if topSplit != "" {
		txt.TopSplit.Set(topSplit)
	} else {
		txt.TopSplit.Set("")
	}
	if bottomSplit != "" {
		txt.NizSplit.Set(bottomSplit)
	} else {
		txt.NizSplit.Set("")
	}
	return
}

func (txt *StrBlock) Clean() {
	txt.Niz.Clean()
	txt.Mid[0].Clean()
	txt.Top.Clean()
	txt.SetSplit("", "")

	return
}

func (txt *StrBlock) AddUD(lines ...string) {
	for _, v := range lines {
		txt.Mid[0].AddUD(v)
	}
	return
}

func (txt *StrBlock) AddDU(lines ...string) {
	for _, v := range lines {
		txt.Mid[0].AddDU(v)
	}
	return
}

func (txt *StrBlock) AddTopUD(lines ...string) {
	for _, v := range lines {
		txt.Top.AddBottomUD(v)
	}
	return
}

func (txt *StrBlock) AddBottomUD(lines ...string) {
	for _, v := range lines {
		txt.Niz.AddBottomUD(v)
	}
}
