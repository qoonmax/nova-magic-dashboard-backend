package generate_service

import (
	"bufio"
	"fmt"
	"magic-dashboard-backend/internal/http/requests"
	"magic-dashboard-backend/internal/services/utils"
	"os"
	"strings"
)

func generateResources(req *requests.GenerateRequest) ([]Instruction, error) {
	var instructions []Instruction

	tables := req.Database.Tables

	for _, table := range tables {
		modelName := utils.GetModelName(table.Name)

		file, err := os.OpenFile("internal/services/generate_service/stubs/resource.stub", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("error opening resource.stub:", err)
			return nil, err
		}

		var resourceContent string
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			line = strings.Replace(line, "{{Model}}", modelName, -1)
			resourceContent += line + "\n"
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("error reading resource.stub:", err)
			return nil, err
		}

		instructions = append(instructions, newInstruction(
			fmt.Sprintf("app/Nova/%s.php", modelName),
			resourceContent,
		))

	}
	return instructions, nil
}
