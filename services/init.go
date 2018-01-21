package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

var wsCreateGTS = `
'%s' 'wtoken' STORE
'%s' 'rtoken' STORE

<'
%s
'>
JSON-> 

<%%
    PARSESELECTOR 'labels' STORE 'classname' STORE
    [ $rtoken $classname $labels ] FIND SIZE 'size' STORE 
    <%% $size 0 == %%> <%%
        // Creating the GTS if it doesn't exists
        NEWGTS $classname RENAME
        $labels RELABEL
        { '.app' '' '.owner' '' } RELABEL // sanity checks
        DUP LABELS 'name' GET 
        <%%  ISNULL !  %%> <%% DUP LABELS 'name' GET '=' SPLIT 1 GET { SWAP 'name' SWAP  }  RELABEL %%> IFT
        1 NaN NaN NaN 0 ADDVALUE
        $wtoken UPDATE
    %%> IFT
%%> FOREACH`

// CreateGTS is ensuring all the GTS are existing by pushing a 0 in it.
func CreateGTS() error {

	log.Info("Creating assets...")

	var selectors []string

	for assetName, _ := range AllAssetReference {
		selectors = append(selectors, "asset{name="+assetName+"}")
	}
	selectors = append(selectors, "money{}")

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(selectors)
	if err != nil {
		return err
	}

	body := fmt.Sprintf(wsCreateGTS, os.Getenv("WTOKEN"), os.Getenv("RTOKEN"), b.String())

	var resp *http.Response
	resp, err = http.DefaultClient.Post("https://"+os.Getenv("ENDPOINT")+"/api/v0/exec", "", strings.NewReader(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 200 {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatal(err)
		}
		return errors.New(string(dump))
	}

	log.Info("Creating assets is OK")

	return nil

}
