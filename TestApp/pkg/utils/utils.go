package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//ParseBody Json decode işlemini döner
func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	//readall methodu başarılı ise err nill döner
	if err == nil {
		err := json.Unmarshal([]byte(body), x)
		if err != nil {
			return
		}
	}
}
