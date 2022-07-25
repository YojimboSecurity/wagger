package file_test

import (
	"testing"
	"wagger/pkg/file"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reading file")
}

var _ = Describe("Reading file", func() {

	var highlight map[string]interface{}

	BeforeEach(func() {
		highlight = map[string]interface{}{
			"error": map[string]interface{}{
				"color":      "#282a36",
				"background": "#ff5555",
				"bold":       true,
			},
		}
	})

	When("the file does not exist", func() {
		Context("test for failures", func() {
			It("ReadFile should fail when file does not exist", func() {
				_, err := file.ReadFile("/tmp/does-not-exist")
				Expect(err).To(MatchError("open /tmp/does-not-exist: no such file or directory"))
			})
			It("ReadFileWithOffset should fail when file does not exist", func() {
				_, err := file.ReadFileWithOffset("/tmp/does-not-exist", 0)
				Expect(err).To(MatchError("open /tmp/does-not-exist: no such file or directory"))
			})
		})
	})

	When("the file does exist", func() {
		Context("test for success", func() {
			// TODO: This tests is not complete.
		})
	})

	When("highlighting lines", func() {
		Context("test for success when onlyHighlight is false", func() {
			It("should return non-highlighted string", func() {
				s := file.HighlightString("Hello World", highlight, false)
				Expect(s).To(Equal("Hello World"))
			})
			It("should return empty string ", func() {
				s := file.HighlightString("", highlight, false)
				Expect(s).To(Equal(""))
			})
		})
		Context("test for success when onlyHighlight is true", func() {
			It("should return highlighted string", func() {
				s := file.HighlightString("Error", highlight, true)
				r := []rune(s)
				exp := []rune{27, 91, 49, 59, 51, 56, 59, 50, 59, 52, 48, 59, 52, 50, 59, 53, 52, 59, 52, 56, 59, 50, 59, 50, 53, 53, 59, 56, 53, 59, 56, 53, 109, 69, 114, 114, 111, 114, 27, 91, 48, 109}
				Expect(r).To(Equal(exp))
			})
			It("should return empty string ", func() {
				s := file.HighlightString("this is a test", highlight, true)
				Expect(s).To(Equal(""))
			})
		})
	})

	When("fixing strings", func() {

	})

	When("ignoring strings", func() {

	})
})

func BenchmarkReadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := file.ReadFile("file.go")
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkHighlight(b *testing.B) {
	for n := 0; n < b.N; n++ {
		highlight := map[string]interface{}{
			"test": map[string]interface{}{
				"color":      "#8be9fd",
				"background": "#282a36",
				"bold":       false,
			},
		}
		file.HighlightString("test", highlight, false)
	}
}

func BenchmarkFixString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		file.FixConfigString("test")
	}
}
