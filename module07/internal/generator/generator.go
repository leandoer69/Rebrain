package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"text/template"
)

const (
	marshallStructKey       = "marshallStruct"
	marshallStructFieldsKey = "marshallStructFields"
)

// Task04 - функция для генерации маршалера структуры в мапу
func MarshallerGenerator(marshallerTemplate string, structName string, inFilePath string, outFilePath string) error {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, "module07/internal/config/config.go",
		nil, parser.ParseComments)
	if err != nil {
		return err
	}

	data := struct {
		packag string
	}{}
	ast.Inspect(node, func(n ast.Node) bool {
		if list, ok := n.(*ast.Field); ok {
			fmt.Println((*list).Names[0])

		}
		return true
	})

	err = generate(marshallerTemplate, outFilePath, data)
	if err != nil {
		return err
	}

	return nil
}

// Task03 - структура для yaml конфигурации
// Если нужно расширить yaml файл, тогда эту структуру нужно также расширить
// нужными параметрами, yaml файл и эта структура должны быть идентичными.
type Config struct {
	Name       string
	Port       string
	ReplicaSet int
	ImageName  string
	Tag        string
	EnvPath    string
}

// Task03 - функция генерации конфига
func ConfigGenerate(tmpl string, outFilePath string) error {
	config := Config{
		Name:       "Project",
		Port:       "8080",
		ReplicaSet: 0,
		ImageName:  "pgAdmin",
		Tag:        "latest",
		EnvPath:    "",
	}

	err := generate(tmpl, outFilePath, config)
	if err != nil {
		return err
	}
	return nil
}

// Основная функция которая должна производить генерацию
func generate(tmpl string, outfilePath string, fields interface{}) error {
	t, err := template.New("config").Parse(tmpl)
	if err != nil {
		return err
	}

	file, err := os.Create(outfilePath)
	if err != nil {
		return err
	}

	err = t.Execute(file, fields)
	if err != nil {
		return err
	}

	return nil
}
