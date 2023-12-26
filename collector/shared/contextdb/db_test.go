package contextdb

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/antchfx/jsonquery"
	"go.uber.org/zap"
)

const SCHEMA_FILE = "test-db-schema.yaml"

func Test(t *testing.T) {
	logger := &zap.Logger{}
	schemaConfig, err := os.ReadFile(SCHEMA_FILE)
	if err != nil {
		t.Errorf("cannot read db schema yaml %s - %v", SCHEMA_FILE, err)
	}

	dbSchemaAbstract, err := ParseDbJsonSchema(schemaConfig)
	if err != nil {
		t.Errorf("cannot parse db schema file %s - %v", SCHEMA_FILE, err)
	}

	dbSchema, err := GetDbSchema(dbSchemaAbstract)
	if err != nil {
		t.Errorf("cannot convert schema to memdb schema %s - %v", SCHEMA_FILE, err)
	}

	db := ContextDb{}
	err = db.Init(dbSchema, logger)
	if err != nil {
		t.Errorf("cannot init DB %s - %v", SCHEMA_FILE, err)
	}

	t.Logf("Successfully initialized database %v", dbSchemaAbstract)

	for _, a := range appdData {
		jsonDoc, _ := json.Marshal(a)
		jsonQueryDoc, _ := jsonquery.Parse(bytes.NewReader(jsonDoc))
		rec := ContextRecord{Data: jsonQueryDoc}
		err = db.InsertOrUpdateRecord("appd", &rec)
		if err != nil {
			t.Logf("cannot store %s to table %s - %v", jsonDoc, "appd", err)
		}
	}

	sel, err := db.GetAllRecords("appd", "id", "Mockup-Cont", "", "")
	for _, rec := range sel {
		node := jsonquery.FindOne(rec.Data, "/node")
		t.Logf("Rec: %v", node.Value())
	}
}
