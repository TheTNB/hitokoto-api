package hitokoto

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/bytedance/sonic"
)

type Sentence struct {
	Id         int    `json:"id"`
	Uuid       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUid int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

//go:embed sentences
var data embed.FS

var sentences map[string][]Sentence
var sentenceTypes []string

func init() {
	sentences = make(map[string][]Sentence)
	dir, err := data.ReadDir("sentences")
	if err != nil {
		panic(fmt.Sprintf("read hitokoto sentences error: %v", err))
	}

	for _, file := range dir {
		fileData, err := data.ReadFile(fmt.Sprintf("sentences/%s", file.Name()))
		if err != nil {
			panic(fmt.Sprintf("read hitokoto sentence file error: %v", err))
		}

		var sentence []Sentence
		err = sonic.Unmarshal(fileData, &sentence)
		if err != nil {
			panic(fmt.Sprintf("unmarshal hitokoto sentence file error: %v", err))
		}

		// type of sentences
		t := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		sentences[t] = sentence
		sentenceTypes = append(sentenceTypes, t)
	}
}
