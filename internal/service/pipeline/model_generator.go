package pipeline

import (
	"bufio"
	"fmt"
	"magic-dashboard-backend/internal/service/utils"
	"os"
	"strings"
)

func GenerateModels(pctx *Context) error {
	tables := pctx.Req.Database.Tables

	for _, table := range tables {
		modelName := utils.GetSingularName(table.Name)
		modelName = strings.ToUpper(modelName[:1]) + modelName[1:]

		file, err := os.OpenFile("internal/service/pipeline/stubs/model.stub", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("error opening model.stub:", err)
			return err
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
			return err
		}

		pctx.AddInstruction(Instruction{
			Path:    fmt.Sprintf("app/Models/%s.php", modelName),
			Content: updatedContent,
		})
	}
	return nil
}
