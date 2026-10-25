package actioninfo

import (
	"fmt"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Шаг 2: Распарсить значение с помощью метода Parse
		err := dp.Parse(data)
		// Шаг 3: Обработать ошибку парсинга
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		// Шаг 4: Сформировать строку с информацией об активности
		infoString, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		// Вывести строку на экран
		fmt.Println(infoString)
	}
}