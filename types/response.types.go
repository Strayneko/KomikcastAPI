package types

type ResponseType struct {
	Status      bool   `json:"status"`
	Code        int16  `json:"code"`
	Message     string `json:"message"`
	Total       int16  `json:"total,omitempty"`
	CurrentPage int16  `json:"current_page,omitempty"`
	PrevPage    int16  `json:"prev_page,omitempty"`
	NextPage    int16  `json:"next_page,omitempty"`
	LastPage    int16  `json:"last_page,omitempty"`
	Data        any    `json:"data,omitempty"`
}
