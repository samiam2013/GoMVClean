package main

//this has a self-explanatory name
func testEverything(hasToPass bool) {
	passM := testModel(hasToPass)
	passV := testView(hasToPass)
	passC := testController(hasToPass)
	if passM && passV && passC {
		print("Tests Pass.\n")
	} else if !passM {
		print("Failed testModel()......\n")
	} else if !passV {
		print("Failed testView()\n")
	} else if !passC {
		print("Failed the testController()\n")
	}
}

// self-explanatory
func testModel(hasToPass bool) bool {
	testModelPath := modelSchemaPub + "/" + schemaFolder + modelMarkup
	schema, err := loadStaticBody(testModelPath)
	if !globalDebug && err {
		print("schema.json:" + string(schema) + "\n")
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
	print("\nTesting model*.go(?).......\n")
	jsonBody, err := loadStaticBody(modelPrivPath)
	if err {
		print("Fail.\n")
		return true
	} else if globalDebug {
		print("testSchema(" + modelPrivPath + ") : false. printing schema...\n\n")
		print(string(jsonBody))
		print("Failed.\n")
		return false
	}
	print("pass.\n")
	return false
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
