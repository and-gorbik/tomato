package models

type Config struct {
	DefaultTag       string `yaml:"default_tag"`
	EditorPath       string `yaml:"editor_path"`
	CurrentTasksPath string `yaml:"current_tasks_path"`
	DBPath           string `yaml:"db_path"`
	Tomato           Tomato `yaml:"tomato"`
}

type Tomato struct {
	SmallBreakMinutes int `yaml:"small_break_period"`
	BigBreakMinutes   int `yaml:"big_break_period"`
	WorkingMinutes    int `yaml:"working_period"`
}
