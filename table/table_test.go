package table_test

import (
	"io/ioutil"

	table "github.com/henderjm/go-feedback/table"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Table", func() {

	Context("Building simple output", func() {
		var data []string = []string{"this", "is", "a", "data", "column"}

		It("Should print a single column to stdout", func() {
			t := table.Table{
				Format: table.Simple,
			}

			r, err := t.Render(data)
			Expect(err).ToNot(HaveOccurred())
			Expect(r).To(Equal(readTable("fixtures/simple_column")))
		})

		It("Should print a single column to stdout with padding", func() {
			t := table.Table{
				Format: table.Padded,
			}

			r, err := t.Render(data)
			Expect(err).ToNot(HaveOccurred())
			Expect(r).To(Equal(readTable("fixtures/simple_column_with_padding")))
		})
	})
})

func readTable(path string) string {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(buf)
}
