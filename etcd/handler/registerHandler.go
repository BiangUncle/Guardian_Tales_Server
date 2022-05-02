package handler

import (
	"etcd/service"
	"fmt"
	"net/http"
)

type RegisterHandler struct{}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	address := r.Header.Get("address")
	reg_ser := r.Header.Get("service")
	fmt.Println("[/register] ", address, ", service: ", reg_ser)

	err := service.NewService().Register(address, reg_ser)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ETCD!")
}
