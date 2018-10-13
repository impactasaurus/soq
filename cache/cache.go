package cache

import (
	"encoding/json"
	"errors"

	"os"
	"path/filepath"

	"sort"

	"fmt"

	soq "github.com/impactasaurus/soq-api"
)

type Cache struct {
	byID   map[string]soq.Questionnaire
	byName []*soq.Questionnaire
}

func loadQuestionnaire(path string) (soq.Questionnaire, error) {
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

func New(questionnaireDirectory string) (*Cache, error) {
	files := make([]string, 0)
	err := filepath.Walk(questionnaireDirectory, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".json" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("found no files at %s\n", questionnaireDirectory)
	}
	c := &Cache{
		byID:   map[string]soq.Questionnaire{},
		byName: make([]*soq.Questionnaire, len(files)),
	}

	for idx, path := range files {
		fmt.Printf("loading %s...\n", path)
		q, err := loadQuestionnaire(path)
		if err != nil {
			return nil, err
		}
		c.byName[idx] = &q
		c.byID[q.ID] = q
	}
	sort.Slice(c.byName, func(i, j int) bool {
		return c.byName[i].Name < c.byName[j].Name
	})
	fmt.Println("questionnaires loaded")
	return c, nil
}

func (c *Cache) Questionnaire(id string) (soq.Questionnaire, error) {
	q, ok := c.byID[id]
	if !ok {
		return soq.Questionnaire{}, errors.New("could not find questionnaire")
	}
	return q, nil
}

func (c *Cache) Questionnaires(page, limit int) (soq.QuestionnaireList, error) {
	offset := page * limit
	if offset >= len(c.byName) {
		return soq.QuestionnaireList{
			PageInfo: soq.PageInfo{
				Limit:       limit,
				HasNextPage: false,
				Page:        page,
			},
			Questionnaires: []*soq.Questionnaire{},
		}, nil
	}
	end := offset + limit
	if end > len(c.byName) {
		end = len(c.byName)
	}
	more := end != len(c.byName)
	return soq.QuestionnaireList{
		PageInfo: soq.PageInfo{
			Limit:       limit,
			HasNextPage: more,
			Page:        page,
		},
		Questionnaires: c.byName[offset:end],
	}, nil
}
