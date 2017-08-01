package feedback

type FeedbackCommand struct {
	RetroId string       `short:"r" long:"retro-id" description:"Retro Board Id" required:"true"`
	Meh     MehCommand   `command:"meh" alias:"m" description:"Raise a potential concern"`
	Happy   HappyCommand `command:"happy" alias:"h" description:"Express your happiness"`
	Sad     SadCommand   `command:"sad" alias:"s" description:"Why so sad??"`
}

type Category string

const (
	CategoryMeh   Category = "meh"
	CategoryHappy Category = "happy"
	CategorySad   Category = "sad"
)

type RetroItem struct {
	Description string   `json:"description"`
	Category    Category `json:"category"`
}

var FeedBack FeedbackCommand
