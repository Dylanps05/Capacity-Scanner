package internal

import (
	"context"
	"log"
	"net/http"

	"github.com/Dylanps05/Capacity-Scanner/internal/logic"
	"github.com/Dylanps05/Capacity-Scanner/internal/mmw"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage"
	"github.com/Dylanps05/Capacity-Scanner/internal/web"
	"github.com/jackc/pgx/v5"
)

type Site struct {
	db *pgx.Conn
	storage.SiteStorage
	logic.Controller
	web.Handler
	mmw.MuxMiddleware
}

func (s *Site) initSQL(db_addr string) {
	conn, err := pgx.Connect(context.Background(), db_addr)
	if err != nil {
		log.Fatal(err)
	}
	s.db = conn
}

func (s *Site) Init(addr string, db_addr string) {
	s.initSQL(db_addr)
	s.SiteStorage = storage.NewDefaultSiteStorage(s.db)
	s.Controller = logic.NewDefaultController(s.SiteStorage)
	s.Handler = web.NewDefaultHandler(s.Controller)
	s.MuxMiddleware = mmw.NewDefaultMuxMiddleware(s.Handler.GetMux(), s.Controller)
	http.ListenAndServe(addr, s.MuxMiddleware.GetMux())
}
