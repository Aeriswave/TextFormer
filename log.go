package main

import (
	"fmt"
	"time"
)

//	Генератор синхроимпульсов
func Sync(freq time.Duration, syncOUT chan<- string, syncIN <-chan string, controlIN <-chan string, controlOUT chan<- string, log chan<- string) {
	log <- "Запущен генератор синхроимпульсов"
	var startTime time.Time = time.Now()
	var localTiming time.Time = time.Now()
	msg := "пустая база для принимаемого сообщения"
	lt := time.Since(localTiming)
	timeout := time.After(freq)
	for {
		select { // Оператор select для чтения данных из канала
		case <-timeout: // Ждет окончания такта генератора
			syncOUT <- "Тик. " + "Раунд " + fmt.Sprint(time.Since(startTime)/freq)
			lt = time.Since(localTiming)
			localTiming = time.Now()
			timeout = time.After(freq - (lt % freq)) // откорректированный на время с уже прошедшее с предыдущего тика		case msg := <-syncIN: // Ждет, когда проснется гофер синхронизации
		case msg = <-controlIN: // ожидание команды управления генератором
			log <- msg        // вывести в log полученную команду
			controlOUT <- msg // отправить ответное сообщение о приеме командны для обратной связи
			return            // остановить генератор
		}
	}
}

// Смешивание текста из нескольких каналов в один общий
// Смешивание выполняется раз в определенное время, строго после получения сопрограммой
// сообщения в потоке синхронизации о начале нового раунда
func ChannelMixer(sync <-chan string, a chan<- string, b ...<-chan string) {
	a <- "Микшер текста запущен"
	syncMSG := "базовое сообщение синхронизации"
	msg := "базовое сообщение терминала"
	for {
		select { // Оператор select для чтения данных из канала
		case syncMSG = <-sync: // Ждет, когда проснется гофер синхронизации
			//			if syncMSG == "Тик" { // обработка по типам сообщений
			for i := range b {
				go func(i int) {
					msg = <-b[i]
					a <- "mixer: synMSG = " + syncMSG + ", " + "гофер = " + fmt.Sprint(i) + "\n\t\tmsg= \n=>>>\n" + msg + "\n<<<="
				}(i)
			}
		}
	}
}

// Вывод текста из потока в консоль терминала
func Terminal(chs CWCH) {
	fmt.Println("Терминал запущен")
	for {
		select { // Оператор select для чтения данных из канала
		case msg := <-chs.data: // Ждет, когда проснется гофер
			fmt.Println("Terminal: \n\t" + msg)
		}
	}
}
