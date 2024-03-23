package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/davecgh/go-spew/spew"
	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v2"
)

//go:embed templates/combat.tpl
var combatRaw string

//go:embed templates/common.go.tpl
var commonGo string

//go:embed templates/parser.go.tpl
var parserGo string

var typeByName = map[string]string{
	"LogTime":         "time.Time",
	"InitiatorID":     "int",
	"RecipientID":     "int",
	"DamageModifiers": "DamageModifiersMap",
	"FriendlyFire":    "bool",
	"DamageTotal":     "float32",
	"DamageHull":      "float32",
	"DamageShield":    "float32",
	"Heal":            "float32",
}

var parserByType = map[string]string{
	"time.Time":          "parseTime(now)",
	"int":                "strconv.Atoi",
	"float32":            "parseFloat",
	"bool":               "parseBool",
	"string":             "",
	"DamageModifiersMap": "parseDamageModifiers",
}

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

func (g *Generator) GenerateRegexps(rawTpl string) (map[string]string, error) {
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

	var regexps map[string]string
	err := yaml.Unmarshal(sb.Bytes(), &regexps)
	if err != nil {
		println(sb.String())
		return nil, fmt.Errorf("yaml.Unmarshal: %w", err)
	}
	return regexps, nil
}

func (g *Generator) MakeConfigs(regexps map[string]string) (configs []FileConfig, err error) {
	for reName, raw := range regexps {
		re, err := regexp.Compile(raw)
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
				Value:        re.String(),
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
	Value        string
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

func (g *Generator) GenerateFiles(configs []FileConfig) (map[string]string, error) {
	res := make(map[string]string)
	var buf bytes.Buffer
	if err := g.commonGoTpl.Execute(&buf, configs); err != nil {
		return nil, fmt.Errorf("generate common.go file: %w", err)
	}
	res["common.go"] = buf.String()

	for _, config := range configs {
		var buf bytes.Buffer
		if err := g.parserGoTpl.Execute(&buf, config); err != nil {
			spew.Dump(config)
			return nil, fmt.Errorf("generate config file(%s): %w", config.TypeName, err)
		}

		res[strcase.ToSnake(config.TypeName)+".go"] = buf.String()
	}

	return res, nil
}

const filePerm = 0600

func (g *Generator) WriteResults(files map[string]string) error {
	for name, contents := range files {
		fileName := filepath.Join(g.cfg.OutputDir, name)
		err := os.WriteFile(fileName, []byte(contents), filePerm)
		if err != nil {
			return fmt.Errorf("write file(%s): %w", fileName, err)
		}
	}

	return nil
}
