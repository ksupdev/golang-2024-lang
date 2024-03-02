package handler

import (
	"bank/errs"
	"fmt"
	"net/http"
)

func handleError(w http.ResponseWriter, err error) {
	switch v := err.(type) {
	case errs.AppError:
		w.WriteHeader(v.Code)
		fmt.Fprintln(w, v.Message)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
	}
}
