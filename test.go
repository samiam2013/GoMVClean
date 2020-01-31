package main

import (
	"fmt"
)

func testEverything(hasToPass bool) {
	passM := testModel(hasToPass)
	passV := testView(hasToPass)
	passC := testController(hasToPass)
	if passM && passV && passC {
		fmt.Println("pass.")
	} else if !passM {
		fmt.Println("failed testModel()......")
	} else if !passV {
		fmt.Println("failed testView()")
	} else if !passC {
		fmt.Println("failed the testController()")
	}
}

func testModel(hasToPass bool) bool {
	fmt.Println("testing the Model........")
	testModelPath := modelSchemaRead + "/" + modelSchemaFolder + modelMarkup
	fmt.Println("testing model path: ", testModelPath)
	schema, err := loadStaticBody(testModelPath)
	if !gDebug && err {
		fmt.Println("schema.json:", string(schema))
		return testSchema(string(schema))
	} else if !gDebug {
		if !err {
			fmt.Println("could not retrieve Json Schema!")
			return false
		}
	}
	return true
}

func testSchema(schema string) bool {
	fmt.Println("testing the Schema.......")
	return true
}

func testView(hasToPass bool) bool {
	fmt.Println("testing the View.........")
	return true
}

func testController(hasToPass bool) bool {
	fmt.Println("testing the Controller!...")
	return true
}
