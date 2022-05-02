package handler

import (
	"etcd/service"
	"fmt"
	"net/http"
)

type GetServeHandler struct{}

func (h *GetServeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reg_ser := r.Header.Get("service")
	fmt.Println("[/get-serve] service: ", reg_ser)

	address := service.NewService().FindServer(reg_ser)
	if address == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("Get nothing")
		return
	}
	fmt.Println(address)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(address))
	return
}
