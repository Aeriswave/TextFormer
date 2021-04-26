package main

// Функции интерфейса IText
//

func (j *StrText) Get() string {
	return string(*j)
}

func (txt *StrText) Set(lines ...string) {
	txt.Clean()
	txt.AddFall(lines...)
	return
}

func (txt *StrText) Split(vSplit string, nSplit string) {
	if vSplit != "" {
		txt.AddFall(vSplit, txt.Get())
	}
	if nSplit != "" {
		txt.AddRise(txt.Get(), nSplit)
	}
	return
}

func (txt *StrText) Clean() {
	*txt = ""
	return
}

func (txt *StrText) AddRise(lines ...string) {
	for _, v := range lines {
		if v != "" {
			*txt = StrText(v + "\n" + txt.Get())
		}
	}
	return
}

func (txt *StrText) AddFall(lines ...string) {
	for _, v := range lines {
		if v != "" {
			*txt += StrText(v + "\n")
		}
	}
	return
}

func (txt *StrText) TopAddFall(lines ...string) {
	for _, v := range lines {
		txt.AddFall(v)
	}
	return
}

func (txt *StrText) SubAddFall(lines ...string) {
	t := ""
	for _, v := range lines {
		if v != "" {
			t += v + "\n"
		}
	}
	*txt += StrText(t)
	return
}

// Функции интерфейса
//
func (j *StrBlock) Get() string {
	return j.Top.Get() + j.TopSplit.Get() + j.Mid.Get() + j.SubSplit.Get() + j.Sub.Get()
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
				txt.Sub.Set(v)
			case 2:
				txt.Mid.Set(v)
			default:
				txt.Mid.TopAddFall(v)
			}
		}
	}
	return
}

func (txt *StrBlock) Split(topSplit string, bottomSplit string) {
	if topSplit != "" {
		txt.TopSplit.Set(topSplit)
	} else {
		txt.TopSplit.Set("")
	}
	if bottomSplit != "" {
		txt.SubSplit.Set(bottomSplit)
	} else {
		txt.SubSplit.Set("")
	}
	return
}

func (txt *StrBlock) Clean() {
	var tmp []StrText = []StrText{"", "", "", "", "", ""}
	var itmp []IText = []IText{&tmp[0], &tmp[1], &tmp[2], &tmp[3], &tmp[4]}
	txt.Top = itmp[0]
	txt.TopSplit = itmp[1]
	txt.Mid = itmp[2]
	txt.SubSplit = itmp[3]
	txt.Sub = itmp[4]

	txt.Sub.Clean()
	txt.Mid.Clean()
	txt.Top.Clean()
	txt.Split("", "")

	return
}

func (txt *StrBlock) AddFall(lines ...string) {
	for _, v := range lines {
		txt.Mid.AddFall(v)
	}
	return
}

func (txt *StrBlock) AddRise(lines ...string) {
	for _, v := range lines {
		txt.Mid.AddRise(v)
	}
	return
}

func (txt *StrBlock) TopAddFall(lines ...string) {
	for _, v := range lines {
		txt.Top.SubAddFall(v)
	}
	return
}

func (txt *StrBlock) SubAddFall(lines ...string) {
	for _, v := range lines {
		txt.Sub.SubAddFall(v)
	}
}
