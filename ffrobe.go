package go_ffrobe

import (
	"os"
	"os/exec"
	"fmt"
	"encoding/json"
)

type FFProbeMeta struct {
	FilePath  string                 `json:"file_path"`
	Meta      map[string]interface{} `json:"meta"`
	FileError error                  `json:"error"`
}

//default, compact, csv, flat, ini, json, xml

func (self *FFProbeMeta) SetFile(filePath string) *FFProbeMeta {
	if _, err := os.Stat(filePath); err != nil {
		panic(err)
	}
	self.FilePath = filePath

	cmdName, err := exec.LookPath("ffprobe")

	if err != nil {
		self.FileError = err
		return self
	}

	var args []string

	args = append(args, "-print_format")
	args = append(args, "json")
	args = append(args, "-show_streams")
	args = append(args, "-v")
	args = append(args, "error")
	args = append(args, self.FilePath)

	cmd := exec.Command(cmdName, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		self.FileError = err
		return self
	}
	var meta map[string]interface{}
	json.Unmarshal([]byte(string(output)), &meta)
	self.Meta = meta
	return self

}

func (self *FFProbeMeta) GetMeta() (err error, meta []interface{}) {
	return self.FileError, self.Meta["streams"].([]interface{})
}

