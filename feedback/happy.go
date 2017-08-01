package feedback

type HappyCommand struct {
	Description string `short:"d" long:"description" description:"Write your message" required:"true"`
}

func (h *HappyCommand) Execute(args []string) error {
	r := RetroItem{
		Description: h.Description,
		Category:    CategoryHappy,
	}
	return Run(r)
}
