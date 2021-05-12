package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ConfigMap struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace"`
	} `yaml:"metadata"`
	BinaryData map[string]string `yaml:"binaryData"`
}

func main() {
	var err error
	defer func(err *error) {
		if *err != nil {
			log.Println("exited with error:", (*err).Error())
			os.Exit(1)
		}
	}(&err)

	var optName string
	var optNamespace string
	flag.StringVar(&optName, "n", "", "name of the ConfigMap")
	flag.StringVar(&optNamespace, "ns", "default", "namespace of the ConfigMap")
	flag.Parse()

	if optName == "" {
		err = errors.New("missing argument -n")
		return
	}

	configMap := ConfigMap{
		APIVersion: "v1",
		Kind:       "ConfigMap",
		BinaryData: map[string]string{},
	}

	configMap.Metadata.Name = optName
	configMap.Metadata.Namespace = optNamespace

	for _, arg := range flag.Args() {
		var key string
		var file string
		argSplit := strings.Split(arg, ":")
		if len(argSplit) == 1 {
			key = filepath.Base(argSplit[0])
			file = argSplit[0]
		} else if len(argSplit) == 2 {
			key = argSplit[1]
			file = argSplit[0]
		}
		if key == "" || file == "" {
			err = errors.New("invalid argument '" + arg + "'")
			return
		}
		var buf []byte
		if buf, err = ioutil.ReadFile(file); err != nil {
			return
		}
		configMap.BinaryData[key] = base64.StdEncoding.EncodeToString(buf)
	}

	enc := yaml.NewEncoder(os.Stdout)
	err = enc.Encode(configMap)
}
