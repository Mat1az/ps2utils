package models

type Game struct {
	Format int       `json:"format"`
	Path   string    `json:"path"`
	Id     string    `json:"id"`
	Name   string    `json:"name"`
	Size   string    `json:"size"`
	OPL    bool      `json:"opl"`
	HDL    bool      `json:"hdl"`
	ZSO    ZSOHeader `json:"zso"`
}
