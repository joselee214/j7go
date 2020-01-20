{{- $short := (shortname .Name "err" "res" "sqlstr" "db" "XOLog") -}}
{{- $table := (schema .Schema .Table.TableName) -}}
{{- if .Comment -}}
// {{ .Comment }}
{{- else -}}
// {{ .Name }} represents a row from '{{ $table }}'.
{{- end }}
type {{ .Name }} struct {
{{- range .Fields }}
	{{ .Name }} {{ retype .Type }} `json:"{{ .Col.ColumnName }}"` // {{ .Col.ColumnName }}
{{- end }}
{{- if .PrimaryKey }}

	// xo fields
	_exists, _deleted bool
{{ end }}
}

{{ if .PrimaryKey }}
// Exists determines if the {{ .Name }} exists in the database.
func ({{ $short }} *{{ .Name }}) Exists() bool {//{{  .Table.TableName  }}
	return {{ $short }}._exists
}

// Deleted provides information if the {{ .Name }} has been deleted from the database.
func ({{ $short }} *{{ .Name }}) Deleted() bool {
	return {{ $short }}._deleted
}

// Get table name
func Get{{  .Name  }}TableName(key... interface{}) (string, error) {
    tableName, err := components.M.GetTable("{{  .Schema  }}","{{  .Table.TableName  }}", key...)
    if err != nil {
        return "", err
    }
	return tableName, nil
}


// Insert inserts the {{ .Name }} to the database.
func ({{ $short }} *{{ .Name }}) Insert(ctx context.Context, key... interface{}) error {
	var err error
    var dbConn *sql.DB
    var res sql.Result
	// if already exist, bail
	if {{ $short }}._exists {
		return errors.New("insert failed: already exists")
	}

    tx, err := components.M.GetConnFromCtx(ctx)
    if err != nil {
        dbConn, err = components.M.GetMasterConn()
        if err != nil {
       		return err
       	}
    }

    tableName, err := Get{{  .Name  }}TableName(key...)
    if err != nil {
        return err
    }


{{ if .Table.ManualPk  }}
	// sql insert query, primary key must be provided
     sqlstr := `INSERT INTO `+ tableName +
	    ` (` +
		`{{ colnames .Fields }}` +
		`) VALUES (` +
		`{{ colvals .Fields }}` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, {{ fieldnames .Fields $short }})))
	if (tx != nil) {
	    res, err = tx.Exec(sqlstr, {{ fieldnames .Fields $short }})
	} else {
	    res, err = dbConn.Exec(sqlstr, {{ fieldnames .Fields $short }})
	}

	if err != nil {
		return err
	}

	// set existence
	{{ $short }}._exists = true
{{ else }}
	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO `+ tableName +
        ` (` +
		`{{ colnames .Fields .PrimaryKey.Name }}` +
		`) VALUES (` +
		`{{ colvals .Fields .PrimaryKey.Name }}` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }})))
	if err != nil {
		return err
	}
	if (tx != nil) {
    	res, err = tx.Exec(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }})
    } else {
    	res, err = dbConn.Exec(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }})
    }
{{ end }}

    if err != nil {
		return err
	}

    // retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
    {{ $short }}.{{ .PrimaryKey.Name }} = {{ .PrimaryKey.Type }}(id)
    {{ $short }}._exists = true

	return nil
}

{{ if ne (fieldnamesmulti .Fields $short .PrimaryKeyFields) "" }}
	// Update updates the {{ .Name }} in the database.
func ({{ $short }} *{{ .Name }}) Update(ctx context.Context, key... interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if {{ $short }}._deleted {
		return errors.New("update failed: marked for deletion")
	}

    tx, err := components.M.GetConnFromCtx(ctx)
    if err != nil {
        dbConn, err = components.M.GetMasterConn()
        if err != nil {
            return err
        }
    }

    tableName, err := Get{{  .Name  }}TableName(key...)
    if err != nil {
        return err
    }

{{ if gt ( len .PrimaryKeyFields ) 1 }}
    // sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET ` +
			`{{ colnamesquerymulti .Fields ", " 0 .PrimaryKeyFields }}` +
			` WHERE {{ colnamesquery .PrimaryKeyFields " AND " }}`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, {{ fieldnamesmulti .Fields $short .PrimaryKeyFields }}, {{ fieldnames .PrimaryKeyFields $short}})))
	if (tx != nil) {
	    _, err = tx.Exec(sqlstr, {{ fieldnamesmulti .Fields $short .PrimaryKeyFields }}, {{ fieldnames .PrimaryKeyFields $short}})
	} else {
	    _, err = dbConn.Exec(sqlstr, {{ fieldnamesmulti .Fields $short .PrimaryKeyFields }}, {{ fieldnames .PrimaryKeyFields $short}})
	}
	return err
{{- else }}
	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`{{ colnamesquery .Fields ", " .PrimaryKey.Name }}` +
		` WHERE {{ colname .PrimaryKey.Col }} = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }}, {{ $short }}.{{ .PrimaryKey.Name }})))
	if (tx != nil) {
   	    _, err = tx.Exec(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }}, {{ $short }}.{{ .PrimaryKey.Name }})
    } else {
        _, err = dbConn.Exec(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }}, {{ $short }}.{{ .PrimaryKey.Name }})
    }
	return err
{{- end }}
}

// Save saves the {{ .Name }} to the database.
func ({{ $short }} *{{ .Name }}) Save(ctx context.Context) error {
	if {{ $short }}.Exists() {
		return {{ $short }}.Update(ctx)
	}

	return {{ $short }}.Insert(ctx)
}
{{ else }}
	// Update statements omitted due to lack of fields other than primary key
{{ end }}

// Delete deletes the {{ .Name }} from the database.
func ({{ $short }} *{{ .Name }}) Delete(ctx context.Context, key... interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if {{ $short }}._deleted {
		return nil
	}

    tx, err := components.M.GetConnFromCtx(ctx)
    if err != nil {
       dbConn, err = components.M.GetMasterConn()
       if err != nil {
           return err
       }
    }

    tableName, err := Get{{  .Name  }}TableName(key...)
    if err != nil {
        return err
    }
    //{{ len .PrimaryKeyFields  }}

	{{ if eq ( len .PrimaryKeyFields ) 1 }}
		// sql query with composite primary key
		sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE {{ colnamesquery .PrimaryKeyFields " AND " }}`

		// run query
		utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, {{ fieldnames .PrimaryKeyFields $short }})))
		if (tx != nil) {
            _, err = tx.Exec(sqlstr, {{ fieldnames .PrimaryKeyFields $short }})
		} else {
		    _, err = dbConn.Exec(sqlstr, {{ fieldnames .PrimaryKeyFields $short }})
		}

		if err != nil {
			return err
		}
	{{- else }}
		// sql query
		sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE {{ colname .PrimaryKey.Col }} = ?`

		// run query
		utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, {{ $short }}.{{ .PrimaryKey.Name }})))

        if (tx != nil) {
            _, err = tx.Exec(sqlstr, {{ $short }}.{{ .PrimaryKey.Name }})
		} else {
		    _, err = dbConn.Exec(sqlstr, {{ $short }}.{{ .PrimaryKey.Name }})
		}
		if err != nil {
			return err
		}
	{{- end }}

	// set deleted
	{{ $short }}._deleted = true

	return nil
}
{{- end }}

