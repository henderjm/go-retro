package feedback

type FeedbackCommand struct {
	RetroId    string            `short:"r" long:"retro-id" description:"Retro Board Id" required:"true"`
	Login      LoginCommand      `command:"login" alias:"l" description:"Login in to retro board"`
	Meh        MehCommand        `command:"meh" alias:"m" description:"Raise a potential concern"`
	Happy      HappyCommand      `command:"happy" alias:"h" description:"Express your happiness"`
	Sad        SadCommand        `command:"sad" alias:"s" description:"Why so sad??"`
	Actions    ActionsCommand    `command:"actions" alias:"a" description:"See all actions"`
	StartRetro StartRetroCommand `command:"start-retro" alias:"sr" description:"Let's start retro-ing"`
}

type Category string

const (
	CategoryMeh   Category = "meh"
	CategoryHappy Category = "happy"
	CategorySad   Category = "sad"
)

type RetroBoard struct {
	Board struct {
		Slug        string `json:"slug"`
		ActionItems []struct {
			Description string `json:"description"`
			ID          uint64 `json:"id"`
			Done        bool   `json:"done"`
		} `json:"action_items"`
		RetroItems []RetroItem `json:"items"`
	} `json:"retro"`
}

func (i *RetroItem) MarkItemAsDone() error {
	i.Done = true
	return nil
}

type RetroItem struct {
	Description string   `json:"description"`
	Category    Category `json:"category"`
	Done        bool     `json:"done,omitempty"`
	ID          int64    `json:"id"`
}

var FeedBack FeedbackCommand
