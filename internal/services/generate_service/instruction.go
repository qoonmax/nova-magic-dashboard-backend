package generate_service

type Instruction struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func newInstruction(path, content string) Instruction {
	return Instruction{
		Path:    path,
		Content: content,
	}
}
