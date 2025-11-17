package internal

import (
	"context"
	"github.com/Dylanps05/Capacity-Scanner/internal/logic"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage"
	"github.com/Dylanps05/Capacity-Scanner/internal/web"
	"github.com/jackc/pgx/v5"
	"log"
)

type Site struct {
	db *pgx.Conn
	storage.SiteStorage
	logic.Controller
	web.Handler
}

func (s *Site) initSQL(db_addr string) {
	conn, err := pgx.Connect(context.Background(), db_addr)
	if err != nil {
		log.Fatal(err)
	}
	s.db = conn
}

func (s *Site) Init(web_addr string, db_addr string) {
	s.initSQL(db_addr)
	s.SiteStorage = storage.NewDefaultSiteStorage(s.db)
	s.Controller = logic.NewDefaultController(s.SiteStorage)
	s.Handler = web.NewDefaultHandler(s.Controller)
	s.Handler.Start(web_addr)
}
