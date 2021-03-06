// Copyright 2020 AWS ElasticRecode Solution Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// @author Su Wei <suwei007@gmail.com>

package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func write(w http.ResponseWriter, data []byte) {

	_, err := w.Write(data)
	if err != nil {
		log.Errorf("Handler: write %s", err.Error())
	}

}

func writeError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	_, writeErr := w.Write([]byte(err.Error()))
	if writeErr != nil {
		log.Error(writeErr)
	}
}

//serverJSON is REST API wrapper
func serverJSON(w http.ResponseWriter, r *http.Request, object interface{}) {
	result, err := json.Marshal(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		write(w, []byte(err.Error()))
		return
	}
	write(w, result)
}
