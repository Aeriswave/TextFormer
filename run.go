package main

import (
	"fmt"
)

// Функции интерфейса IText
//

func (j *StrText) FallList(h int, s int) string {
	if j.header != nil {
		if h != 0 {
			return j.GetText() + j.header.FallList(h-1, s)
		}
	}
	return j.GetText()
}

func (j *StrText) RiseList(h int, s int) string {
	if j.header != nil {
		if h != 0 {
			return j.header.RiseList(h-1, s) + j.GetText()
		}
	}
	return j.GetText()
}

func (j *StrBlock) RiseList(h int, s int) string {
	if j.header != nil {
		if h != 0 {
			return j.header.RiseList(h-1, s) + j.GetText()
		}
	}
	return j.GetText()
}

func (j *StrBlock) FallList(h int, s int) string {
	if j.header != nil {
		if h != 0 {
			return j.GetText() + j.header.FallList(h-1, s)
		}
	}
	return j.GetText()
}

func (j *StrModule) RiseList(h int, s int) string {
	var txt string = j.GetText()
	if j.header != nil {
		if h != 0 {
			j.header.RiseList(h-1, s)
			switch len(j.subText) {
			case 0:
				return txt
				//			case 1:
				//				if s>0 {txt=j.subText[0].GetText()+txt
			default:
				for i, v := range j.subText {
					if i > s && s >= 0 {
						return txt
					}
					txt = v.GetText() + txt
				}
			}
		}
	}
	return txt
}

func (j *StrModule) FallList(h int, s int) string {
	var txt string = j.GetText()
	if j.header != nil {
		if h != 0 {
			j.header.FallList(h-1, s)
			switch len(j.subText) {
			case 0:
				return txt
				//			case 1:
				//				if s>0 {txt=j.subText[0].GetText()+txt
			default:
				for i, v := range j.subText {
					if i > s && s >= 0 {
						return txt
					}
					txt = txt + v.GetText()
				}
			}
		}
	}
	return txt
}

func (j *StrModule) GetText() string {
	if j.text != "" {
		return j.text + "\n"
	}
	return ""
}

func (j *StrModule) Clean() {
	if j.subText != nil {
		for _, v := range j.subText {
			if v != nil {
				v.Clean()
			}
		}
		j.header.Clean()
		j.header = nil
		j.text = ""
		j.subText = nil
	}
	return
}

func (j *StrText) CheckChild(ch IText, ii int) bool {
	return false
}
func (j *StrBlock) CheckChild(ch IText, ii int) bool {
	if ii >= 0 {
		return false
	}
	if j.Mid == ch || j.Sub == ch || j.Top == ch || j.SubSplit == ch || j.TopSplit == ch {
		return true
	}
	return false
}
func (j *StrModule) CheckChild(ch IText, ii int) bool {
	if (ii < 0) || (ii >= len(j.subText)) {
		return false
	}
	if j.subText[ii] == ch {
		return true
	}
	return false
}

func (j *StrModule) Set(pp IText, ii int, lines ...string) {
	j.Clean()
	if pp.CheckChild(j, ii) == false {
		fmt.Print("Неверный указатель и индекс на родительский текстовый блок")
		j.header = nil
		j.index = -1
	} else {
		j.header = pp
		j.index = ii
	}
	switch len(lines) {
	case 0:
		j.text = ""
		j.subText = nil
		return
	case 1:
		j.text = lines[0]
	}
	if len(lines) == 0 {
		j.text = ""
		j.subText = nil
		return
	}
	txt := new(StrText)
	txt.text = lines[0]
	txt.Clean()
	lines = lines[:len(lines)-1]
	txt.AddFall(nil, lines...)
	return
}

func (j *StrModule) Split(string, string) {
	return
}

func (j *StrModule) AddFall(iTxt IText, lines ...string) {
	if lines != nil {
		var arr []StrText
		arr = make([]StrText, len(lines))
		for i, v := range lines {
			arr[i].Set(j, i, v)
			var nn IText = arr[i]
			j.subText = append(j.subText, iTxt, nn)
			//			arr[i] = IText
		}
		if iTxt != nil {
		} else {
			//			j.subText = append(j.subText, arr)
		}
	} else if iTxt != nil {
		j.subText = append(j.subText, iTxt)
	}
	return
}

func (j *StrModule) AddRise(iTxt IText, lines ...string) {
	return
}

func (j *StrModule) TopAddFall(lines ...string) {}

func (j *StrModule) SubAddFall(lines ...string) {}

func (j *StrText) GetText() string {
	return string(j.text)
}
func (txt *StrText) Set(pp IText, ii int, lines ...string) {
	txt.Clean()
	if pp != nil {
		if pp.CheckChild(txt, ii) == false {
			pp = nil
			ii = -1
			fmt.Print("Ошибка при создании простого текста StrText")
			return
		} // ошибка
		txt.header = pp
		txt.index = ii
	}
	txt.AddFall(nil, lines...)
	return
}
func (txt *StrText) Split(vSplit string, nSplit string) {

	if vSplit != "" {
		txt.AddFall(nil, vSplit, txt.GetText())
	}
	if nSplit != "" {
		txt.AddRise(nil, txt.GetText(), nSplit)
	}
	return
}

func (txt *StrText) Clean() {
	txt.text = ""
	return
}

func (txt *StrText) AddRise(iTxt IText, lines ...string) {
	for _, v := range lines {
		if v != "" {
			txt.text = v + "\n" + txt.GetText()
		}
	}
	return
}

func (txt *StrText) AddFall(iTxt IText, lines ...string) {
	for _, v := range lines {
		if v != "" {
			txt.text += v + "\n"
		}
	}
	return
}

func (txt *StrText) TopAddFall(lines ...string) {
	for _, v := range lines {
		txt.AddFall(nil, v)
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
	txt.text += t
	return
}

// Функции интерфейса
//
func (j *StrBlock) GetText() string {
	return j.Top.GetText() + j.TopSplit.GetText() + j.Mid.GetText() + j.SubSplit.GetText() + j.Sub.GetText()
}

// Задает значения текстовых блоков в следующем порядке: верхний, нижний, средние
func (txt *StrBlock) Set(tt IText, ii int, lines ...string) {
	txt.Clean()
	for i, v := range lines {
		if v != "" {
			switch i {
			case 0:
				txt.Top.Set(tt, -1, v)
			case 1:
				txt.Sub.Set(tt, -1, v)
			case 2:
				txt.Mid.Set(tt, -1, v)
			default:
				txt.Mid.TopAddFall(v)
			}
		}
	}
	return
}

func (txt *StrBlock) Split(topSplit string, bottomSplit string) {
	if topSplit != "" {
		txt.TopSplit.Set(txt, -1, topSplit)
	} else {
		txt.TopSplit.Set(txt, -1, "")
	}
	if bottomSplit != "" {
		txt.SubSplit.Set(txt, -1, bottomSplit)
	} else {
		txt.SubSplit.Set(txt, -1, "")
	}
	return
}

func (txt *StrBlock) Clean() {
	var tmp []StrText = []StrText{
		{txt, -1, ""},
		{txt, -1, ""},
		{txt, -1, ""},
		{txt, -1, ""},
		{txt, -1, ""}}
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

func (txt *StrBlock) AddFall(iTxt IText, lines ...string) {
	for _, v := range lines {
		txt.Mid.AddFall(nil, v)
	}
	return
}

func (txt *StrBlock) AddRise(iTxt IText, lines ...string) {
	for _, v := range lines {
		txt.Mid.AddRise(nil, v)
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
