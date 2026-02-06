package http

type Request struct {
	Board [3][3]uint8 `json:"board"`
}

type Response struct {
	UUID   string      `json:"uuid"`
	Board  [3][3]uint8 `json:"board"`
	Winner *uint8      `json:"winner,omitempty"`
}
