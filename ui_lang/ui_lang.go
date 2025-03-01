package ui_lang


import (
	s "strings"
	// exe "github.com/mcuadros/OctoPrint-TFT/ui_exec"
)


var	lang map[string]string = map[string]string{}
var	langs [countLangs]string
var	CurrentLang string

const countLangs = 4

func init(){
	langs[0] = "en"
	langs[1] = "kz"
	langs[2] = "qz"
	langs[3] = "ru"
}

func FindAndTranslate(needle string, text string) string{
	return s.Replace(text, needle, Translate(needle), -1)
}

func GetLanguagesList() [countLangs]string{
	return langs
}

func Translate(word string) string{
 var lng = lang_en()
	switch CurrentLang {
	case "ru":
		lng = lang_ru()
	case "en":
		lng = lang_en()
	case "kz":
		lng = lang_kz()
	case "qz":
		lng = lang_qz()
	}
	
	value, ok := lng[word]
	if ok {
		return value
	} 
	return lng[word]
}

func lang_ru() map[string]string{
	var l map[string]string = map[string]string{}
  l["en"] = "English"
  l["ru"] = "Русский"
  l["kz"] = "Қазақ тілі"
  l["qz"] = "Qazaq tіlі"
  l["Back"] = "Назад"
  l["New background task started"] = "Запущено фоновое задание"
  l["Background task closed"] = "Фоновое задание завершено"
  l["IdleAdd() failed:"] = "IdleAdd() ошибка:"
  l["Motor Off"] = "Выкл мотор"
  l["Fan On"] = "Вкл обдув"
  l["Fan Off"] = "Выкл обдув"
  l["Retrieving custom controls"] = "Получение пользовательских элементов управления"
  l["Executing command %q"] = "Запуск команды %q"
  l["Retrieving custom commands"] = "Получение пользователских команд"
  l["Extrude"] = "Загрузить"
  l["Retract"] = "Выгрузить"
  l["5mm"] = "5мм"
  l["10mm"] = "10мм"
  l["50mm"] = "50мм"
  l["1mm"] = "1мм"
  l["New tool detected %s"] = "Найден новый эксртудер %s"
  l["Changing tool to %s"] = "Экструдер заменен на %s"
  l["Normal"] = "Норм"
  l["High"] = "Быстро"
  l["Slow"] = "Медленно"
  l["Changing flowrate to %d%%"] = "Изменена скорость выдавливания на %d%%"
  l["Sending extrude request, with amount %d"] = "Отправлен запрос на выдавливание, со скоростью %d"
  l["Refreshing list of files"] = "Обновление списка файлов"
  l["<small>Uploaded: <b>%s</b> - Size: <b>%s</b></small>"] = "<small>Загружен: <b>%s</b> - Размер: <b>%s</b></small>"
  l["Are you sure you want to proceed?"] = "Вы действительно хотите продолжить?"
  l["Loading file %q, printing: %v"] = "Загрузка файла %q, печать: %v"
  l["Home All"] = "Полная парковка"
  l["Home X"] = "Парковка X"
  l["Home Y"] = "Парковка Y"
  l["Home Z"] = "Парковка Z"
  l["Homing the print head in %s axes"] = "Парковка экструдера по оси %s"
  l["X+"] = "X+"
  l["Y+"] = "Y+"
  l["Z+"] = "Z+"
  l["X-"] = "X-"
  l["Y-"] = "Y-"
  l["Z-"] = "Z-"
  l["Jogging print head axis %s in %dmm"] = "Движение экструдера по оси %s на %dmm"
  l["Connecting to OctoPrint..."] = "Соединяемся с OctoPrint..."
  l["Starting a new job"] = "Запуск печати..."
  l["Print"] = "Печать"
  l["Pause"] = "Пауза"
  l["Pausing/Resuming job"] = "Пауза/Продолжение работы"
  l["Stop"] = "Стоп"
  l["Stopping job"] = "Остановка печати"
  l["bed"] = "стол"
  l["tool0"] = "экструдер"
  l["tool1"] = "экструдер2"
  l["Resume"] = "Продолжить"
  l["File: %s"] = "Файл: %s"
  l["Printer is ready"] = "Принтер готов"
  l["Job Completed in %s"] = "Печать завершена за %s"
  l["Elapsed/Left: %s / %s"] = "Прошло/Осталось: %s / %s"
  l["Elapsed: %s"] = "Прошло: %s"
  l["Bed"] = "Стол"
  l["Tool0"] = "Экструдер"
  l["Tool1"] = "Экструдер2"
  l["Warming up ..."] = "Разогрев ..."
  l["<i>not-set</i>"] = "<i>не выбран</i>"
  l["Detecting baudrate"] = "Определение скорости порта"
  l["<b>Versions Information</b>"] = "<b>Версии ПО</b>"
  l["<b>OctoPrint-TFT Version</b>"] = "<b>Версия прошивки</b>"
  l["OctoPi Version: <b>%s</b>"] = "Версия OctoPi: <b>%s</b>"
  l["OctoPrint Version: <b>%s (%s)</b>"] = "Версия OctoPrint: <b>%s (%s)</b>"
  l["<b>System Information</b>"] = "<b>Системная информация</b>"
  l["Memory Total / Free: <b>%s / %s</b>"] = "Память Всего / Свободно: <b>%s / %s</b>"
  l["Load Average: <b>%.2f, %.2f, %.2f</b>"] = "Нагрузка: <b>%.2f, %.2f, %.2f</b>"
  l["Increase"] = "Прибавить"
  l["Decrease"] = "Уменьшить"
  l["Profiles"] = "Преднастроки"
  l["Setting target temperature for %s to %1.f°C."] = "Изменение температуры %s до %1.f°C."
  l["unable to find tool %q"] = "не удается найти %q"
  l["Cool Down"] = "Охлаждение"
  l["Setting temperature profile %s."] = "Выбран температурный режим %s."
  l["minutes ago"] = "минут назад"
  l["minute ago"] = "минуту назад"
  l["hours ago"] = "часов назад"
  l["hour ago"] = "час назад"
  l["hours from now"] = "часов назад"
  l["days ago"] = "дней назад"
  l["day ago"] = "день назад"
  l["week ago"] = "неделю назад"
  l["weeks ago"] = "недели назад"
  l["month ago"] = "месяц назад"
  l["months ago"] = "месяцев назад"
  l["year ago"] = "год назад"
  l["years ago"] = "лет назад"
  l["a long while ago"] = "давным давно"
  l["ago"] = "назад"
  l["kB"] = "Кбайт"
  l["mB"] = "Мбайт"
  l["B"] = "Байт"
  l["Language"] = "Язык"
  l["Error getting GDK screen: %s"] = "Ошибка отображения: %s"
  l["Unexpected error: %s"] = "Неизвестная ошибка: %s"
  l["Printing a job"] = "Идет печать..."
  l["Connection offline, connecting: %s"] = "Нет соединения, подключение: %s"
  l["Error connecting to printer: %s"] = "Ошибка подключения к принтеру: %s"
  l["Waiting for connection: %s"] = "Ожидание соединения: %s"
  l["Unable to connect to %q (Key: %v), \nmaybe OctoPrint not running?"] = "Не удается подключиться к %q (Key: %v), \nможет быть OctoPrint не запущен?"
  l["Status"] = "Статус"
  l["Heat Up"] = "Нагрев"
  l["Move"] = "Движение"
  l["Home"] = "Парковка"
  l["Settings"] = "Настройки"
  l["Filament"] = "Пластик"
  l["Control"] = "Управление"
  l["Files"] = "Файлы"
  l["System"] = "Система"
  l["Printer is not operational"] = "Принтер не исправен"
  l["Rebooting MKS..."] = "Перезапуск платы управления..."
  l["Save"] = "Сохранить"
  l["Network"] = "Сеть"
  l["Selected language: %s"] = "Выбран язык: %s"
  l["WIFI SSID: %s"] = "Имя WIFI точки: %s"
  l["WIFI password: %s"] = "Пароль WIFI: %s"
  l["Load..."] = "Загрузить..."
  l["File not found or wrong"] = "Файл не найден или имеет неверный формат"
  l["<b>CURRENT PARAMS:</b>"] = "<b>ТЕКУЩИЕ ПАРАМЕТРЫ:</b>"
  l["<b>LOADED PARAMS:</b>"] = "<b>ЗАГРУЖЕННЫЕ ПАРАМЕТРЫ:</b>"
  l["WIFI no pass"] = "Не задан пароль точки доступа"
  l["New settings saved"] = "Настройки сохранены"
  l["Settings not saved"] = "Ошибка сохранения настроек"
  l["<b>IP LAN: %s</b>"] = "IP адрес LAN: %s"
  l["<b>IP wifi: %s</b>"] = "IP адрес WIFI: %s"
  l["UP DIR"] = "НАЗАД"
  l["Go to prev directory"] = "Перейти на директорию выше"
  l["Calibrate"] = "Калибровать"
  l["IP address: %s"] = "IP адрес: %s"
  l["Gateway address: %s"] = "IP шлюза: %s"
  l["DNS address: %s"] = "DNS адрес: %s"
  l["IP address settings not found"] = "Настройки IP адреса не найдены"
  l["Gateway IP  not found"] = "IP адрес шлюза не найден"
  l["DNS IP not found"] = "DNS адрес не найден"
  l["IP address settings saved"] = "IP адреса сохранены"
  l["IP address settings not saved"] = "Настройки IP адреса не сохранены"
  l["LED On"] = "Вкл свет"
  l["LED Off"] = "Выкл свет"
  l["Faster"] = "Быстрее"
  l["Slower"] = "Медленнее"
  l["Speed: %d%%"] = "Скорость: %d%%"
  l["<b>InterPrint PiSoft LCD version:</b>"] = "<b>InterPrint PiSoft LCD версия:</b>"
  l["Fan +"] = "Обдув +"
  l["Fan -"] = "Обдув -"
  l["Refresh"] = "Обновить"
  l["Gateway IP LAN not found"] = "Настройки IP адреса LAN не найдены"
  l["LAN IP address: %s"] = "LAN IP адрес: %s"
  l["LAN Gateway address: %s"] = "LAN IP адрес шлюза: %s"
  l["LAN DNS address: %s"] = "LAN DNS адрес: %s"
  l["Update Software"] = "Обновить ПО"
  l["Current version is latest"] = "Установалена актуальная версия ПО"
  l["Update error: md5 sum incorrect"] = "Ошибка обновления: md5 сумма не верна"
  l["Update error"] = "Ошибка обновления"
  l["Reboot"] = "Перезагрузить"
  l["Updates will be instaled within 1 minute"] = "Обновления будут установлены в течении 1 минуты"
  l["Current temperature below 150°C"] = "Невозможно выполнить действия, температура экструдера ниже 150°C"
  l["Error: printer is busy"] = "Не возможно выполнить действие: идет процесс печати"
  l["Press down"] = "Прижать сопло"
  l["Push up"] = "Поднять сопло"
	
	return l
}

func lang_en() map[string]string{
	var l map[string]string = map[string]string{}
  l["en"] = "English"
  l["ru"] = "Русский"
  l["kz"] = "Қазақ тілі"
  l["qz"] = "Qazaq tіlі"
  l["Back"] = "Back"
  l["New background task started"] = "New background task started"
  l["Background task closed"] = "Background task closed"	
  l["IdleAdd() failed:"] = "IdleAdd() failed:"
  l["Motor Off"] = "Motor Off"
  l["Fan On"] = "Fan On"
  l["Fan Off"] = "Fan Off"
  l["Retrieving custom controls"] = "Retrieving custom controls"
  l["Executing command %q"] = "Executing command %q"
  l["Retrieving custom commands"] = "Retrieving custom commands"
  l["Extrude"] = "Extrude"
  l["Retract"] = "Retract"
  l["5mm"] = "5mm"
  l["10mm"] = "10mm"
  l["50mm"] = "50mm"
  l["1mm"] = "1mm"
  l["New tool detected %s"] = "New tool detected %s"
  l["Changing tool to %s"] = "Changing tool to %s"
  l["Normal"] = "Normal"
  l["High"] = "High"
  l["Slow"] = "Slow"
  l["Changing flowrate to %d%%"] = "Changing flowrate to %d%%"
  l["Sending extrude request, with amount %d"] = "Sending extrude request, with amount %d"
  l["Refreshing list of files"] = "Refreshing list of files"
  l["<small>Uploaded: <b>%s</b> - Size: <b>%s</b></small>"] = "<small>Uploaded: <b>%s</b> - Size: <b>%s</b></small>"
  l["Are you sure you want to proceed?"] = "Are you sure you want to proceed?"
  l["Loading file %q, printing: %v"] = "Loading file %q, printing: %v"
  l["Home All"] = "Home All"
  l["Home X"] = "Home X"
  l["Home Y"] = "Home Y"
  l["Home Z"] = "Home Z"
  l["Homing the print head in %s axes"] = "Homing the print head in %s axes"
  l["X+"] = "X+"
  l["Y+"] = "Y+"
  l["Z+"] = "Z+"
  l["X-"] = "X-"
  l["Y-"] = "Y-"
  l["Z-"] = "Z-"
  l["Jogging print head axis %s in %dmm"] = "Jogging print head axis %s in %dmm"
  l["Connecting to OctoPrint..."] = "Connecting to OctoPrint..."
	l["Starting a new job"] = "Starting a new job"
  l["Print"] = "Print"
  l["Pause"] = "Pause"
  l["Pausing/Resuming job"] = "Pausing/Resuming job"
  l["Stop"] = "Stop"
  l["Stopping job"] = "Stopping job"
  l["bed"] = "bed"
  l["tool0"] = "tool"
  l["tool1"] = "tool2"
  l["Resume"] = "Resume"
  l["File: %s"] = "File: %s"
  l["Printer is ready"] = "Printer is ready"
  l["Job Completed in %s"] = "Job Completed in %s"
  l["Elapsed/Left: %s / %s"] = "Elapsed/Left: %s / %s"
  l["Elapsed: %s"] = "Elapsed: %s"
  l["Bed"] = "Bed"
  l["Tool0"] = "Tool"
  l["Tool1"] = "Tool2"
  l["Warming up ..."] = "Warming up ..."
  l["<i>not-set</i>"] = "<i>not-set</i>"
  l["Detecting baudrate"] = "Detecting baudrate"
  l["<b>Versions Information</b>"] = "<b>Versions Information</b>"
  l["<b>OctoPrint-TFT Version</b>"] = "<b>Firmware Version</b>"
  l["OctoPi Version: <b>%s</b>"] = "OctoPi Version: <b>%s</b>"
  l["OctoPrint Version: <b>%s (%s)</b>"] = "OctoPrint Version: <b>%s (%s)</b>"
  l["<b>System Information</b>"] = "<b>System Information</b>"
  l["Memory Total / Free: <b>%s / %s</b>"] = "Memory Total / Free: <b>%s / %s</b>"
  l["Load Average: <b>%.2f, %.2f, %.2f</b>"] = "Load Average: <b>%.2f, %.2f, %.2f</b>"
	l["Increase"] = "Increase"
  l["Decrease"] = "Decrease"
  l["Profiles"] = "Profiles"
  l["Setting target temperature for %s to %1.f°C."] = "Setting target temperature for %s to %1.f°C."
  l["unable to find tool %q"] = "unable to find tool %q"
  l["Cool Down"] = "Cool Down"
  l["Setting temperature profile %s."] = "Setting temperature profile %s."
  l["minutes ago"] = "minutes ago"
  l["minute ago"] = "minute ago"
  l["hours ago"] = "hours ago"
  l["hour ago"] = "hour ago"
  l["hours from now"] = "hours from now"
  l["days ago"] = "days ago"
  l["day ago"] = "day ago"
  l["week ago"] = "week ago"
  l["weeks ago"] = "weeks ago"
  l["month ago"] = "month ago"
  l["months ago"] = "months ago"
  l["year ago"] = "year ago"
  l["years ago"] = "years ago"
  l["a long while ago"] = "a long while ago"
  l["kB"] = "kB"
  l["mB"] = "mB"
  l["B"] = "B"
  l["Language"] = "Language"
  l["Error getting GDK screen: %s"] = "Error getting GDK screen: %s"
  l["Unexpected error: %s"] = "Unexpected error: %s"
  l["Printing a job"] = "Printing a job"
  l["Connection offline, connecting: %s"] = "Connection offline, connecting: %s"
  l["Error connecting to printer: %s"] = "Error connecting to printer: %s"
  l["Waiting for connection: %s"] = "Waiting for connection: %s"
  l["Unable to connect to %q (Key: %v), \nmaybe OctoPrint not running?"] = "Unable to connect to %q (Key: %v), \nmaybe OctoPrint not running?"
  l["Status"] = "Status"
  l["Heat Up"] = "Heat Up"
  l["Move"] = "Move"
  l["Home"] = "Home"
  l["Settings"] = "Settings"
  l["Filament"] = "Filament"
  l["Control"] = "Control"
  l["Files"] = "Files"
  l["System"] = "System"
  l["Printer is not operational"] = "Printer is not operational"
  l["Rebooting MKS..."] = "Rebooting control board..."
  l["Save"] = "Save"
  l["Network"] = "Network"
  l["Selected language: %s"] = "Selected language: %s"
  l["WIFI SSID: %s"] = "SSID: %s"
  l["WIFI password: %s"] = "Password: %s"
  l["Load..."] = "Load..."
  l["File not found or wrong"] = "File not found or wrong"
	l["<b>CURRENT PARAMS:</b>"] = "<b>CURRENT PARAMS:</b>"
  l["<b>LOADED PARAMS:</b>"] = "<b>LOADED PARAMS:</b>"
  l["WIFI no pass"] = "WIFI no pass"
  l["New settings saved"] = "New settings saved"
  l["Settings not saved"] = "Settings not saved"
  l["<b>IP LAN: %s</b>"] = "<b>IP LAN: %s</b>"
  l["<b>IP wifi: %s</b>"] = "<b>IP WIFI: %s</b>"
  l["UP DIR"] = "UP DIR"
  l["Go to prev directory"] = "Go to prev directory"
  l["Calibrate"] = "Calibrate"
  l["IP address: %s"] = "IP address: %s"
  l["Gateway address: %s"] = "Gateway address: %s"
  l["DNS address: %s"] = "DNS address: %s"
  l["IP address settings not found"] = "IP address settings not found"
  l["Gateway IP  not found"] = "Gateway IP  not found"
  l["DNS IP not found"] = "DNS IP not found"
  l["IP address settings saved"] = "IP address settings saved"
  l["IP address settings not saved"] = "IP address settings not saved"
  l["LED On"] = "LED On"
  l["LED Off"] = "LED Off"
  l["Faster"] = "Faster"
  l["Slower"] = "Slower"
  l["Speed: %d%%"] = "Speed: %d%%"
  l["<b>InterPrint PiSoft LCD version:</b>"] = "<b>InterPrint PiSoft LCD version:</b>"
  l["Fan +"] = "Fan +"
  l["Fan -"] = "Fan -"
  l["Refresh"] = "Refresh"
  l["Gateway IP LAN not found"] = "Gateway IP LAN not found"
  l["DNS IP LAN not found"] = "DNS IP LAN not found"
  l["LAN IP address: %s"] = "LAN IP address: %s"
  l["LAN Gateway address: %s"] = "LAN Gateway address: %s"
  l["LAN DNS address: %s"] = "LAN DNS address: %s"
  l["Update Software"] = "Update Software"
  l["Current version is latest"] = "Current version is latest"
  l["Update error: md5 sum incorrect"] = "Update error: md5 sum incorrect"
  l["Update error"] = "Update error"
  l["Reboot"] = "Reboot"
  l["Updates will be instaled within 1 minute"] = "Updates will be instaled within 1 minute"
  l["Current temperature below 150°C"] = "Current temperature below 150°C"
  l["Error: printer is busy"] = "Error: printer is busy"
  l["Press down"] = "Press down"
  l["Push up"] = "Push up"
	return l
}

func lang_kz() map[string]string{
	var l map[string]string = map[string]string{}
  l["en"] = "English"
  l["ru"] = "Русский"
  l["kz"] = "Қазақ"
  l["qz"] = "Qazaq"
  l["Back"] = "Қайту"
  l["New background task started"] = "Фондық тапсырма қосылды"
  l["Background task closed"] = "Фондық тапсырма аяқталды"	
  l["IdleAdd() failed:"] = "IdleAdd() қате:"
  l["Motor Off"] = "Мотор сөнд."
  l["Fan On"] = "Салқын қосылды"
  l["Fan Off"] = "Салқын сөнді"
  l["Retrieving custom controls"] = "Қолданушының басқару элементін алу"
  l["Executing command %q"] = "%q командасын іске қосу"
  l["Retrieving custom commands"] = "Қолданушылардың командаларын алу"
  l["Extrude"] = "Енгізу"
  l["Retract"] = "Шығару"
  l["5mm"] = "5мм"
  l["10mm"] = "10мм"
  l["50mm"] = "50мм"
  l["1mm"] = "1мм"
  l["New tool detected %s"] = "Жана %s экструдер табылды"
  l["Changing tool to %s"] = "Экструдер %s аустырылды"
  l["Normal"] = "Кәдімгі"
  l["High"] = "Тез"
  l["Slow"] = "Баяу"
  l["Changing flowrate to %d%%"] = "Экструзия жылдамдығы %d%% өзгерді"
  l["Sending extrude request, with amount %d"] = "Экструзия командасы %d жылдамдығымен жіберілді"
  l["Refreshing list of files"] = "Файлдар тізімі жаңаруда"
  l["<small>Uploaded: <b>%s</b> - Size: <b>%s</b></small>"] = "<small><b>%s</b> енгізілді, көлемі: <b>%s</b></small>"
  l["Are you sure you want to proceed?"] = "Басып шығаруды растаныз"
  l["Loading file %q, printing: %v"] = "%q файл енгізілді, шығару: %v"
  l["Home All"] = "Толық тұрақ"
  l["Home X"] = "X тұрақ"
  l["Home Y"] = "Y тұрақ"
  l["Home Z"] = "Z тұрақ"
  l["Homing the print head in %s axes"] = "Экструдер %s осінің тұрағы"
  l["X+"] = "X+"
  l["Y+"] = "Y+"
  l["Z+"] = "Z+"
  l["X-"] = "X-"
  l["Y-"] = "Y-"
  l["Z-"] = "Z-"
  l["Jogging print head axis %s in %dmm"] = "Экструдер %s осімен %dмм қозғалуда"
  l["Connecting to OctoPrint..."] = "OctoPrint-ке қосылуда..."
	l["Starting a new job"] = "Басып шығару қосылуда..."
  l["Print"] = "Шығару"
  l["Pause"] = "Кідірту"
  l["Pausing/Resuming job"] = "Кідірту/Басып шығаруды жалғастыру"
  l["Stop"] = "Тоқтату"
  l["Stopping job"] = "Тоқтатылуда"
  l["bed"] = "алаң"
  l["tool0"] = "экструдер"
  l["tool1"] = "экструдер2"
  l["Resume"] = "Жалғастыру"
  l["File: %s"] = "Файл: %s"
  l["Printer is ready"] = "Принтер дайын"
  l["Job Completed in %s"] = "Басып шығару %s аяқталды"
  l["Elapsed/Left: %s / %s"] = "Өтті/Қалды: %s / %s"
  l["Elapsed: %s"] = "%s өтті"
  l["Bed"] = "Алаң"
  l["Tool0"] = "Экструдер"
  l["Tool1"] = "Экструдер2"
  l["Warming up ..."] = "Қыздыру ..."
  l["<i>not-set</i>"] = "<i>таңдалмаған</i>"
  l["Detecting baudrate"] = "Порт жылдамдығын анықтау"
  l["<b>Versions Information</b>"] = "<b>Бағдарлама нұсқасы</b>"
  l["<b>OctoPrint-TFT Version</b>"] = "<b>Прошивка нұсқасы</b>"
  l["OctoPi Version: <b>%s</b>"] = "OctoPi нұсқасы: <b>%s</b>"
  l["OctoPrint Version: <b>%s (%s)</b>"] = "OctoPrint нұсқасы: <b>%s (%s)</b>"
  l["<b>System Information</b>"] = "<b>Жүйе ақпараты</b>"
  l["Memory Total / Free: <b>%s / %s</b>"] = "Барлық жады / Бос жады: <b>%s / %s</b>"
  l["Load Average: <b>%.2f, %.2f, %.2f</b>"] = "CPU жүктемесі: <b>%.2f, %.2f, %.2f</b>"
	l["Increase"] = "Қосу"
  l["Decrease"] = "Алу"
  l["Profiles"] = "Үлгілер"
  l["Setting target temperature for %s to %1.f°C."] = "%s температурасын %1.f°C дейін өзгерту"
  l["unable to find tool %q"] = "%q табылмады"
  l["Cool Down"] = "Салқындату"
  l["Setting temperature profile %s."] = "%s температура үлгісі тандалды"
  l["minutes ago"] = "минут бұрын"
  l["minute ago"] = "минут бұрын"
  l["hours ago"] = "сағат бұрын"
  l["hour ago"] = "сағат бұрын"
  l["hours from now"] = "сағат бұрын"
  l["days ago"] = "күн бұрын"
  l["day ago"] = "күн бұрын"
  l["week ago"] = "апта бұрын"
  l["weeks ago"] = "апта бұрын"
  l["month ago"] = "ай бұрын"
  l["months ago"] = "ай бұрын"
  l["year ago"] = "жыл бұрын"
  l["years ago"] = "жыл бұрын"
  l["a long while ago"] = "көп уақыт бұрын"
  l["ago"] = "бұрын"
  l["kB"] = "Кбайт"
  l["mB"] = "Мбайт"
  l["B"] = "Байт"
  l["Language"] = "Тілі"
  l["Error getting GDK screen: %s"] = "%s көрсетілу қатесі"
  l["Unexpected error: %s"] = "Белгісіз қате: %s"
  l["Printing a job"] = "Басып шығаруда..."
  l["Connection offline, connecting: %s"] = "Сервер статусы: %s. Байланыс үзілді, байланыс қалпына келтіруде"
  l["Error connecting to printer: %s"] = "Принтерге косылуда қате: %s"
  l["Waiting for connection: %s"] = "Қосылу... Сервер статусы: %s"
  l["Unable to connect to %q (Key: %v), \nmaybe OctoPrint not running?"] = "%q (Key: %v) байланыс қосылған жоқ. \nOctoPrint қосылғанын тексеріңіз."
  l["Status"] = "Статус"
  l["Heat Up"] = "Қыздыру"
  l["Move"] = "Жылжыту"
  l["Home"] = "Тұрақ"
  l["Settings"] = "Параметрлер"
  l["Filament"] = "Пластик"
  l["Control"] = "Басқару"
  l["Files"] = "Файлдар"
  l["System"] = "Жүйе"
  l["Printer is not operational"] = "Принтер жұмыс істемейді"
  l["Rebooting MKS..."] = "Басқару платасы қайта іске қосылуда..."
  l["Save"] = "Сақтау"
  l["Network"] = "Желі"
  l["Selected language: %s"] = "%s тілі тандалды"
  l["WIFI SSID: %s"] = "WIFI желісінің аты: %s"
  l["WIFI password: %s"] = "WIFI паролі: %s"
  l["Load..."] = "Енгізу..."
  l["File not found or wrong"] = "Файл табылмады, немесе форматы белгісіз"
	l["<b>CURRENT PARAMS:</b>"] = "<b>Қазіргі параметрлер:</b>"
  l["<b>LOADED PARAMS:</b>"] = "<b>Параметрлер енгізу:</b>"
  l["WIFI no pass"] = "WIFI желесінің паролі табылмады"
  l["New settings saved"] = "Жаңа параметрлер сақталды"
  l["Settings not saved"] = "Параметрлер сақталған жоқ"
  l["<b>IP LAN: %s</b>"] = "<b>IP LAN: %s</b>"
  l["<b>IP wifi: %s</b>"] = "<b>IP WIFI: %s</b>"
  l["UP DIR"] = "АРТҚА ҚАЙТУ"
  l["Go to prev directory"] = "Алдыңғы каталогқа қайту"
  l["Calibrate"] = "Калибрлеу"
  l["IP address: %s"] = "IP адрес: %s"
  l["Gateway address: %s"] = "Шлюз адресі: %s"
  l["DNS address: %s"] = "DNS адресі: %s"
  l["IP address settings not found"] = "IP адрес параметрлері табылмады"
  l["Gateway IP  not found"] = "Шлюз адресі табылмады"
  l["DNS IP not found"] = "DNS адресі табылмады"
  l["IP address settings saved"] = "IP адрес параметрлері сақталды"
  l["IP address settings not saved"] = "IP адрес параметрлері сақталған жоқ"
  l["LED On"] = "Жарық қосу"
  l["LED Off"] = "Жарықсыз"
  l["Faster"] = "Тез"
  l["Slower"] = "Баяу"
  l["Speed: %d%%"] = "  Жылдамдық:\n        %d%%"
  l["<b>InterPrint PiSoft LCD version:</b>"] = "<b>InterPrint PiSoft LCD версия:</b>"
  l["Fan +"] = "Салқындық +"
  l["Fan -"] = "Салқындық -"
  l["Refresh"] = "Жаңарту"
  l["Gateway IP LAN not found"] = "Шлюз LAN адресі табылмады"
  l["DNS IP LAN not found"] = "DNS LAN адресі табылмады"
  l["LAN IP address: %s"] = "LAN IP адрес: %s"
  l["LAN Gateway address: %s"] = "LAN Шлюз адресі: %s"
  l["LAN DNS address: %s"] = "LAN DNS адресі: %s"
  l["Update Software"] = "Жүйені жаңарту"
  l["Current version is latest"] = "Current version is latest"
  l["Update error: md5 sum incorrect"] = "Update error: md5 sum incorrect"
  l["Update error"] = "Update error"
  l["Reboot"] = "Қайта жүктеу"
  l["Updates will be instaled within 1 minute"] = "Updates will be instaled within 1 minute"
  l["Current temperature below 150°C"] = "Current temperature below 150°C"
  l["Error: printer is busy"] = "Error: printer is busy"
  l["Press down"] = "Press down"
  l["Push up"] = "Push up"
	return l
}

func lang_qz() map[string]string{
	var l map[string]string = map[string]string{}
  l["en"] = "English"
  l["ru"] = "Русский"
  l["kz"] = "Қазақ"
  l["qz"] = "Qazaq"
  l["Back"] = "Qaitý"
  l["New background task started"] = "Fondyq tapsyrma qosyldy"
  l["Background task closed"] = "Fondyq tapsyrma aıaqtaldy"	
  l["IdleAdd() failed:"] = "IdleAdd() qate:"
  l["Motor Off"] = "Motor sónd."
  l["Fan On"] = "Salqyn qosyldy"
  l["Fan Off"] = "Salqyn sóndi"
  l["Retrieving custom controls"] = "Qoldanýshynyń basqarý elementin alý"
  l["Executing command %q"] = "%q komandasyn iske qosý"
  l["Retrieving custom commands"] = "Qoldanýshylardyń komandalaryn alý"
  l["Extrude"] = "Engizý"
  l["Retract"] = "Shyǵarý"
  l["5mm"] = "5mm"
  l["10mm"] = "10mm"
  l["50mm"] = "50mm"
  l["1mm"] = "1mm"
  l["New tool detected %s"] = "Jana %s ekstrýder tabyldy"
  l["Changing tool to %s"] = "Ekstrýder %s aýstyryldy"
  l["Normal"] = "Kádimgi"
  l["High"] = "Tez"
  l["Slow"] = "Baıaý"
  l["Changing flowrate to %d%%"] = "Ekstrýzıa jyldamdyǵy %d%% ózgerdi"
  l["Sending extrude request, with amount %d"] = "Ekstrýzıa komandasy %d jyldamdyǵymen jiberildi"
  l["Refreshing list of files"] = "Faıldar tizimi jańarýda"
  l["<small>Uploaded: <b>%s</b> - Size: <b>%s</b></small>"] = "<small><b>%s</b> engizildi, kólemi: <b>%s</b></small>"
  l["Are you sure you want to proceed?"] = "Basyp shyǵarýdy rastanyz"
  l["Loading file %q, printing: %v"] = "%q faıl engizildi, shyǵarý: %v"
  l["Home All"] = "Tolyq turaq"
  l["Home X"] = "X turaq"
  l["Home Y"] = "Y turaq"
  l["Home Z"] = "Z turaq"
  l["Homing the print head in %s axes"] = "Ekstrýder %s osiniń turaǵy"
  l["X+"] = "X+"
  l["Y+"] = "Y+"
  l["Z+"] = "Z+"
  l["X-"] = "X-"
  l["Y-"] = "Y-"
  l["Z-"] = "Z-"
  l["Jogging print head axis %s in %dmm"] = "Ekstrýder %s osimen %dmm qozǵalýda"
  l["Connecting to OctoPrint..."] = "OctoPrint-ke qosylýda..."
	l["Starting a new job"] = "Basyp shyǵarý qosylýda..."
  l["Print"] = "Shyǵarý"
  l["Pause"] = "Kidirtý"
  l["Pausing/Resuming job"] = "Kidirtý/Basyp shyǵarýdy jalǵastyrý"
  l["Stop"] = "Toqtatý"
  l["Stopping job"] = "Toqtatylýda"
  l["bed"] = "alań"
  l["tool0"] = "ekstrýder"
  l["tool1"] = "ekstrýder2"
  l["Resume"] = "Jalǵastyrý"
  l["File: %s"] = "Faıl: %s"
  l["Printer is ready"] = "Prınter daıyn"
  l["Job Completed in %s"] = "Basyp shyǵarý %s aıaqtaldy"
  l["Elapsed/Left: %s / %s"] = "Ótti/Qaldy: %s / %s"
  l["Elapsed: %s"] = "%s ótti"
  l["Bed"] = "Alań"
  l["Tool0"] = "Ekstrýder"
  l["Tool1"] = "Ekstrýder2"
  l["Warming up ..."] = "Qyzdyrý ..."
  l["<i>not-set</i>"] = "<i>tańdalmaǵan</i>"
  l["Detecting baudrate"] = "Port jyldamdyǵyn anyqtaý"
  l["<b>Versions Information</b>"] = "<b>Baǵdarlama nusqasy</b>"
  l["<b>OctoPrint-TFT Version</b>"] = "<b>Proshıvka nusqasy</b>"
  l["OctoPi Version: <b>%s</b>"] = "OctoPi nusqasy: <b>%s</b>"
  l["OctoPrint Version: <b>%s (%s)</b>"] = "OctoPrint nusqasy: <b>%s (%s)</b>"
  l["<b>System Information</b>"] = "<b>Júıe aqparaty</b>"
  l["Memory Total / Free: <b>%s / %s</b>"] = "Barlyq jady / Bos jady: <b>%s / %s</b>"
  l["Load Average: <b>%.2f, %.2f, %.2f</b>"] = "CPU júktemesi: <b>%.2f, %.2f, %.2f</b>"
	l["Increase"] = "Qosý"
  l["Decrease"] = "Alý"
  l["Profiles"] = "Úlgiler"
  l["Setting target temperature for %s to %1.f°C."] = "%s temperatýrasyn %1.f°C deıin ózgertý"
  l["unable to find tool %q"] = "%q tabylmady"
  l["Cool Down"] = "Salqyndatý"
  l["Setting temperature profile %s."] = "%s temperatýra úlgisi tandaldy"
  l["minutes ago"] = "mınýt buryn"
  l["minute ago"] = "mınýt buryn"
  l["hours ago"] = "saǵat buryn"
  l["hour ago"] = "saǵat buryn"
  l["hours from now"] = "saǵat buryn"
  l["days ago"] = "kún buryn"
  l["day ago"] = "kún buryn"
  l["week ago"] = "apta buryn"
  l["weeks ago"] = "apta buryn"
  l["month ago"] = "aı buryn"
  l["months ago"] = "aı buryn"
  l["year ago"] = "jyl buryn"
  l["years ago"] = "jyl buryn"
  l["a long while ago"] = "kóp ýaqyt buryn"
  l["ago"] = "buryn"
  l["kB"] = "Kbaıt"
  l["mB"] = "Mbaıt"
  l["B"] = "Baıt"
  l["Language"] = "Tili"
  l["Error getting GDK screen: %s"] = "%s kórsetilý qatesi"
  l["Unexpected error: %s"] = "Belgisiz qate: %s"
  l["Printing a job"] = "Basyp shyǵarýda..."
  l["Connection offline, connecting: %s"] = "Server statýsy: %s. Baılanys úzildi, baılanys qalpyna keltirýde"
  l["Error connecting to printer: %s"] = "Prınterge kosylýda qate: %s"
  l["Waiting for connection: %s"] = "Qosylý... Server statýsy: %s"
  l["Unable to connect to %q (Key: %v), \nmaybe OctoPrint not running?"] = "%q (Key: %v) baılanys qosylǵan joq. \nOctoPrint qosylǵanyn tekserińiz."
  l["Status"] = "Statýs"
  l["Heat Up"] = "Qyzdyrý"
  l["Move"] = "Jyljytý"
  l["Home"] = "Turaq"
  l["Settings"] = "Parametrler"
  l["Filament"] = "Plastık"
  l["Control"] = "Basqarý"
  l["Files"] = "Faıldar"
  l["System"] = "Júıe"
  l["Printer is not operational"] = "Prınter jumys istemeıdi"
  l["Rebooting MKS..."] = "Basqarý platasy qaıta iske qosylýda..."
  l["Save"] = "Saqtaý"
  l["Network"] = "Jeli"
  l["Selected language: %s"] = "%s tili tandaldy"
  l["WIFI SSID: %s"] = "WIFI jelisiniń aty: %s"
  l["WIFI password: %s"] = "WIFI paroli: %s"
  l["Load..."] = "Engizý..."
  l["File not found or wrong"] = "Faıl tabylmady, nemese formaty belgisiz"
	l["<b>CURRENT PARAMS:</b>"] = "<b>Qazirgi parametrler:</b>"
  l["<b>LOADED PARAMS:</b>"] = "<b>Parametrler engizý:</b>"
  l["WIFI no pass"] = "WIFI jelesiniń paroli tabylmady"
  l["New settings saved"] = "Jańa parametrler saqtaldy"
  l["Settings not saved"] = "Parametrler saqtalǵan joq"
  l["<b>IP LAN: %s</b>"] = "<b>IP LAN: %s</b>"
  l["<b>IP wifi: %s</b>"] = "<b>IP WIFI: %s</b>"
  l["UP DIR"] = "ARTQA QAITÝ"
  l["Go to prev directory"] = "Aldyńǵy katalogqa qaıtý"
  l["Calibrate"] = "Kalıbrleý"
  l["IP address: %s"] = "IP adres: %s"
  l["Gateway address: %s"] = "Shlúz adresi: %s"
  l["DNS address: %s"] = "DNS adresi: %s"
  l["IP address settings not found"] = "IP adres parametrleri tabylmady"
  l["Gateway IP  not found"] = "Shlúz adresi tabylmady"
  l["DNS IP not found"] = "DNS adresi tabylmady"
  l["IP address settings saved"] = "IP adres parametrleri saqtaldy"
  l["IP address settings not saved"] = "IP adres parametrleri saqtalǵan joq"
  l["LED On"] = "Jaryq qosý"
  l["LED Off"] = "Jaryqsyz"
  l["Faster"] = "Tez"
  l["Slower"] = "Baıaý"
  l["Speed: %d%%"] = "   Jyldamdyq:\n        %d%%"
  l["<b>InterPrint PiSoft LCD version:</b>"] = "<b>InterPrint PiSoft LCD versıa:</b>"
  l["Fan +"] = "Salqyndyq +"
  l["Fan -"] = "Salqyndyq -"
  l["Refresh"] = "Jańartý"
  l["Gateway IP LAN not found"] = "Shlúz LAN adresi tabylmady"
  l["DNS IP LAN not found"] = "DNS LAN adresi tabylmady"
  l["LAN IP address: %s"] = "LAN IP adres: %s"
  l["LAN Gateway address: %s"] = "LAN Shlúz adresi: %s"
  l["LAN DNS address: %s"] = "LAN DNS adresi: %s"
  l["Update Software"] = "Жүйені жаңарту"
  l["Current version is latest"] = "Current version is latest"
  l["Update error: md5 sum incorrect"] = "Update error: md5 sum incorrect"
  l["Update error"] = "Update error"
  l["Reboot"] = "Qaıta júkteý"
  l["Updates will be instaled within 1 minute"] = "Updates will be instaled within 1 minute"
  l["Current temperature below 150°C"] = "Current temperature below 150°C"
  l["Error: printer is busy"] = "Error: printer is busy"
  l["Press down"] = "Press down"
  l["Push up"] = "Push up"
	return l
}
