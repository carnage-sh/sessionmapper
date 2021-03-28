package main

import (
	"encoding/json"
	"net/http"

	"github.com/blaqkube/sessionmapper"
)

func ServiceHandler(w http.ResponseWriter, r *http.Request) {
	keys := map[string]string{}
	if r.Body != nil {
		err := json.NewDecoder(r.Body).Decode(&keys)
		if err != nil {
			keys = map[string]string{}
		}
	}
	output := map[string]string{"they": "too"}
	if v, ok := keys["me"]; ok {
		output["you"] = v
	}
	o, _ := json.Marshal(&sessionmapper.Response{Upstream: output})
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(o)
}
