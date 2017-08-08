package feedback_test

import (
	"os/exec"

	"encoding/json"

	"github.com/henderjm/go-feedback/feedback"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Start", func() {

	var retro string = `
	{
		"retro": {
			"id": 330,
			"slug": "330",
			"name": "Services Pipeline Team",
			"items": [{
				"id": 144146,
				"description": "Hello World",
				"category": "happy",
				"vote_count": 0,
				"done": false,
				"created_at": "2017-08-01T21:41:00.536Z",
				"archived_at": null
			}, {
				"id": 144147,
				"description": "Hello World",
				"category": "sad",
				"vote_count": 0,
				"done": false,
				"created_at": "2017-08-01T21:41:53.344Z",
				"archived_at": null
			}, {
				"id": 144148,
				"description": "Hello World",
				"category": "meh",
				"vote_count": 0,
				"done": false,
				"created_at": "2017-08-01T21:41:57.539Z",
				"archived_at": null
			}],
			"action_items": [{
				"id": 23389,
				"description": "[Mark] Tell teams and mention at anchors meeting to move off zumba!",
				"done": true,
				"created_at": "2017-06-02T15:27:36.041Z",
				"archived_at": null
			}, {
				"id": 24104,
				"description": "[jMe/David] Setup an email list so that we can blast for new releases",
				"done": false,
				"created_at": "2017-06-08T11:16:16.939Z",
				"archived_at": null
			}]
		}
	}
		`
	Context("when it is executable", func() {
		var cmd *exec.Cmd
		var actions feedback.RetroBoard

		BeforeSuite(func() {
			cmd = exec.Command("go", "run main.go", "start-retro")
			json.Unmarshal([]byte(retro), &actions)
		})

		It("should execute with no errors", func() {
			_, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should parse the json correctly", func() {
			items := actions.Board.RetroItems
			Expect(len(items)).To(Equal(3))
			Expect(items[0].Description).To(Equal("Hello World"))
			Expect(items[0].Category).To(Equal(feedback.CategoryHappy))
			Expect(items[0].Done).To(BeFalse())
		})

		It("should change Done flag to true when marked done", func() {
			item := actions.Board.RetroItems[0]
			Expect(item.Done).To(BeFalse())
			item.MarkItemAsDone()
			Expect(item.Done).To(BeTrue())
		})
	})
})
