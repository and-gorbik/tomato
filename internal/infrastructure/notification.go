package infrastructure

import (
	"os/exec"
	"strings"
	"time"
)

func PlanNotifies() error {
	endWork := time.Now().Add(time.Minute * 1)
	cmd := exec.Command("at", endWork.Format("15:04"))
	cmd.Stdin = strings.NewReader("notify-send \"Пора отдыхать!\"")
	if err := cmd.Run(); err != nil {
		return err
	}

	endBreak := endWork.Add(time.Minute * 2)
	cmd = exec.Command("at", endBreak.Format("15:04"))
	cmd.Stdin = strings.NewReader("notify-send \"Пора работать!\"")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
