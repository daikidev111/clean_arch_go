package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Entity struct {
	Package string
	Model   string
	Models  string
}
type DomainRepository struct {
	Package string
	Model   string
}
type DomainService struct {
	Package        string
	Model          string
	ModelLowerCase string
}
type DatabaseRepository struct {
	Package string
	Model   string
}
type Interactor struct {
	Package        string
	Model          string
	ModelLowerCase string
}
type Controller struct {
	Package        string
	Model          string
	ModelLowerCase string
}

const (
	fileExtension          = ".go"
	entityPath             = "./../../pkg/domain/entity/"
	domainRepositoryPath   = "./../../pkg/domain/repository/"
	domainServicePath      = "./../../pkg/domain/service/"
	databaseRepositoryPath = "./../../pkg/driver/mysql/database/"
	interactorPath         = "./../../pkg/usecase/"
	controllerPath         = "./../../pkg/adapter/controllers/"
)

//  go run generator.go -model="User" -file=""

// WARNING: space after the , will be determined as an invalid command
func main() {
	m := flag.String("model", "", "This is to generate files based on the model name")
	f := flag.String("file", "", "This is to determine which files to generate")
	flag.Parse()
	file := strings.Split(*f, ",")

	model, modelFCharLowerCase := convertModelInputToModel(m)
	if model == "" && modelFCharLowerCase == "" {
		return
	}

	errors := fileSelection(model, modelFCharLowerCase, file)
	for _, err := range errors {
		if err != nil {
			fmt.Println("[ERROR] createFile functions or from fileSelection(model, modelFCharLowerCase, file) = ", err)
		}
	}
}

func createEntityFile(model string) error {
	if _, err := os.Stat(entityPath); os.IsNotExist(err) {
		return err
	}
	data := Entity{
		Package: "entity",
		Model:   model,
		Models:  model + "s",
	}
	tmpl, err := template.New("entity.tmpl").ParseFiles("entity.tmpl")
	if err != nil {
		return err
	}

	fileName := strings.ToLower(model) + fileExtension

	// ファイルの生成
	fp, err := os.Create(filepath.Join(entityPath, fileName))
	if err != nil {
		return err
	}
	defer fp.Close()

	// ファイルへの書き込み
	if err := tmpl.Execute(fp, data); err != nil {
		return err
	}
	return nil
}

func createDomainRepositoryFile(model string) error {
	if _, err := os.Stat(domainRepositoryPath); os.IsNotExist(err) {
		return err
	}
	data := DomainRepository{
		Package: "repository",
		Model:   model,
	}
	tmpl, err := template.New("domain_repo.tmpl").ParseGlob("domain_repo.tmpl")
	if err != nil {
		return err
	}

	fileName := strings.ToLower(model) + "_repository" + fileExtension

	// ファイルの生成
	fp, err := os.Create(filepath.Join(domainRepositoryPath, fileName))
	if err != nil {
		return err
	}
	defer fp.Close()

	// ファイルへの書き込み
	if err := tmpl.Execute(fp, data); err != nil {
		return err
	}
	return nil
}

func createServiceFile(model, modelLower string) error {
	if _, err := os.Stat(domainServicePath); os.IsNotExist(err) {
		return err
	}
	data := DomainService{
		Package:        "service",
		Model:          model,
		ModelLowerCase: modelLower,
	}
	tmpl, err := template.New("domain_service.tmpl").ParseGlob("domain_service.tmpl")
	if err != nil {
		return err
	}

	fileName := data.ModelLowerCase + "_service" + fileExtension

	fp, err := os.Create(filepath.Join(domainServicePath, fileName))
	if err != nil {
		return err
	}
	defer fp.Close()

	// ファイルへの書き込み
	if err := tmpl.Execute(fp, data); err != nil {
		return err
	}
	return nil
}

func createDatabaseRepositoryFile(model string) error {
	if _, err := os.Stat(databaseRepositoryPath); os.IsNotExist(err) {
		return err
	}
	data := DatabaseRepository{
		Package: "database",
		Model:   model,
	}
	tmpl, err := template.New("database_repository.tmpl").ParseGlob("database_repository.tmpl")
	if err != nil {
		return err
	}

	fileName := strings.ToLower(model) + "_repository" + fileExtension

	fp, err := os.Create(filepath.Join(databaseRepositoryPath, fileName))
	if err != nil {
		return err
	}
	defer fp.Close()

	// ファイルへの書き込み
	if err := tmpl.Execute(fp, data); err != nil {
		return err
	}
	return nil
}

func createInteractorFile(model, modelLower string) error {
	if _, err := os.Stat(interactorPath); os.IsNotExist(err) {
		return err
	}
	data := Interactor{
		Package:        "usecase",
		Model:          model,
		ModelLowerCase: modelLower,
	}
	tmpl, err := template.New("interactor.tmpl").ParseGlob("interactor.tmpl")
	if err != nil {
		return err
	}

	fileName := data.ModelLowerCase + "_interactor" + fileExtension

	fp, err := os.Create(filepath.Join(interactorPath, fileName))
	if err != nil {
		return err
	}
	defer fp.Close()

	// ファイルへの書き込み
	if err := tmpl.Execute(fp, data); err != nil {
		return err
	}
	return nil
}

func createControllerFile(model, modelLower string) error {
	if _, err := os.Stat(controllerPath); os.IsNotExist(err) {
		return err
	}
	data := Controller{
		Package:        "controllers",
		Model:          model,
		ModelLowerCase: modelLower,
	}
	tmpl, err := template.New("controller.tmpl").ParseGlob("controller.tmpl")
	if err != nil {
		return err
	}

	fileName := data.ModelLowerCase + "_controller" + fileExtension

	fp, err := os.Create(filepath.Join(controllerPath, fileName))
	if err != nil {
		return err
	}
	defer fp.Close()

	// ファイルへの書き込み
	if err := tmpl.Execute(fp, data); err != nil {
		return err
	}
	return nil
}

// error とfile name をマップの中に詰め込む
func FirstCharToLowerCase(model string) string {
	s := model
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	s = string(a)
	return s
}

//nolint: gocyclo // this is why
func fileSelection(model, modelFCharLowerCase string, file []string) []error {
	for _, v := range file {
		v = strings.ToLower(v) // make all the chars lower case in case of some capitalized chars
		validate := true
		var err error
		errors := make([]error, 6)
		switch validate {
		case v == "":
			errors = append(errors, fmt.Errorf("does not accept empty file or invalid file name"))
			return errors

		case v == "databaserepo" || v == "databaseRepo" || v == "Databaserepo":
			err = createDatabaseRepositoryFile(model)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to generate a database repository file = %w", err))
				return errors
			}
			fmt.Printf("successfully generated a file %s\n", v)

		case v == "entity":
			err = createEntityFile(model)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to generate an entity file = %w", err))
				return errors
			}
			fmt.Printf("successfully generated a file %s\n", v)

		case v == "domainRepository" || v == "Domainrepo" || v == "domainrepo" || v == "Domainrepository":
			err = createDomainRepositoryFile(model)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to generate a domain repository file = %w", err))
				return errors
			}
			fmt.Printf("successfully generated a file %s\n", v)

		case v == "domainService" || v == "Domainservice":
			err = createServiceFile(model, modelFCharLowerCase)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to generate a service file = %w", err))
				return errors
			}
			fmt.Printf("successfully generated a file %s\n", v)

		case v == "interactor" || v == "int":
			err = createInteractorFile(model, modelFCharLowerCase)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to generate a service file = %w", err))
				return errors
			}
			fmt.Printf("successfully generated a file %s\n", v)

		case v == "controller" || v == "cont":
			err = createControllerFile(model, modelFCharLowerCase)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to generate a service file = %w", err))
				return errors
			}
			fmt.Printf("successfully generated a file %s\n", v)

		case v == "all":
			if len(file) > 1 {
				errors = append(errors, fmt.Errorf("cannot choose all along with other file selections"))
			}

			err = createEntityFile(model)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to generate an entity file = %w", err))
			}

			err = createDomainRepositoryFile(model)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to generate a domain repository file = %w", err))
			}

			err = createServiceFile(model, modelFCharLowerCase)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to generate a service file = %w", err))
			}

			err = createDatabaseRepositoryFile(model)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to generate a database repository file = %w", err))
			}

			err = createInteractorFile(model, modelFCharLowerCase)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to generate a service file = %w", err))
			}

			err = createControllerFile(model, modelFCharLowerCase)
			if err != nil {
				errors = append(errors, fmt.Errorf("failed to generate a service file = %w", err))
			}
			return errors
		default:
			errors = append(errors, fmt.Errorf("error with the file generation. Please check if it is entered according to its instruction"))
			return errors
		}
	}
	return nil
}

func convertModelInputToModel(modelInput *string) (model, modelFCharLowerCase string) {
	words := strings.Fields(*modelInput)

	if len(words) > 1 {
		for _, v := range words {
			model += strings.Title(strings.ToLower(v))
		}
		modelFCharLowerCase = FirstCharToLowerCase(model)
	} else if len(words) == 1 {
		model = strings.Title(strings.ToLower(*modelInput))
		modelFCharLowerCase = FirstCharToLowerCase(model)
	} else {
		log.Println("Does not accept empty word")
		return "", ""
	}
	return model, modelFCharLowerCase
}
