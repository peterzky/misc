package lib

import (
	"fmt"
	"strings"
)

type basicField struct {
	Phonetic   string   `json:"phonetic"`
	UkPhonetic string   `json:"uk-phonetic"`
	UsPhonetic string   `json:"us-phonetic"`
	Explains   []string `json:"explains"`
}

type webField struct {
	Value []string `json:"value"`
	Key   string   `json:"key"`
}

type Result struct {
	ErrorCode   string       `json:"error_code"`
	Query       string       `json:"query"`
	SpeakUrl    string       `json:"speakUrl"`
	TSpeakUrl   string       `json:"tspeakUrl"`
	Translation *[]string    `json:"translation"`
	Basic       *basicField  `json:"basic"`
	Web         *[]*webField `json:"web"`
}

func formatWeb(list []*webField) string {
	var str string
	for _, wf := range list {

		s := fmt.Sprintf("%s\n%s\n", wf.Key, strings.Join(wf.Value, " "))
		str += s
	}
	return str
}

func (r *Result) Format() string {
	if r.Basic == nil {
		return "Sorry\n No result"
	} else {
		return fmt.Sprintf("%s\n%s\n--------------------\n%s",
			r.Query,
			strings.Join(r.Basic.Explains, "\n"),
			formatWeb(*r.Web),
		)
	}
}
