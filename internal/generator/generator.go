package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/davecgh/go-spew/spew"
	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v2"

	"golang.org/x/exp/maps"
	"golang.org/x/tools/imports"
)

//go:embed templates/common.gotpl
var commonGo string

//go:embed templates/parser.gotpl
var parserGo string

type Generator struct {
	parserGoTpl *template.Template
	commonGoTpl *template.Template
	cfg         *Config
}

func NewGenerator(cfg *Config) (*Generator, error) {
	parserTpl := template.Must(template.New("parser.go.tpl").Funcs(sprig.TxtFuncMap()).Parse(parserGo))
	commonTpl := template.Must(template.New("common.go.tpl").Funcs(sprig.TxtFuncMap()).Parse(commonGo))

	return &Generator{
		parserGoTpl: parserTpl,
		commonGoTpl: commonTpl,
		cfg:         cfg,
	}, nil
}

// Generate создает файлы со структурами регулярок из данного шаблона, который формирует yaml файл
func (g *Generator) Generate(rawTpl string) error {
	regexps, err := g.GenerateRegexps(rawTpl)
	if err != nil {
		return fmt.Errorf("error generating regexps: %v", err)
	}
	configs, err := g.MakeConfigs(regexps)
	if err != nil {
		return fmt.Errorf("make configs: %w", err)
	}

	files, err := g.GenerateFiles(configs)
	if err != nil {
		return fmt.Errorf("error generating files: %w", err)
	}

	if err := g.WriteResults(files); err != nil {
		return fmt.Errorf("write results: %w", err)
	}

	return nil
}

type Regexps struct {
	Full  string
	Short string
}

type RegexpsMap map[string]Regexps

func (g *Generator) GenerateRegexps(rawTpl string) (RegexpsMap, error) {
	tpl := template.Must(
		template.New("").
			Funcs(template.FuncMap{
				"wrap": func(name, re string) string {
					return fmt.Sprintf("(?P<%s>%s)", name, re)
				},
			}).
			Parse(rawTpl))
	var sb bytes.Buffer
	if err := tpl.Execute(&sb, nil); err != nil {
		return nil, fmt.Errorf("template.Execute: %w", err)
	}

	var data struct {
		Regexps []struct {
			Name  string `yaml:"name"`
			Full  string `yaml:"full"`
			Short string `yaml:"short"`
		} `yaml:"regexps"`
	}
	err := yaml.Unmarshal(sb.Bytes(), &data)
	if err != nil {
		for idx, line := range strings.Split(sb.String(), "\n") {
			fmt.Printf("%3d | %s\n", idx, line)
		}
		return nil, fmt.Errorf("yaml.Unmarshal: %w", err)
	}

	regexps := make(RegexpsMap, len(data.Regexps))
	for _, re := range data.Regexps {
		if re.Short == "" {
			re.Short = re.Full
		}
		regexps[re.Name] = Regexps{
			Full:  re.Full,
			Short: re.Short,
		}
	}
	return regexps, nil
}

func (g *Generator) MakeConfigs(regexps RegexpsMap) (configs []FileConfig, err error) {
	keys := maps.Keys(regexps)
	slices.Sort(keys)
	for _, reName := range keys {
		raw := regexps[reName]
		re, err := regexp.Compile(raw.Full)
		if err != nil {
			return nil, fmt.Errorf("compile regexp: %w", err)
		}

		fieldNamesConst := re.SubexpNames()
		var fieldNames []string

		for _, name := range fieldNamesConst {
			if name != "" {
				fieldNames = append(fieldNames, name)
			}
		}

		var fields []FieldConfig
		for _, name := range fieldNames {
			typeName, ok := typeByName[name]
			if !ok {
				typeName = "string"
			}
			parserFunc, ok := parserByType[typeName]
			if !ok {
				return nil, fmt.Errorf("unexpected type: %v", typeName)
			}
			fields = append(fields, FieldConfig{
				Name: name,
				Type: FieldType{
					Name:      typeName,
					ParseFunc: parserFunc,
					IsString:  typeName == "string",
				},
				RegexpIndex: re.SubexpIndex(name),
			})
		}

		configs = append(configs, FileConfig{
			PackageName: g.cfg.OutputPackage,
			TypeName:    strcase.ToCamel(reName),
			Regexp: RegexpParams{
				Name:         "re" + strcase.ToCamel(reName),
				ShortName:    "shortRe" + strcase.ToCamel(reName),
				Value:        raw,
				TotalMatches: re.NumSubexp() + 1,
			},
			Fields: fields,
		})
	}
	return configs, nil
}

type FileConfig struct {
	PackageName string
	TypeName    string
	Regexp      RegexpParams
	Fields      []FieldConfig
}

type RegexpParams struct {
	Name         string
	ShortName    string
	Value        Regexps
	TotalMatches int
}

type FieldConfig struct {
	Name        string
	Type        FieldType
	RegexpIndex int
}

type FieldType struct {
	Name      string
	ParseFunc string
	IsString  bool
}

func (g *Generator) GenerateFiles(configs []FileConfig) (map[string][]byte, error) {
	res := make(map[string][]byte)
	var buf bytes.Buffer
	if err := g.commonGoTpl.Execute(&buf, configs); err != nil {
		return nil, fmt.Errorf("generate common.go file: %w", err)
	}

	formatted, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		return nil, fmt.Errorf("format common.go file: %w", err)
	}
	res["common.go"] = formatted

	for _, config := range configs {
		var buf bytes.Buffer
		if err := g.parserGoTpl.Execute(&buf, config); err != nil {
			spew.Dump(config)
			return nil, fmt.Errorf("generate config file(%s): %w", config.TypeName, err)
		}

		fileName := strcase.ToSnake(config.TypeName) + ".go"
		formatted, err := imports.Process("", buf.Bytes(), nil)
		if err != nil {
			println(buf.String())
			return nil, fmt.Errorf("format %s file: %w", fileName, err)
		}
		res[fileName] = formatted
	}

	return res, nil
}

const filePerm = 0600

func (g *Generator) WriteResults(files map[string][]byte) error {
	for name, contents := range files {
		fileName := filepath.Join(g.cfg.OutputDir, name)
		err := os.WriteFile(fileName, contents, filePerm)
		if err != nil {
			return fmt.Errorf("write file(%s): %w", fileName, err)
		}
	}

	return nil
}
