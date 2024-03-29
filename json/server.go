package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
 具有通过在结构或接口中嵌入类型来借用一部分实现典型的，类型驱动的子类化概念的能力
 可以通过嵌入将接口组成新的接口，结构也可以如此
 同时，注意将会公开所有嵌入类型的公共方法和字段
 滥用嵌入将会污染你的 API 并暴露你的类型的内部信息
 路由 完全支持 http.Handler 接口，因为可以将路由分配给 Handler，路由本身也是 Handler

*/

const jsonContentType = "application/json"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type Player struct {
	Name string
	Wins int
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	// not initialized
	//p := &PlayerServer{
	//	store,
	//	http.NewServeMux(),
	//}

	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	p.Handler = router

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	// 要创建一个 Encoder，需要一个 http.ResponseWriter 实现的 io.Writer
	// 要创建一个 Decoder，需要一个 io.Writer，由我们的响应 Body 字段实现
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
