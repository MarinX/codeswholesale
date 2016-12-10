// util provides utility for common usage
package codeswholesale

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
)

func responseToJSON(resp *http.Response, v interface{}) error {

	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	errJSON := json.Unmarshal(buff, v)

	if err := isCWError(buff); err != nil {
		return err
	}

	return errJSON
}

func isCWError(buff []byte) error {

	cwerr := new(CWError)

	err := json.Unmarshal(buff, cwerr)
	if err == nil {

		if cwerr.Status != API_STATUS_SUCCESS {
			return fmt.Errorf("%s", cwerr.Message)
		}
	}

	return nil
}
