package main

func (self *TextItem) Clean() {
	self.text = ""
}

func (self *TextItem) Destroy() {
	self.Clean()
	self.address.parent.delete(self, self.address.index)
	self.address.parent = nil
	self.address.index = 0
	return
}

func (self *TextBlock) Destroy() {
	self.Clean()
	self.address.parent.delete(self, self.address.index)
	self.address.parent = nil
	self.address.index = 0
	return
}

func (self *TextModule) Destroy() {
	self.Clean()
	self.address.parent.delete(self, self.address.index)
	self.address.parent = nil
	self.address.index = 0
	return
}

func (self *TextItem) delete(child IText, i int) bool {
	return false
}

func (self *TextBlock) delete(child IText, i int) bool {
	if self.checkChild(child, i) {
		switch i {
		case -2:
			self.sub = nil
			return true
		case -1:
			self.subSplit = nil
			return true
		case 0:
			self.mid = nil
			return true
		case 1:
			self.topSplit = nil
			return true
		case 2:
			self.top = nil
			return true
		}
	}
	return false
}

func (self *TextModule) delete(child IText, i int) bool {
	if self.checkChild(child, i) {
		if i > 0 {
			delete(self.subRise, i-1)
			return true
		} else if i < 0 {
			delete(self.subRise, -i-1)
			return true
		}
	}
	return false
}

func (self *TextItem) SetParent(parent IText, i int) bool {
	if parent.checkChild(self, i) {
		self.address.parent = parent
		self.address.index = i
		return true
	}
	return false
}

func (self *TextItem) checkChild(parent IText, i int) bool {
	return false
}

func (self *TextBlock) checkChild(child IText, i int) bool {
	switch i {
	case -2:
		if self.sub == child {
			return true
		}
	case -1:
		if self.subSplit == child {
			return true
		}
	case 0:
		if self.mid == child {
			return true
		}
	case 1:
		if self.topSplit == child {
			return true
		}
	case 2:
		if self.top == child {
			return true
		}
	}
	return false
}

func (self *TextModule) checkChild(child IText, i int) bool {
	if i > 0 && i <= len(self.subRise) {
		if self.subRise[i-1] == child {
			return true
		}
	} else if i < 0 && -i <= len(self.subRise) {
		if self.subFall[-i-1] == child {
			return true
		}
	}

	return false
}

func (self *TextItem) SetText(m IText, s ...IText) IText {
	var t TextString = ""
	if m != nil {
		t = m.GetText()
	}
	if len(s) > 0 {
		for _, v := range s {
			if v != nil {
				t = t + v.GetText()
			}
		}
	}
	self.text = t
	return self
}

func (self *TextItem) GetText() TextString {
	return self.text
}

func (self *TextBlock) SetParent(parent IText, i int) bool {
	self.address.parent = parent
	self.address.index = i
	return false
}

func (self *TextModule) SetParent(parent IText, i int) bool {
	self.address.parent = parent
	self.address.index = i
	return false
}

func (self *TextBlock) Clean() {
	if self.mid != nil {
		self.mid.Clean()
		self.mid = nil
	}
	if self.top != nil {
		self.top.Clean()
		self.top = nil
	}
	if self.sub != nil {
		self.sub.Clean()
		self.sub = nil
	}
	if self.topSplit != nil {
		self.topSplit.Clean()
		self.topSplit = nil
	}
	if self.subSplit != nil {
		self.subSplit.Clean()
		self.subSplit = nil
	}
}

func (self *TextBlock) SetText(m IText, s ...IText) IText {
	// заменить текст в середине блока
	if m != nil {
		self.mid.SetText(m)
	}
	if len(s) > 0 {
		var tt TextItem = TextItem{}
		tt.text = ""
		if self.mid != nil {
			tt.text += self.mid.GetText()
		}
		for i, v := range s {
			if v != nil {
				switch i {
				case 0: // заменить текст наверху блока
					self.top.SetText(v)
				case 1: // заменить текст верхнего разделителя блока
					self.topSplit.SetText(v)
				case 2: // заменить текст нижнего разделителя блока
					self.subSplit.SetText(v)
				case 3: // заменить текст внизу блока
					self.sub.SetText(v)
				default: // добавить текст в середину блока
					tt.text += v.GetText()
				}
			}
		}
		self.mid.SetText(self, &tt)
	}
	return self
}

func (self *TextBlock) GetText() TextString {
	return self.top.GetText() + self.topSplit.GetText() + self.mid.GetText() + self.subSplit.GetText() + self.sub.GetText()
}

func (self *TextModule) Clean() {
	// очистить модуль
	return
}

func (self *TextModule) ReCreate() TextModule {
	// переписать массивы модуля, удалив nil-значения из массивов
	var m TextModule = TextModule{}
	for _, v := range self.subFall {
		if v != nil {
			m.AddFall(v)
		}
	}
	for _, vv := range self.subRise {
		if vv != nil {
			m.AddRise(vv)
		}
	}
	self.subFall = m.subFall
	self.subRise = m.subRise
	return *self
}

func (self *TextModule) SetText(m IText, s ...IText) IText {
	self.text.SetText(m, s...)
	return self
}

func (self *TextModule) GetText() TextString {
	return self.text.GetText()
}

// Добавить нисходящий текст
func (self *TextBlock) AddFall(s ...IText) IText {
	if len(s) > 0 {
		var tt TextItem = TextItem{}
		tt.text = ""
		if self.mid != nil {
			tt.text += self.mid.GetText()
		}
		for i, v := range s {
			if v != nil {
				switch i {
				case 0: // заменить текст наверху блока
					self.top.SetText(v)
				case 1: // заменить текст верхнего разделителя блока
					self.topSplit.SetText(v)
				case 2: // заменить текст нижнего разделителя блока
					self.subSplit.SetText(v)
				case 3: // заменить текст внизу блока
					self.sub.SetText(v)
				default: // добавить текст в середину блока
					tt.text += v.GetText()
				}
			}
		}
		self.mid.SetText(&tt)
	}
	return self
}

func (self *TextModule) AddFall(s ...IText) IText {
	switch len(s) {
	case 0:
	default:
		for _, v := range s {
			if v != nil {
				if self.maxIndex > 0 {
					self.maxIndex = -self.maxIndex
				}
				self.maxIndex--
				self.subFall[self.maxIndex] = v
				self.subRise[self.maxIndex].SetParent(self, self.maxIndex)
			}
		}
	}
	return self
}

func (self *TextModule) AddRise(s ...IText) IText {
	switch len(s) {
	case 0:
	default:
		for _, v := range s {
			if v != nil {
				if self.maxIndex < 0 {
					self.maxIndex = -self.maxIndex
				}
				self.maxIndex++
				self.subRise[self.maxIndex] = v
				self.subRise[self.maxIndex].SetParent(self, self.maxIndex)
			}
		}
	}
	return self
}

func (self *TextItem) GetType() string {
	return "TextString"
}

func (self *TextBlock) GetType() string {
	return "TextBlock"
}

func (self *TextModule) GetType() string {
	return "TextModule"
}

// Добавить восходящий текст
func (self *TextBlock) AddRise(s ...IText) IText {
	if len(s) > 0 {
		var tt TextItem = TextItem{}
		tt.text = ""
		for _, v := range s {
			if v != nil {
				tt.text = v.GetText() + tt.text
			}
		}
		if self.mid != nil {
			tt.text = self.mid.GetText() + tt.text
		}
		self.mid.SetText(&tt)
	}
	return self
}
