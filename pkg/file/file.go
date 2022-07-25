package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"wagger/log"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ReadFileWithOffset(filePath string, offset int64) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	file.Seek(offset, 0)
	return ioutil.ReadAll(file)
}

func ReadFile(filePath string) ([]byte, error) {
	return ReadFileWithOffset(filePath, 0)
}

func HighlightString(line string, highlight map[string]interface{}, onlyHighlight bool) string {
	toLower := cases.Lower(language.Und)
	newS := toLower.String(line)
	for k := range highlight {
		h := highlight[k].(map[string]interface{})
		newK := k
		newK = FixConfigString(newK)
		if strings.Contains(newS, newK) {
			style := lipgloss.NewStyle().Foreground(lipgloss.Color(h["color"].(string))).Background(lipgloss.Color(h["background"].(string))).Bold(h["bold"].(bool))
			return style.Render(line)
		}
	}
	if !onlyHighlight {
		return line
	}
	return ""
}

func FixConfigString(s string) string {
	return strings.ReplaceAll(s, "_", " ")
}

func IgnoreString(s string, ignore []string) bool {
	for _, k := range ignore {
		if strings.Contains(s, k) {
			return true
		}
	}
	return false
}

func HighlightAndIgnore(ss []string, highlight map[string]interface{}, ign []string, onlyHighlight, ignore bool) {
	for _, line := range ss {
		if ignore {
			if IgnoreString(line, ign) {
				continue
			}
		}
		line = HighlightString(line, highlight, onlyHighlight)
		if line != "" {
			fmt.Println(line)
		}
	}
}

func TailFile(filePath string, highlight map[string]interface{}, ign []string, onlyHighlight, ignore bool, d time.Duration) {
	var flen int64
	st, err := os.Stat(filePath)
	if err != nil {
		log.Error(err)
		return
	}
	flen = st.Size()
	log.Debug("Start file length: ", flen)
	f, err := os.Open(filePath)
	if err != nil {
		log.Error(err)
		return
	}
	dat := make([]byte, flen)
	f.Read(dat)
	f.Close()
	ds := strings.Split(string(dat), "\n")
	HighlightAndIgnore(ds, highlight, ign, onlyHighlight, ignore)
	for {
		// time.Sleep(time.Second * 1)
		time.Sleep(d)
		sta, err := os.Stat(filePath)
		if err != nil {
			log.Error(err)
			return
		}
		if sta.Size() == flen {
			log.Debug("File length not changed ", sta.Size()-flen)
			continue
		}
		dif := sta.Size() - flen
		log.Debug("File length changed ", dif)
		if dif < 0 {
			dif = sta.Size()
			flen = sta.Size()
		}
		data := make([]byte, dif)
		f, err := os.Open(filePath)
		if err != nil {
			log.Error(err)
			return
		}
		f.Seek(flen, 0)
		f.Read(data)
		f.Close()
		ds := strings.Split(string(data), "\n")
		HighlightAndIgnore(ds, highlight, ign, onlyHighlight, ignore)
		flen = sta.Size()
	}
}
