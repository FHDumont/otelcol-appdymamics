package expressions

import (
	"fmt"

	"github.com/antchfx/jsonquery"
	"github.com/chrlic/otelcol-cust/collector/shared/contextdb"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"
)

func (c *ExpressionEnvironment) SetContextDB(db *contextdb.ContextDb) {
	c.db = db
}

func (c *ExpressionEnvironment) dbFunctions() []cel.EnvOption {
	functions := []cel.EnvOption{}

	var dbGetAllRecordsFunctionImpl = cel.FunctionBinding(func(args ...ref.Val) ref.Val {

		jsonQuery, ok := args[0].Value().(string)
		if !ok {
			return types.NewErr("invalid operand of type '%v' - a string", args[0].Type())
		}
		table, ok := args[1].Value().(string)
		if !ok {
			return types.NewErr("invalid operand of type '%v' - a string", args[1].Type())
		}
		index, ok := args[2].Value().(string)
		if !ok {
			return types.NewErr("invalid operand of type '%v' - a string", args[2].Type())
		}
		fieldsIn, ok := args[3].(traits.Lister)
		if !ok {
			return types.NewErr("invalid operand of type '%v' - should be list of strings", args[3].Type())
		}

		fields := []string{}
		iter := fieldsIn.Iterator()
		for iter.HasNext().Value().(bool) {
			i := iter.Next()
			field := i.Value().(string)
			fields = append(fields, field)
		}

		records, err := c.db.GetAllRecords(table, index, fields...)
		if err != nil {
			return types.NewErr("cannot get db data - table %s, index %s, fields %v - %v", table, index, fields, err)
		}

		values := []string{}

		for _, rec := range records {
			jsonqueryNode := rec.Data
			valSlicePtr := jsonquery.Find(jsonqueryNode, jsonQuery)
			for _, valPtr := range valSlicePtr {
				values = append(values, fmt.Sprintf("%s", valPtr.Value()))
			}
		}

		return types.NewStringList(AnyAdapter{}, values)
	})

	var dbGetFirstRecordFunctionImpl = cel.FunctionBinding(func(args ...ref.Val) ref.Val {

		jsonQuery, ok := args[0].Value().(string)
		if !ok {
			return types.NewErr("invalid operand of type '%v' - a string", args[0].Type())
		}
		table, ok := args[1].Value().(string)
		if !ok {
			return types.NewErr("invalid operand of type '%v' - a string", args[1].Type())
		}
		index, ok := args[2].Value().(string)
		if !ok {
			return types.NewErr("invalid operand of type '%v' - a string", args[2].Type())
		}
		fieldsIn, ok := args[3].(traits.Lister)
		if !ok {
			return types.NewErr("invalid operand of type '%v' - should be list of strings", args[3].Type())
		}

		fields := []string{}
		iter := fieldsIn.Iterator()
		for iter.HasNext().Value().(bool) {
			i := iter.Next()
			field := i.Value().(string)
			fields = append(fields, field)
		}

		record, err := c.db.GetOneRecord(table, index, fields...)
		if err != nil {
			return types.NewErr("cannot get db data - table %s, index %s, fields %v - %v", table, index, fields, err)
		}

		valPtr := jsonquery.FindOne(record.Data, jsonQuery)
		valueStr := fmt.Sprintf("%s", valPtr.Value())

		return types.String(valueStr)

	})

	var dbGetAllRecords = cel.Function("dbGetAll",
		cel.Overload("dbGetAll_string_string_string_list", // table, index, fields...
			[]*cel.Type{cel.StringType, cel.StringType, cel.ListType(cel.StringType)},
			cel.ListType(cel.AnyType),
			dbGetAllRecordsFunctionImpl,
		),
	)

	var dbGetFirstRecord = cel.Function("dbGetFirst",
		cel.Overload("dbGetFirst_string_string_string_list", // table, index, fields...
			[]*cel.Type{cel.StringType, cel.StringType, cel.ListType(cel.StringType)},
			cel.ListType(cel.AnyType),
			dbGetFirstRecordFunctionImpl,
		),
	)

	functions = append(functions, dbGetAllRecords)
	functions = append(functions, dbGetFirstRecord)

	return functions
}
