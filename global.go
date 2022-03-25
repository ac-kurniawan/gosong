package gosong

var GlobalDependencies []Dependency

func AddGlobalComponent(name string, component interface{}) {
	buffer := Dependency{
		Name:   name,
		Inface: component,
	}
	GlobalDependencies = append(GlobalDependencies, buffer)
}
