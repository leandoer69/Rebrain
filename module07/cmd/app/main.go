package main

import (
	"Rebrain/module07/internal/generator"
	"Rebrain/module07/internal/monitors"
	"fmt"
	"io/ioutil"
)

func Task03() {
	yamlConfiTemplate, err := ioutil.ReadFile("module07/assets/template/config_template.yml")
	if err != nil {
		panic(err)
	}

	_ = generator.ConfigGenerate(string(yamlConfiTemplate), "module07/internal/generator")
}

func Task04() {
	marshallerTemplate, err := ioutil.ReadFile("module07/assets/template/marshaller.gotmpl")
	if err != nil {
		panic(err)
	}

	_ = generator.MarshallerGenerator(
		string(marshallerTemplate),
		"Config",
		"module07/internal/config/config.go",
		"module07/internal/config/codegen_marshaller.go",
	)
}

// Заполнить логикой, когда дойдем до задания номер 5
func Task05() {
	_ = monitors.NewSimpleMonitor()
}

func main() {
	// Task03()
	// Раскоментировать, когда дойдем до задания номер 4
	Task04()
	// Раскоментировать, когда дойдем до задания номер 5
	// Task05()

	fmt.Println("Hello, from module07 module")
}
