package templates

import (
	"bytes"
	"text/template"
)

var templateGetAssets = `
'{{.ReadToken}}' 'token' STORE
NOW 'end' STORE
[ $token 'asset' {} $end NOW ] FETCH
[ SWAP bucketizer.sum $end 0 1 ] BUCKETIZE
<%
    DUP LABELS 'name' GET 
    SWAP VALUES 0 GET
%> FOREACH

DEPTH ->MAP
`

var templateGetMoney = `
'{{.ReadToken}}' 'token' STORE
NOW 'end' STORE
[ $token 'money' {} $end -1 ] FETCH
0 GET VALUES 0 GET`

type templateGetAssetsData struct {
	ReadToken string
}

// GetAssets is creating the right MC2 to get all Assets and their number
func GetAssets(token string) string {
	tmpl, err := template.New("getAssets").Parse(templateGetAssets)
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, templateGetAssetsData{ReadToken: token})

	return tpl.String()
}

// GetMoney is creating the right MC2 to get all Assets and their number
func GetMoney(token string) string {
	tmpl, err := template.New("getMoney").Parse(templateGetMoney)
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, templateGetAssetsData{ReadToken: token})

	return tpl.String()
}
