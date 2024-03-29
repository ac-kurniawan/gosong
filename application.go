package gosong

import (
	"fmt"
	"reflect"
)

type IApplication interface {
	Run()
}

type Dependency struct {
	Name   string
	Inface interface{}
}
type Tag struct {
	Name string
	Tag  string
}

type Application struct {
	Name string
	IApplication
	singleton   []Dependency
	controllers []Dependency
	entries     []func()
}

func (a *Application) Run() {
	for _, entry := range a.entries {
		entry()
	}
}

// AddEntry controller entries or route setting
func (a *Application) AddEntry(entry func()) {
	a.entries = append(a.entries, entry)
}

// AddController register controller/resolver/consumer
func (a *Application) AddController(name string, controller interface{}) {
	for _, tag := range a.findTag(controller) {
		result := a.findDependenciesByName(tag.Tag)

		if result != nil {
			a.bind(controller, result, tag.Name)
		}
	}
	buffer := Dependency{
		Name:   name,
		Inface: controller,
	}
	a.controllers = append(a.controllers, buffer)
}

// AddSingleton register providers/service/repository
func (a *Application) AddSingleton(name string, provider interface{}) {
	for _, tag := range a.findTag(provider) {
		result := a.findDependenciesByName(tag.Tag)
		if result != nil {
			a.bind(provider, result, tag.Name)
		}
	}
	buffer := Dependency{
		Name:   name,
		Inface: provider,
	}
	a.singleton = append(a.singleton, buffer)
}

func (a *Application) findDependenciesByName(name string) interface{} {
	var merge []Dependency
	merge = append(merge, GlobalDependencies...)
	merge = append(merge, a.singleton...)
	var found Dependency

	for _, dep := range merge {
		if dep.Name == name {
			found = dep
		}
	}

	return found.Inface
}

func (a *Application) bind(srv interface{}, repo interface{}, fieldName string) {
	v := reflect.ValueOf(srv).Elem()
	v.FieldByName(fieldName).Set(reflect.ValueOf(repo))
}

func (a *Application) findTag(srv interface{}) []Tag {
	f := reflect.TypeOf(srv).Elem()
	var tags []Tag

	for i := 0; i < f.NumField(); i++ {
		field := f.Field(i)
		tag := string(field.Tag.Get("import"))
		if tag != "" {
			buffer := Tag{
				Name: field.Name,
				Tag:  tag,
			}
			tags = append(tags, buffer)
		}

	}
	return tags
}

func RunApplications(apps ...Application) {
	for _, app := range apps {
		app.Run()
		fmt.Printf("Application: %s%s %s- started \n", string("\033[32m"), app.Name, string("\033[0m"))
	}
}
