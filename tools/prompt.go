package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

func mustRun(p promptui.Prompt) string {
	s, err := p.Run()
	if err != nil {
		log.Fatalf("Creating questionnaire failed: %s", err.Error())
	}

	return strings.Trim(s)
}

func mustRunI(p promptui.Prompt) int {
	s := mustRun(p)
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Fatalf("Creating questionnaire failed: %s", err.Error())
	}
	return int(i)
}

func mustRunSP(p promptui.Prompt) *string {
	s := mustRun(p)
	var sp *string
	if s != "" {
		sp = &s
	}
	return sp
}

func mustBool(question string) bool {
	p := promptui.Prompt{
		IsConfirm: true,
		Label:     question,
	}
	_, err := p.Run()
	return err == nil
}
