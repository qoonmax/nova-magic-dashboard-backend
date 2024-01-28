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
		modelName := utils.GetModelName(table.Name)

		file, err := os.OpenFile("internal/services/generate_service/stubs/model.stub", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("error opening model.stub:", err)
			return nil, err
		}

		var modelContent string
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			line = strings.Replace(line, "{{Model}}", modelName, -1)
			line = strings.Replace(line, "{{Table}}", table.Name, -1)
			modelContent += line + "\n"
		}

		modelContent += "\n" + fmt.Sprintf("    %s\n", getModelRelations(req))
		modelContent += "}"

		if err := scanner.Err(); err != nil {
			fmt.Println("error reading model.stub:", err)
			return nil, err
		}

		instructions = append(instructions, newInstruction(
			fmt.Sprintf("app/Models/%s.php", modelName),
			modelContent,
		))

	}
	return instructions, nil
}

func getModelRelations(req *requests.GenerateRequest) string {

	return "public function quiz (): \\Illuminate\\Database\\Eloquent\\Relations\\BelongsTo\n    {\n        return $this->belongsTo(Quiz::class);\n    }"
}
