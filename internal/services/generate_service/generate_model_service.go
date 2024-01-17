package generate_service

import (
	"bufio"
	"fmt"
	"magic-dashboard-backend/internal/http/requests"
	"magic-dashboard-backend/internal/services/utils"
	"os"
	"strings"
)

func generateModels(req *requests.GenerateRequest) ([]Instruction, error) {
	var instructions []Instruction

	tables := req.Database.Tables

	for _, table := range tables {
		modelName := getModelName(table.Name)

		file, err := os.OpenFile("internal/services/generate_service/stubs/model.stub", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("error opening model.stub:", err)
			return nil, err
		}

		var updatedContent string
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			line = strings.Replace(line, "{{Model}}", modelName, -1)
			updatedContent += line + "\n"
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("error reading model.stub:", err)
			return nil, err
		}

		instructions = append(instructions, newInstruction(
			fmt.Sprintf("app/Models/%s.php", modelName),
			updatedContent,
		))

	}
	return instructions, nil
}

func getModelName(table string) string {
	modelName := utils.GetSingularName(table)
	modelName = strings.ToUpper(modelName[:1]) + modelName[1:]
	return modelName
}
