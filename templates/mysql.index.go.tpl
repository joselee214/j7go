{{- $short := (shortname .Type.Name "err" "sqlstr" "q" "res" .Fields) -}}
{{- $table := (schema .Schema .Type.Table.TableName) -}}
// {{ .FuncName }} retrieves a row from '{{ $table }}' as a {{ .Type.Name }}.
//
// Generated from index '{{ .Index.IndexName }}'.
func {{ .FuncName }}(ctx context.Context, {{ goparamlist .Fields false true }}, key... interface{}) ({{ if not .Index.IsUnique }}[]{{ end }}*{{ .Type.Name }}, error) {
	var err error
	var dbConn *sql.DB

    tableName, err := Get{{  .Type.Name  }}TableName(key...)
    if err != nil {
        return nil, err
    }

	// sql query
	sqlstr := `SELECT ` +
		`{{ colnames .Type.Fields }} ` +
		`FROM ` + tableName +
		` WHERE {{ colnamesquery .Fields " AND " }}`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr{{ goparamlist .Fields true false }})))

    tx, err := components.M.GetConnFromCtx(ctx)
    if err != nil {
       dbConn, err = components.M.GetSlaveConn()
       if err != nil {
           return nil, err
       }
    }

{{- if .Index.IsUnique }}
	{{ $short }} := {{ .Type.Name }}{
	{{- if .Type.PrimaryKey }}
		_exists: true,
	{{ end -}}
	}

	if (tx != nil) {
        err = tx.QueryRow(sqlstr{{ goparamlist .Fields true false }}).Scan({{ fieldnames .Type.Fields (print "&" $short) }})
    	if err != nil {
    		return nil, err
    	}
	} else {
	    err = dbConn.QueryRow(sqlstr{{ goparamlist .Fields true false }}).Scan({{ fieldnames .Type.Fields (print "&" $short) }})
    	if err != nil {
    		return nil, err
    	}
	}

	return &{{ $short }}, nil
{{- else }}
    var queryData *sql.Rows
	if (tx != nil) {
        queryData, err = tx.Query(sqlstr{{ goparamlist .Fields true false }})
    	if err != nil {
    		return nil, err
    	}
	} else {
	    queryData, err = dbConn.Query(sqlstr{{ goparamlist .Fields true false }})
    	if err != nil {
    		return nil, err
    	}
	}

	defer queryData.Close()

	// load results
	res := make([]*{{ .Type.Name }}, 0)
	for queryData.Next() {
		{{ $short }} := {{ .Type.Name }}{
		{{- if .Type.PrimaryKey }}
			_exists: true,
		{{ end -}}
		}

		// scan
		err = queryData.Scan({{ fieldnames .Type.Fields (print "&" $short) }})
		if err != nil {
			return nil, err
		}

		res = append(res, &{{ $short }})
	}

	return res, nil
{{- end }}
}

