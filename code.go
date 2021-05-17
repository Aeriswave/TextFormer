package main

import (
	"fmt"
)

func (self *TextItem) NewText(text ...string) IText {
	t := ""
	for _, v := range text {
		t = t + v + "\n"
	}
	self.text = TextString(t)
	return nil
}

func (self *TextBlock) NewText(txt ...string) IText {
	self.Clean()
	if len(txt) > 0 {
		var tt TextItem = TextItem{
			text: TextString(txt[0])}
		var ttt []IText = make([]IText, len(txt)-1)
		for i, v := range txt {
			if i > 0 {
				ttt[i-1] = &TextItem{
					text: TextString(v),
				}
				//fmt.Printf("xx %s xx\n", tt.GetText())
			}
		}
		self.SetText(&tt, ttt...)
	}
	return self
}

func (self *TextModule) NewText(text ...string) IText {
	var b []TextBlock = make([]TextBlock, len(text))
	var ib []IText = make([]IText, len(text))
	for i, v := range text {
		b[i].NewText(v)
		ib[i] = &b[i]
	}
	self.AddRise(ib...)
	return self
}

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

func (self *TextItem) init() {

}

func (self *TextItem) SetText(m IText, s ...IText) IText {
	self.init()
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
	if len(self.text) > 0 {
		if self.text[len(self.text)-1] != '\n' {
			self.text = self.text + "\n"
		}
	}
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

func (self *TextModule) init() {
	if self.text == nil {
		self.text = &TextItem{}
	}
}

func (self *TextBlock) init() {
	if self.text == nil {
		self.text = &TextItem{}
	}
	if self.top == nil {
		self.top = &TextItem{}
	}
	if self.sub == nil {
		self.sub = &TextItem{}
	}
	if self.topSplit == nil {
		self.topSplit = &TextItem{}
	}
	if self.subSplit == nil {
		self.subSplit = &TextItem{}
	}
}

func (self *TextBlock) SetText(m IText, s ...IText) IText {
	self.init()
	if len(s) > 0 {
		var tt TextItem = TextItem{}
		tt.text = ""
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
		self.text.SetText(m, &tt)
	} else {
		// заменить текст в середине блока
		if m != nil {
			self.text.SetText(m)
		}
	}
	return self
}

func (self *TextBlock) GetText() TextString {
	var tt TextString = ""
	if self.top != nil {
		tt += self.top.GetText()
	}
	if self.topSplit != nil {
		tt += self.topSplit.GetText()
	}
	if self.text != nil {
		tt += self.text.GetText()
	}
	if self.subSplit != nil {
		tt += self.subSplit.GetText()
	}
	if self.sub != nil {
		tt += self.sub.GetText()
	}
	return tt
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
	self.init()
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

	var riseTxt TextString = ""
	if self.maxRise != nil {
		var n int = self.maxRise.getHash().this
		//		riseTxt = self.maxRise.GetText() + riseTxt
		for iii := 0; iii >= 0; iii++ {
			if self.subRise != nil {
				if self.subRise[n] != nil {
					riseTxt = self.subRise[n].GetText() + riseTxt
					if n == self.subRise[n].getHash().prev {
						iii = -100
					} else {
						n = self.subRise[n].getHash().prev
					}
				} else {
					iii = -100
					fmt.Printf("опять ошибка")
				}
			} else {
				iii = -100
				fmt.Printf("ошибка")
			}
		}
	}
	//	for _, v := range self.subRise {
	//	n = self.subRise[n].getHash().prev
	//	txt += self.subRise[n].GetText()
	//	}
	return txt + riseTxt
}

// Добавить нисходящий текст
func (self *TextBlock) AddFall(s ...IText) IText {
	self.init()
	if len(s) > 0 {
		var tt TextItem = TextItem{}
		tt.text = ""
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
		if self.text != nil {
			tt.text += self.text.GetText()
		}
		self.text.SetText(&tt)
	}
	return self
}

func (self *TextModule) AddFall(s ...IText) IText {
	self.init()
	var l int = len(s)
	switch l {
	case 0:
	default:
		var index int = 0
		for i, v := range s {
			if v != nil {
				if self.minFall != nil {
					index = self.minFall.getHash().this
				} else {
					self.subFall = make(map[int]IText, len(s))
					self.minFall = IText(s[i])
					self.minFall.SetParent(
						self,
						Hash{this: 0, next: 0, prev: 0},
					)
				}
				self.subFall[index-1] = s[l-i-1]
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
	return self
}

func (self *TextModule) AddRise(s ...IText) IText {
	self.init()
	var l int = len(s)
	switch l {
	case 0:
	default:
		var index int = 0
		for i, v := range s {
			if v != nil {
				if self.maxRise != nil {
					index = self.maxRise.getHash().this
				} else {
					self.subRise = make(map[int]IText, len(s))
					self.maxRise = IText(s[i])
					self.maxRise.SetParent(
						self,
						Hash{this: 0, next: 0, prev: 0},
					)
				}
				self.subRise[index+1] = s[i]
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
	self.init()
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
