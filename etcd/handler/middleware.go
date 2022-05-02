package handler

import (
	"fmt"
	"net/http"
)

// assessLog 日志中间件
func assessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("log")
		next.ServeHTTP(w, r)
	})
}
