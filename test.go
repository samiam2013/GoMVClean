package main

import "log"

//this has a self-explanatory name
func testEverything(hasToPass bool) {
	print("\nTesting 'Everything'*...")
	passM := testModel(hasToPass)
	passV := testView(hasToPass)
	passC := testController(hasToPass)
	if passM && passV && passC {
		print("Tests Pass.\n")
		return // return so that you don't hit log.Fatal()
	} else if !passM {
		print("Failed testModel()......\n")
	} else if !passV {
		print("Failed testView()\n")
	} else if !passC {
		print("Failed the testController()\n")
	}
	if hasToPass {
		log.Fatal("A test failed and hasToPass (bool) set `true`.")
	}
}

// self-explanatory
func testModel(hasToPass bool) bool {
	print("\nTesting model...")
	testModelPath := modelSchemaPub + "/" + schemaFolder + modelMarkup
	schema, err := loadStaticBody(testModelPath)
	if globalDebug || err {
		//print("schema.json file contents:\n" + string(schema) + "\n\n\n")
		return testSchema(string(schema))
	} else if !globalDebug {
		if !err {
			print("Could not retrieve JSON schema!\n")
			return false
		}
	}
	return true
}

// ...
func testSchema(schema string) bool {
	print("model*.go schema file exists.......")
	if len(schema) == 0 {
		print("Schema file len() == 0...Fail.\n")
		return false
	}
	print("pass.\n")
	return true
}

// (ultra) *selfExplanatory
func testView(hasToPass bool) bool {
	print("Testing view.go...test not written.\n")
	return true
}

// superMega selfExplanatory
func testController(hasToPass bool) bool {
	print("Testing controller.go...test not written.\n")
	return true
}
