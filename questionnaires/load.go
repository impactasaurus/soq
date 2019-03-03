package questionnaires

import (
	"encoding/json"
	"os"

	"fmt"
	"path/filepath"

	"github.com/impactasaurus/soq"
)

func Load(path string) (soq.Questionnaire, error) {
	f, err := os.Open(path)
	if err != nil {
		return soq.Questionnaire{}, err
	}
	defer f.Close()

	q := soq.Questionnaire{}
	jsonParser := json.NewDecoder(f)
	err = jsonParser.Decode(&q)
	return q, err
}

func LoadDirectory(dir string) ([]soq.Questionnaire, error) {
	files := make([]string, 0)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("found no files at %s\n", dir)
	}
	out := make([]soq.Questionnaire, len(files))
	for idx, path := range files {
		fmt.Printf("loading %s...\n", path)
		q, err := Load(path)
		if err != nil {
			return nil, err
		}
		out[idx] = q
	}

	fmt.Println("questionnaires loaded")
	return out, nil
}
