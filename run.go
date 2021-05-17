package main

import (
	"fmt"
	"os"
	"time"
)

//	"fmt"

func main() {
	var orco Mainer = "Самообучение"
	cwl := orco.Init(16) // запуск до 16 сопрограмм
	fmt.Println("Запущена программа: " + orco)

	icw, log := cwl.Create("LOG-Терминал")
	synch := make(chan string) // канал для сообщений синхронизации
	logch := make(chan string) // канал для вывода текста в терминал
	tch1 := make(chan string)  // подканал 1 для вывода текста
	log += "ds"
	tch2 := make(chan string) // подканал 2 для вывода текста
	//	runtime.GOMAXPROCS(4)
	go Sync(time.Millisecond*1000, synch, nil, nil, nil, logch) // запуск гофера с генератором синхроимпульсов
	go ChannelMixer(synch, logch, tch1, tch2)                   // смешивание нескольких текстовых каналов в один
	go Terminal(logch)
	go aparser(tch1)

	icw.Init()
	//	fmt.Scanf("...")
	//	fmt.Println("_+_+_+_")

	var oo string = ""
	time.Sleep(time.Second * 3)
	tch2 <- "Введите любой текст: "
	fmt.Fscan(os.Stdin, &oo)
	tch2 <- "сообщение из основной программы = " + oo
	tch2 <- "Выход\n"
	fmt.Scanf("...")
	fmt.Println("Программа остановлена")
}

// Запустить новый сопроцесс
func (self CWList) Create(name string) (icw ICW, log string) {
	self.index = (self.index + 1) % (self.coworks)
	_, have := self.coworksList[name]
	if have == false {
		_, have := self.cowoIDsList[self.index]
		if have == false {
			self.coworksList[name] = CW(self.index)
			self.cowoIDsList[self.index] = name
			icw = self.coworksList[name]
			return icw, "Новый процесс успешно создан\n"
		} else {
			return nil, "Ошибка: Невозможно создать процесс\n"
		}
	} else {
		log = "Ошибка: Невозможно создать новый сопроцесс с таким именем, т.к. сопроцесс с именем '" + name + "' уже существует.\n"
		return nil, log
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

func (self CW) Init() (iset ISetter, irun IRunner, idrive IDriver, log string) {
	// iset.Off()
	// irun.Wait()
	//settings:=
	return iset, irun, idrive, log
}
