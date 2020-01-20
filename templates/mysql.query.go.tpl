{{- $short := (shortname .Type.Name "err" "sqlstr" "db" "q" "res" .QueryParams) -}}
{{- $queryComments := .QueryComments -}}
{{- if .Comment -}}
// {{ .Comment }}
{{- else -}}
// {{ .Name }} runs a custom query, returning results as {{ .Type.Name }}.
{{- end }}
func {{ .Name }} (ctx context.Context, {{ range .QueryParams }} {{ .Name }} {{ .Type }}{{ end }}) ({{ if not .OnlyOne }}[]{{ end }}*{{ .Type.Name }}, error) {
	var err error
	var dbConn *sql.DB

	// sql query
	{{ if .Interpolate }}var{{ else }}const{{ end }} sqlstr = {{ range $i, $l := .Query }}{{ if $i }} +{{ end }}{{ if (index $queryComments $i) }} // {{ index $queryComments $i }}{{ end }}{{ if $i }}
	{{end -}}`{{ $l }}`{{ end }}

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr{{ range .QueryParams }}{{ if not .Interpolate }}, {{ .Name }}{{ end }}{{ end }})))

    tx, err := components.M.GetConnFromCtx(ctx)
    if err != nil {
       dbConn, err = components.M.GetSlaveConn()
       if err != nil {
           return nil, err
       }
    }

{{- if .OnlyOne }}
	var {{ $short }} {{ .Type.Name }}
	if (tx != nil) {
        err = tx.QueryRow(sqlstr{{ range .QueryParams }}, {{ .Name }}{{ end }}).Scan({{ fieldnames .Type.Fields (print "&" $short) }})
        if err != nil {
            return nil, err
        }
	} else {
	    err = m.QueryRow(sqlstr{{ range .QueryParams }}, {{ .Name }}{{ end }}).Scan({{ fieldnames .Type.Fields (print "&" $short) }})
    	if err != nil {
    		return nil, err
    	}
	}

	return &{{ $short }}, nil
{{- else }}
    var q *sql.Rows
    if (tx != nil) {
        q, err = tx.Query(sqlstr{{ range .QueryParams }}, {{ .Name }}{{ end }})
        if err != nil {
        	return nil, err
        }
	} else {
	    q, err = m.Query(sqlstr{{ range .QueryParams }}, {{ .Name }}{{ end }})
        if err != nil {
        	return nil, err
        }
        defer m.Close()
	}

	// load results
	res := make([]*{{ .Type.Name }}, 0)
	for q.Next() {
		{{ $short }} := {{ .Type.Name }}{}

		// scan
		err = q.Scan({{ fieldnames .Type.Fields (print "&" $short) }})
		if err != nil {
			return nil, err
		}

		res = append(res, &{{ $short }})
	}

	return res, nil
{{- end }}
}

