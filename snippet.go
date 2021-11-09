package snippet

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/google/uuid"
	"howett.net/plist"
)

type (
	Snippet struct {
		Snippet string `json:"snippet"`
		UID     string `json:"uid"`
		Name    string `json:"name"`
		Keyword string `json:"keyword"`
	}
	Snippets struct {
		KeywordPrefix string
		KeywordSuffix string
		Snippets      []Snippet
	}
)

func New(keywordPrefix string, keywordSuffix string) *Snippets {
	return &Snippets{keywordPrefix, keywordSuffix, nil}
}

func (s *Snippets) Add(name string, snippet string, keyword string) {
	uid := uuid.New()

	s.Snippets = append(s.Snippets, Snippet{
		Snippet: snippet,
		UID:     uid.String(),
		Name:    name,
		Keyword: keyword,
	})
}

func (s *Snippets) Save(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	w := zip.NewWriter(f)
	defer w.Close()
	if f, err := w.Create("info.plist"); err == nil {
		s.writePlist(f)
	} else {
		return err
	}
	for _, snippet := range s.Snippets {
		if f, err := w.Create(snippet.fileName()); err == nil {
			if err := s.writeSnippet(snippet, f); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func (s *Snippets) writePlist(w io.Writer) error {
	enc := plist.NewEncoder(w)
	return enc.Encode(map[string]string{
		"snippetkeywordprefix": s.KeywordPrefix,
		"snippetkeywordsuffix": s.KeywordSuffix,
	})
}

func (s Snippet) fileName() string {
	return escapeFilename(fmt.Sprintf("%s [%s].json", s.Name, s.UID))
}

func escapeFilename(name string) string {
	r := regexp.MustCompile(`[\n~"#%&\*:<>\?/\{\|\}]`)
	return r.ReplaceAllString(name, " ")
}

func (s *Snippets) writeSnippet(snippet Snippet, w io.Writer) error {
	v := make(map[string]interface{})
	v["alfredsnippet"] = snippet
	enc := json.NewEncoder(w)
	return enc.Encode(v)
}
