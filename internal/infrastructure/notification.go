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

// TODO: delete notifications correctly
func RemoveNotifies() error {
	cmd := exec.Command("atq")
	if err := cmd.Run(); err != nil {
		return err
	}

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	lines := strings.Split(string(out), "\n")
	ids := make([]string, 0, len(lines))
	for _, line := range lines {
		ids = append(ids, strings.Fields(line)[0])
	}

	cmd = exec.Command("atrm", ids...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
