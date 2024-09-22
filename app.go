package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type App struct {
	ctx context.Context
}

type Game struct {
	Format int       `json:"format"`
	Url    string    `json:"url"`
	Id     string    `json:"id"`
	Name   string    `json:"name"`
	Size   string    `json:"size"`
	OPL    bool      `json:"opl"`
	HDL    bool      `json:"hdl"`
	ZSO    ZSOHeader `json:"zso"`
}

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

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetGame(path string) Game {
	f, _ := os.Open(path)
	zso := getZSOHeader(f)
	isZSO := (ZSOHeader{}) != zso
	stat, _ := f.Stat()
	var fSize = formatSize(stat.Size())
	var name = getName(stat.Name())
	var format = getFormat(stat.Size(), isZSO)
	g := Game{
		format,
		f.Name(),
		name[0],
		name[1],
		fSize,
		isOPL(filepath.Ext(name[1])),
		isHDL(int(stat.Size())),
		zso}
	f.Close()
	return g
}

func (a *App) SelectFile() []Game {
	//FIXME Linux: App stop working/freeze after closing/cancel this dialog
	var files, _ = runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{})
	var games []Game
	for _, s := range files {
		games = append(games, a.GetGame(s))
	}
	return games
}

func (a *App) RepairFile(path string, id int) string {
	switch id {
	case 0:
		//missing ID
		/** ISO Priority
		 * 1. Extract ID
		 * 2. Checksum
		 * 3. Hex Find
		 *
		 * ZSO Priority
		 * 1. ZSO -> ISO
		 */
		return "Soon..."
	case 1:
		//opl
		return "Soon..."
	case 2:
		//hdl
		f, _ := os.OpenFile(path, os.O_RDWR|os.O_APPEND, os.ModeAppend)
		stat, _ := f.Stat()
		add := make([]byte, 2048-stat.Size()%2048)
		f.Write(add)
		f.Close()
		return strconv.Itoa(len(add))
	case 3:
		//hdl, ZSO without ID
		//same as case 0, ZSO
		return "Soon..."
	default:
		return "Soon..."
	}
}

func (a *App) GetHDL(path string) string {
	g := a.GetGame(path)
	//inject_cd/dvd ps2hdd_path game_name game_path game_startup
	cmd := fmt.Sprintf("sudo hdl-dump %s %s '%s' '%s' '%s'",
		map[bool]string{true: "inject_cd", false: "inject_dvd"}[g.Format == 1],
		"/dev/sdx",
		g.Name,
		path,
		g.Id)
	fmt.Println(cmd)
	return cmd
}

func getFormat(s int64, isZSO bool) int {
	//0: ISO DVD; 1: ISO CD; 2: ZSO; 3: ZSO DVD; 4: ZSO CD
	switch {
	case s%2048 == 0:
		// ISO DVD or ZSO DVD
		return map[bool]int{true: 3, false: 0}[isZSO]
	case s%2352 == 0:
		// ISO CD or ZSO CD
		return map[bool]int{true: 4, false: 1}[isZSO]
	default:
		// ZSO or Invalid
		return map[bool]int{true: 2, false: -1}[isZSO]
	}
}

func getZSOHeader(f *os.File) ZSOHeader {
	h := make([]byte, 24)
	f.Read(h)
	magic := h[0:4]
	if strings.EqualFold(string(magic), "ziso") {
		var ver byte
		var is byte
		hsize := h[4:8]
		osize := h[8:16]
		bs := h[16:20]
		ver = h[20]
		is = h[21]
		unused := h[21:23]
		return ZSOHeader{true, string(magic),
			int(binary.LittleEndian.Uint32(hsize)),
			formatSize(int64(binary.LittleEndian.Uint64(osize))),
			int(binary.LittleEndian.Uint32(bs)),
			int(int8(ver)),
			int(int8(is)),
			string(unused),
		}
	} else {
		return ZSOHeader{}
	}
}

func getName(s string) []string {
	r := regexp.MustCompile("(\\w{4}_\\d{3}\\.\\d{2})\\.(.+)").FindStringSubmatch(s)
	if len(r) > 2 {
		//XXXX_XX.XX.X
		return []string{r[1], r[2]}
	} else {
		//filename doesn't contains id
		return []string{"", s}
	}
}

func formatSize(s int64) string {
	if s < 1024*1024*1024 {
		return fmt.Sprintf("%.2f %s", float64(s)/1024/1024, "MB")
	} else {
		return fmt.Sprintf("%.2f %s", float64(s)/1024/1024/1024, "GB")
	}
}

func isOPL(s string) bool {
	return strings.EqualFold(s, ".zso") || strings.EqualFold(s, ".iso")
}

func isHDL(s int) bool {
	return s%2048 == 0 || s%2352 == 0
}
