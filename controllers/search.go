package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aabukhader/backEnd/models"
	"github.com/gorilla/mux"
)

func startSearch(query string) ([]*models.PostItem, error) {
	resp, err := http.Get("https://demo.dataverse.org/api/search?q=" + query)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var postList *models.SearchBody
	json.Unmarshal(bodyBytes, &postList)

	defer resp.Body.Close()
	return postList.Data.Items, err

}

// Search function
func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	result, err := startSearch(params["query"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		var res models.StatusRes
		res.Status = 500
		res.Msg = "Something went wrong"
		json.NewEncoder(w).Encode(res)
	} else {
		var res models.PostItemResSuccss
		res.Status = 200
		res.Msg = "success"
		res.Data = result
		json.NewEncoder(w).Encode(res)
	}

}
