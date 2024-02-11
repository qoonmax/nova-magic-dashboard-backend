package generate_service

import (
	"fmt"
	"magic-dashboard-backend/internal/http/requests"
	"magic-dashboard-backend/internal/services/utils"
	"os"
	"strings"
)

func generateModels(req *requests.GenerateRequest) ([]Instruction, error) {
	var instructions []Instruction

	// Получим таблицы из запроса
	tables := req.Database.Tables

	// Откроем файл с шаблоном модели
	file, err := os.OpenFile("internal/services/generate_service/stubs/model.stub", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("error opening model.stub:", err)
		return nil, err
	}

	// Закроем файл после завершения работы
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing model.stub:", err)
		}
	}(file)

	// Для каждой таблицы создадим модель
	for _, table := range tables {
		// Получим имя будущей модели, из названия таблицы
		modelName := utils.GetModelName(table.Name)

		// Создадим строку для записи в файл
		var modelContent string

		// Создадим сканер для чтения файла
		scanner, err := getStubScanner("model")
		if err != nil {
			return nil, err
		}

		// Прочитаем файл и заменим в нем все переменные на нужные значения
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.Replace(line, "{{Model}}", modelName, -1)
			line = strings.Replace(line, "{{Table}}", table.Name, -1)
			modelContent += line + "\n"
		}

		// Добавим отношения в модель
		if modelRelations, err := getModelRelations(table); err != nil {
			return nil, err
		} else {
			modelContent += "\n" + fmt.Sprintf("    %s\n", modelRelations)
			modelContent += "}"
		}

		// Добавим модель в 'инструкции' для создания файлов
		instructions = append(instructions, newInstruction(
			fmt.Sprintf("app/Models/%s.php", modelName),
			modelContent,
		))

	}
	return instructions, nil
}

func getModelRelations(table requests.Table) (string, error) {
	var modelRelations string

	// Найдем колонки с внешними ключами
	for _, column := range table.Columns {
		//TODO: проверить, что таблица является ли таблица промежуточной связующей таблицей

		// Если таблица является обычной таблицей, и колонка является внешним ключом (BelongsTo)
		if column.IsForeignKey == true {
			if belongsToRelation, err := getBelongsToRelation(column); err == nil {
				modelRelations += belongsToRelation
			} else {
				return "", err
			}
		}
	}

	return modelRelations, nil
}

func getBelongsToRelation(column requests.Column) (string, error) {
	var belongsToRelation string

	referencedTableName := column.ReferencedTableName
	referencedModelName := utils.GetModelName(referencedTableName)

	scanner, err := getStubScanner("model_relation")
	if err != nil {
		return "", err
	}

	// Прочитаем файл и заменим в нем все переменные на нужные значения
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "{{Entity}}", strings.ToLower(referencedModelName), -1)
		line = strings.Replace(line, "{{ReturnedType}}", "\\Illuminate\\Database\\Eloquent\\Relations\\BelongsTo", -1)
		//TODO: id заменить на переменную, в роли внешнего ownerKey может быть не только id
		line = strings.Replace(line, "{{RelationMethod}}", "belongsTo(\\App\\Models\\"+referencedModelName+"::class, "+column.Name+", id)", -1)
		belongsToRelation += line + "\n"
	}

	return belongsToRelation, nil
}

func getHasOneRelation() string {
	return ""
}

func getHasManyRelation() string {
	return ""
}
