package commands

import (
	utilsJson "github.com/Korbeil/slurp/utils/json"
	"github.com/Korbeil/slurp/utils/directory"
	"encoding/json"
	"os"
	"fmt"
)

func loadEnv(homeDir string) Environment {
	envPath := homeDir+"/.slurp/env.json"

	// env doesn't exists, let's create it and return an empty one
	if directory.CheckExists(envPath) == false {
		fmt.Printf("Bash history file: `%s`", os.Getenv("HISTFILE"))
		env := Environment{Project: "", OldBashHistoryPath: os.Getenv("HISTFILE")}
		utilsJson.WriteJsonInFile(
			env,
			envPath)
		return env
	}

	// or just read it :)
	var b []byte
	b = utilsJson.ReadJson(envPath)

	var env Environment
	json.Unmarshal(b, &env)
	return env
}

func loadProject(homeDir string, projectName string) Project {
	var b []byte
	b = utilsJson.ReadJson(homeDir+"/.slurp/projects/"+projectName+"/config.json")

	var conf Project
	json.Unmarshal(b, &conf)
	return conf
}