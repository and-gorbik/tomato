package app

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"tomato/internal/models"
)

type TaskFile struct {
	Path string
}

func (t *TaskFile) Load(fname string, allowRewriteNonEmptyFile bool) (err error) {
	src, err := os.Open(fname)
	if err != nil {
		return
	}

	defer src.Close()

	if err = t.validate(src); err != nil {
		return
	}

	if _, err = src.Seek(0, 0); err != nil {
		return
	}

	dst, err := os.OpenFile(t.Path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return
	}

	defer dst.Close()

	dstInfo, err := dst.Stat()
	if err != nil {
		return
	}

	if dstInfo.Size() != 0 && !allowRewriteNonEmptyFile {
		return fmt.Errorf("Current tasks is not finished")
	}

	_, err = io.Copy(dst, src)
	return
}

func (t *TaskFile) Print() error {
	f, err := os.Open(t.Path)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := io.Copy(os.Stdout, f); err != nil {
		return err
	}

	return nil
}

// Next reads the next line from the file.
// If count > 1, its value decrements, else line is removed.
func (t *TaskFile) Next() (models.Task, error) {
	return models.Task{}, nil
}

func (t *TaskFile) Prepend(title string, tag *string) error {
	return nil
}

func (t *TaskFile) Validate() error {
	return nil
}

func (t *TaskFile) validate(src io.Reader) error {
	lines, err := csv.NewReader(src).ReadAll()
	if err != nil {
		return err
	}

	// title;[tag;][count;][position;]
	positions := make(map[int64]struct{})
	for _, line := range lines {
		switch len(line) {
		case 1, 2:
		case 3:
			pos, err := strconv.ParseInt(line[2], 10, 32)
			if err != nil {
				return err
			}
			if _, ok := positions[pos]; ok {
				return fmt.Errorf("Repeatable position: %d", pos)
			}
			positions[pos] = struct{}{}
		case 4:
			pos, err := strconv.ParseInt(line[2], 10, 32)
			if err != nil {
				return err
			}
			if _, ok := positions[pos]; ok {
				return fmt.Errorf("Repeatable position: %d", pos)
			}
			positions[pos] = struct{}{}

			if _, err := strconv.ParseInt(line[3], 10, 32); err != nil {
				return err
			}
		}
	}

	return nil
}
