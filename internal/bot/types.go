package bot

type Updates struct {
	ID      int    `json:"update_id"`
	Message string `jason:"Messaage"`
}

type UpdateResponse struct {
	Ok     bool      `json:"ok"`
	Result []Updates `json:"result"`
}
