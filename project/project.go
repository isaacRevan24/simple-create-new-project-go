package project

import (
	"github.com/isaacRevan24/simple-create-new-project/properties"
)

// TODO: Make structureType only accept concrete and indeterminate type
// Can be concrete or indeterminate
type structureType interface{}

// Project properties definition
type Project struct {
	ProjectName      string
	Icon             uint8
	Manager          string
	Members          map[string]string
	Description      string
	ProjectStructure structureType
}

// generateIndeterminateStructure generate the Indeterminate project structure
func generateIndeterminateStructure(sections []string, startDate string) structureType {
	std := properties.Indeterminate{}
	std = std.CreateIndeterminateProjectStructure(sections, startDate)
	return structureType(std)
}

// generateConcreteStructure generate the concrete project structure
func generateConcreteStructure(sections []string, startDate string, endDate string) structureType {
	std := properties.Concrete{}
	std = std.CreateConcreteProjectStructure(sections, startDate, endDate)
	return structureType(std)
}

// GenerateProject return a project instance
func (p *Project) GenerateProject(typeStruct bool, projectName string, icon uint8, manager string, members map[string]string, description string, sections []string, dates [2]string) Project {

	/*
		1. typeStruct bool
		false = concrete
		true = indeterminate

		2. projectName string

		3. icon uint8

		4. manager string

		5. members [string]string

		6. description string

		7. sections []string

		8. dates [2]string
		dates[0] = start date
		dates[1] = end date

		return structureType
	*/

	// Initialize structure with the return type
	var structure structureType
	// start and end date as independent variables
	startDate := dates[0]
	endDate := dates[1]

	// If true generate a indeterminate structure else a concrete structure
	if typeStruct {
		// indefinite
		structure = generateIndeterminateStructure(sections, startDate)
	} else {
		// concrete
		structure = generateConcreteStructure(sections, startDate, endDate)
	}

	// Project declaration and initialization
	project := Project{
		ProjectName:      projectName,
		Icon:             icon,
		Manager:          manager,
		Members:          members,
		Description:      description,
		ProjectStructure: structure,
	}

	return project
}

// TODO: Add ProjectID as a parameter of the project type struct
