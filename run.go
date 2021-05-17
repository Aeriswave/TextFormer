package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var orco Mainer = "Самообучение"
	cwl := orco.Init(16) // запуск до 16 сопрограмм
	terminal, log := cwl.Create("LOG-Терминал", Terminal)

	synch := make(chan string)                                          // канал для сообщений синхронизации
	go Sync(time.Millisecond*1000, synch, nil, nil, nil, terminal.data) // запуск генератора синхроимпульсов

	tch1 := make(chan string) // подканал 1 для вывода текста основной программы
	tch2 := make(chan string) // подканал 2 для вывода текста антипарсера
	//	runtime.GOMAXPROCS(4)
	go ChannelMixer(synch, terminal.data, tch1, tch2) // смешивание нескольких текстовых каналов в один

	terminal.data <- "Запущена программа: " + string(orco)
	terminal.data <- log

	go aparser(tch1)

	var oo string = ""
	tch2 <- "Введите любой текст: "
	fmt.Fscan(os.Stdin, &oo)
	tch2 <- "сообщение из основной программы = " + oo
	tch2 <- "Выход\n"
	fmt.Scanf("...")
	fmt.Println("Программа остановлена, окно автоматически закроется через 10 секунд")
	time.Sleep(time.Second * 10)
}

// Запустить новый сопроцесс
func (self CWList) Create(name string, f func(CWCH)) (icw CWCH, log string) {
	self.index = (self.index + 1) % (self.coworks)
	_, have := self.coworksList[name]
	if have == false {
		_, have := self.cowoIDsList[self.index]
		if have == false {
			self.coworksList[name] = CW(self.index)   // выделение памяти под сопроцесс
			self.cowoIDsList[self.index] = name       // добавление индекса сопроцесса в список созданных сопроцессов
			icw, log = self.coworksList[name].Init(f) // инициализация и запуск сопроцесса
			return icw, "Новый процесс успешно создан, инициализирован и запущен\n"
		} else {
			return icw, "Ошибка: Невозможно создать процесс\n"
		}
	} else {
		log = "Ошибка: Невозможно создать новый сопроцесс с таким именем, т.к. сопроцесс с именем '" + name + "' уже существует.\n"
		return icw, log
	}
}

// Остановить сопроцесс
func (self CWList) Destroy(name string) (icw ICW, err string) {
	return nil, err
}

func (self Mainer) Init(coworks uint) ICWList {
	var cwl CWList = CWList{
		index:       -1,
		coworks:     int(1 + coworks),
		cowoIDsList: make(map[int]string, coworks),
		coworksList: make(map[string]CW, coworks)}

	return cwl
}

func (self CW) Init(f func(CWCH)) (chs CWCH, log string) {
	chs = self.start(f)
	chs.data <- "Cопрограмма инициализирована и запущена в горутине"
	return chs, log
}

// Создание каналов для работы с сопрограммой
// Запуск сопрограммы в горутине
// Функция возвращает набор каналов для связи с горутиной сопрограммы
func (self CW) start(f func(CWCH)) (chs CWCH) {
	chs = CWCH{
		logOUT: make(chan string),
		//		setIN:   make(chan string),
		//		setOUT:  make(chan string),
		//		runIN:   make(chan string),
		//		runOUT:  make(chan string),
		//		doIN:    make(chan string),
		//		doOUT:   make(chan string),
		data: make(chan string),
	}
	go f(chs)
	return chs
}
