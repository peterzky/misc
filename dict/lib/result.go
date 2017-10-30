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

func (r *Result) Format() string {
	return fmt.Sprintf("%s\n%s\n%s",
		r.Query, r.Basic.Phonetic,
		strings.Join(r.Basic.Explains, "\n"),
	)
}
