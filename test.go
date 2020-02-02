package main

//this has a self-explanatory name
func testEverything(hasToPass bool) {
	passM := testModel(hasToPass)
	passV := testView(hasToPass)
	passC := testController(hasToPass)
	if passM && passV && passC {
		print("pass.\n")
	} else if !passM {
		print("failed testModel()......\n")
	} else if !passV {
		print("failed testView()\n")
	} else if !passC {
		print("failed the testController()\n")
	}
}

// self-explanatory
func testModel(hasToPass bool) bool {
	print("Testing the Model........\n")
	testModelPath := modelSchemaPub + "/" + schemaFolder + modelMarkup
	print("Testing model path: " + testModelPath + "\n")
	schema, err := loadStaticBody(testModelPath)
	if !gDebug && err {
		print("schema.json:" + string(schema) + "\n")
		return testSchema(string(schema))
	} else if !gDebug {
		if !err {
			print("Could not retrieve JSON schema!\n")
			return false
		}
	}
	return true
}

// ...
func testSchema(schema string) bool {
	print("Testing the Public Schema.......\n")
	jsonBody, err := loadStaticBody(modelPrivPath)
	if err {
		print("Fail.\n")
		return true
	} else if gDebug {
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
	print("Testing the View.........test not written.\n")
	return true
}

// superMega selfExplanatory
func testController(hasToPass bool) bool {
	print("testing the Controller!..test not written.\n")
	return true
}
