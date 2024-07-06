package types

type ResponseType struct {
	Status  bool         `json:"status"`
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *[]ComicType `json:"data"`
}
