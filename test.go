package main

import (
	"fmt"
)

//this has a self-explanatory name
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

// self-explanatory
func testModel(hasToPass bool) bool {
	fmt.Println("testing the Model........")
	testModelPath := modelSchemaPub + "/" + schemaFolder + modelMarkup
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

// ...
func testSchema(schema string) bool {
	fmt.Println("testing the Public schema.......")
	jsonBody, err := loadStaticBody(modelPrivPath)
	if err {
		fmt.Print("fail.")
		return true
	} else if gDebug {
		fmt.Println("testSchema(" + modelPrivPath + ") : false. printing schema...")
		fmt.Println(string(jsonBody))
		fmt.Println("failed.")
		return false
	}
	fmt.Print("pass.")
	return false
}

// (ultra) *selfExplanatory
func testView(hasToPass bool) bool {
	fmt.Println("testing the View.........test not written.")
	return true
}

// superMega selfExplanatory
func testController(hasToPass bool) bool {
	fmt.Println("testing the Controller!..test not written.")
	return true
}
