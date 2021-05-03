package main

func (self *TextItem) Clean() {
	self.text = ""
}

func (self *TextItem) Destroy() {
	self.Clean()
	self.address.parent.delete(self, self.address.hash)
	self.address.parent = nil
	self.address.hash = Hash{this: 0, next: 0, prev: 0}
	return
}

func (self *TextBlock) Destroy() {
	self.Clean()
	self.address.parent.delete(self, self.address.hash)
	self.address.parent = nil
	self.address.hash = Hash{this: 0, next: 0, prev: 0}
	return
}

func (self *TextModule) Destroy() {
	self.Clean()
	self.address.parent.delete(self, self.address.hash)
	self.address.parent = nil
	self.address.hash = Hash{this: 0, next: 0, prev: 0}
	return
}

func (self *TextItem) delete(child IText, i Hash) bool {
	return false
}

func (self *TextBlock) delete(child IText, i Hash) bool {
	if self.checkChild(child, i.this) {
		switch i.this {
		case -2:
			self.sub = nil
			return true
		case -1:
			self.subSplit = nil
			return true
		case 0:
			self.text = nil
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

func (self *TextModule) delete(child IText, i Hash) bool {
	if self.checkChild(child, i.this) {
		if i.this > 0 {
			if self.maxRise.getHash().this == i.this {
				self.maxRise = self.subRise[i.prev]
				self.subRise[i.prev].SetParent(self, Hash{
					next: i.this,
					this: self.subRise[i.prev].getHash().this,
					prev: self.subRise[i.prev].getHash().prev})
			} else {
				self.subRise[i.prev].SetParent(self, Hash{
					next: i.next,
					this: self.subRise[i.prev].getHash().this,
					prev: self.subRise[i.prev].getHash().prev})
				self.subRise[i.next].SetParent(self, Hash{
					prev: i.prev,
					this: self.subRise[i.next].getHash().this,
					next: self.subRise[i.next].getHash().prev})
			}
			delete(self.subRise, i.this)
			return true
		} else if i.this < 0 {
			if self.minFall.getHash().this == i.this {
				self.minFall = self.subFall[i.prev]
				self.subFall[i.prev].SetParent(self, Hash{next: i.this, this: self.subFall[i.prev].getHash().this, prev: self.subRise[i.prev].getHash().prev})
			} else {
				self.subFall[i.prev].SetParent(self, Hash{next: i.next, this: self.subFall[i.prev].getHash().this, prev: self.subRise[i.prev].getHash().prev})
				self.subFall[i.next].SetParent(self, Hash{prev: i.prev, this: self.subFall[i.next].getHash().this, next: self.subRise[i.next].getHash().prev})
			}
			delete(self.subFall, i.this)
			return true
		}
	}
	return false
}

func (self *TextItem) SetParent(parent IText, i Hash) bool {
	if parent.checkChild(self, i.this) {
		self.address.parent = parent
		self.address.hash = i
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
		if self.text == child {
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

func (self *TextBlock) SetParent(parent IText, i Hash) bool {
	self.address.parent = parent
	self.address.hash = i
	return false
}

func (self *TextModule) SetParent(parent IText, hash Hash) bool {
	self.address.parent = parent
	self.address.hash = hash
	return false
}

func (self *TextBlock) Clean() {
	if self.text != nil {
		self.text.Clean()
		self.text = nil
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
		self.text.SetText(m)
	}
	if len(s) > 0 {
		var tt TextItem = TextItem{}
		tt.text = ""
		if self.text != nil {
			tt.text += self.text.GetText()
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
		self.text.SetText(self, &tt)
	}
	return self
}

func (self *TextBlock) GetText() TextString {
	return self.top.GetText() + self.topSplit.GetText() + self.text.GetText() + self.subSplit.GetText() + self.sub.GetText()
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

func (self *TextItem) GetFullText() TextString {
	return self.GetText()
}

func (self *TextBlock) GetFullText() TextString {
	return self.GetText()
}

func (self *TextModule) GetFullText() TextString {
	var txt TextString = ""
	txt = self.text.GetFullText()
	for _, v := range self.subFall {
		txt += v.GetFullText()
	}
	for _, v := range self.subRise {
		txt += v.GetFullText()
	}
	return self.text.GetFullText()
}

// Добавить нисходящий текст
func (self *TextBlock) AddFall(s ...IText) IText {
	if len(s) > 0 {
		var tt TextItem = TextItem{}
		tt.text = ""
		if self.text != nil {
			tt.text += self.text.GetText()
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
		self.text.SetText(&tt)
	}
	return self
}

func (self *TextModule) AddFall(s ...IText) IText {
	switch len(s) {
	case 0:
	default:
		for _, v := range s {
			if v != nil {
				var index int = 0
				if self.minFall != nil {
					index = self.minFall.getHash().this
					self.subFall[0] = v
					self.subFall[0].SetParent(self,
						Hash{
							prev: 0,
							this: 0,
							next: 0,
						})
					self.minFall = self.subFall[0]
				} else {
					self.subFall[index-1] = v
					self.subFall[index-1].SetParent(self,
						Hash{
							prev: self.minFall.getHash().this,
							this: index - 1,
							next: index - 1,
						})
					self.minFall.SetParent(self,
						Hash{
							prev: self.minFall.getHash().prev,
							this: self.minFall.getHash().this,
							next: index - 1,
						})
					self.minFall = self.subFall[index-1]
				}
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
				var index int = 0
				if self.maxRise != nil {
					index = self.maxRise.getHash().this
					self.subRise[0] = v
					self.subRise[0].SetParent(self,
						Hash{
							prev: 0,
							this: 0,
							next: 0,
						})
					self.maxRise = self.subRise[0]
				} else {
					self.subRise[index+1] = v
					self.subRise[index+1].SetParent(self,
						Hash{
							prev: self.maxRise.getHash().this,
							this: index + 1,
							next: index + 1,
						})
					self.maxRise.SetParent(self,
						Hash{
							prev: self.maxRise.getHash().prev,
							this: self.maxRise.getHash().this,
							next: index + 1,
						})
					self.maxRise = self.subRise[index+1]
				}
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
		if self.text != nil {
			tt.text = self.text.GetText() + tt.text
		}
		self.text.SetText(&tt)
	}
	return self
}

func (self *TextItem) getHash() Hash {
	return self.address.hash
}

func (self *TextBlock) getHash() Hash {
	return self.address.hash
}

func (self *TextModule) getHash() Hash {
	return self.address.hash
}
