package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

func listTasksHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, height, err := clientCtx.QueryWithData("/custom/taskbounty/tasks", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func getTaskHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskID := vars["id"]

		params := []byte(fmt.Sprintf(`{"task_id":"%s"}`, taskID))
		res, height, err := clientCtx.QueryWithData("/custom/taskbounty/task", params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func getBalanceHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		address := vars["address"]

		params := []byte(fmt.Sprintf(`{"address":"%s"}`, address))
		res, height, err := clientCtx.QueryWithData("/custom/taskbounty/balance", params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}
