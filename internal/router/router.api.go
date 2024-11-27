package router

import (
	"log"
	"net/http"
)

func (router *Router) handleAPIRoutes(w http.ResponseWriter, r *http.Request) {
	//r.URL.Path = r.URL.Path[1:]
	//
	//page := r.PathValue("page")
	//log.Println(page)
	//if page == "" {
	//	page = "home"
	//}
	//t := router.templateHandler.GetTemplate(page + ".html")
	//err := t.Execute(w, nil)
	//if err != nil {
	//	log.Println(err)
	//	w.WriteHeader(500)
	//	_, err = w.Write([]byte("internal error"))
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	return
	//}
	log.Println("is api route")
	w.WriteHeader(200)
}
