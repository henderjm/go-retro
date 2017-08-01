package feedback

type SadCommand struct {
	Description string `short:"d" long:"description" description:"Write your message" required:"true"`
}

func (s *SadCommand) Execute(args []string) error {
	r := RetroItem{
		Description: s.Description,
		Category:    CategorySad,
	}

	return Run(r)
}
