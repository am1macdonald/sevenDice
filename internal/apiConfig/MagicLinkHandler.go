package apiConfig

//func (cfg *apiConfig) HandleMagicLink(w http.ResponseWriter, r *http.Request) {
//	t := r.URL.Query().Get("token")
//	if t == "" {
//		errorResponse(errors.New("token is required"), 404, w)
//		return
//	}
//
//	_, err := magiclink.ValidateMagicLink(t, *cfg.valKey, r.Context())
//	if err != nil {
//		errorResponse(errors.New("token is invalid"), 404, w)
//		return
//	}
//
//}
