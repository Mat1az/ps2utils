package models

type ZSOHeader struct {
	IsZSO   bool   `json:"is_zso"`
	Magic   string `json:"magic"`
	HSize   int    `json:"header_size"`
	OSize   string `json:"orig_size"`
	BS      int    `json:"block_size"`
	Version int    `json:"version"`
	IS      int    `json:"index_shift"`
	Unused  string `json:"unused"`
}
