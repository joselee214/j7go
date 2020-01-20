{{- $notVoid := (ne .Proc.ReturnType "void") -}}
{{- $proc := (schema .Schema .Proc.ProcName) -}}
{{- if ne .Proc.ReturnType "trigger" -}}
// {{ .Name }} calls the stored procedure '{{ $proc }}({{ .ProcParams }}) {{ .Proc.ReturnType }}' on db.
func {{ .Name }}(ctx context.Context, {{ goparamlist .Params true true }}) ({{ if $notVoid }}{{ retype .Return.Type }}, {{ end }}error) {
	var err error
	var dbConn *sql.DB

	// sql query
	const sqlstr = `SELECT {{ $proc }}({{ colvals .Params }})`

	// run query
	tx, err = components.M.GetConnFromCtx(ctx)
	if err != nil {
	    dbConn, err = components.M.GetSlaveConn()
        if err != nil {
            return nil, err
        }
	}
{{- if $notVoid }}
	var ret {{ retype .Return.Type }}
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr{{ goparamlist .Params true false }})))
	if (tx != nil) {
        err = tx.QueryRow(sqlstr{{ goparamlist .Params true false }}).Scan(&ret)
        if err != nil {
            return {{ reniltype .Return.NilType }}, err
        }
	} else {
	    err = dbConn.QueryRow(sqlstr{{ goparamlist .Params true false }}).Scan(&ret)
        if err != nil {
            return {{ reniltype .Return.NilType }}, err
        }
	}

	return ret, nil
{{- else }}
	components.L.Debug("DB", zap.String("SQL", sqlstr)))
	_, err = m.Exec(sqlstr)
	if (tx != nil) {
	    _, err = tx.Exec(sqlstr)
	} else {
	    _, err = dbConn.Exec(sqlstr)
	}
	return err
{{- end }}
}
{{- end }}

