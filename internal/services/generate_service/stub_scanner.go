package generate_service

import (
	"bufio"
	"fmt"
	"os"
)

func getStubScanner(fileName string) (*bufio.Scanner, error) {
	// Откроем файл с шаблоном
	file, err := os.OpenFile("internal/services/generate_service/stubs/"+fileName+".stub", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("error opening "+fileName+".stub:", err)
		return nil, err
	}

	// Закроем файл после завершения работы
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing "+fileName+".stub:", err)
		}
	}(file)

	// Создадим сканер для чтения файла
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading "+fileName+".stub:", err)
		return nil, err
	}

	return scanner, nil
}
