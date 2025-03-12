package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(datastring string) (err error) // Метод парсит строку с данными
	ActionInfo() (string, error)         // получить строку действия и ошибку
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {

	for _, v := range dataset {

		err := dp.Parse(v)

		if err != nil {
			fmt.Println(err)
			continue
		}

		str, err := dp.ActionInfo()

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(str)

	}
}
