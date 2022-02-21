package questionnaires

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/impactasaurus/soq"
	"github.com/pkg/errors"
)

//go:embed *.json
var qFS embed.FS

func Load(name string) (soq.Questionnaire, error) {
	f, err := qFS.Open(name)
	if err != nil {
		return soq.Questionnaire{}, err
	}
	defer f.Close()

	q := soq.Questionnaire{}
	jsonParser := json.NewDecoder(f)
	err = jsonParser.Decode(&q)
	return q, err
}

func LoadAll() ([]soq.Questionnaire, error) {
	files, err := qFS.ReadDir(".")
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, errors.New("found no questionnaires")
	}
	out := make([]soq.Questionnaire, len(files))
	for idx, path := range files {
		q, err := Load(path.Name())
		if err != nil {
			return nil, err
		}
		out[idx] = q
	}

	fmt.Printf("%d questionnaires loaded\n", len(files))
	return out, nil
}
