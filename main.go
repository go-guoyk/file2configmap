package main

import (
	"encoding/base64"
	"flag"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
	flag.StringVar(&optName, "name", "unknown", "name of the ConfigMap")
	flag.StringVar(&optNamespace, "namespace", "default", "namespace of the ConfigMap")
	flag.Parse()

	configMap := ConfigMap{
		APIVersion: "v1",
		Kind:       "ConfigMap",
		BinaryData: map[string]string{},
	}

	configMap.Metadata.Name = optName
	configMap.Metadata.Namespace = optNamespace

	for _, file := range flag.Args() {
		var buf []byte
		if buf, err = ioutil.ReadFile(file); err != nil {
			return
		}
		configMap.BinaryData[filepath.Base(file)] = base64.StdEncoding.EncodeToString(buf)
	}

	enc := yaml.NewEncoder(os.Stdout)
	err = enc.Encode(configMap)
}
