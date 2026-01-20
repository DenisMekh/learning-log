package main

import "context"

type App struct {
	server *Server
}

func New() (*App, error) {
	app := &App{}
	handler := NewHandler()
	router := NewRouter(*handler)
	app.server = NewServer(router)
	return app, nil
}

func (app *App) Start() error {
	return app.server.Start()
}

func (app *App) Stop(ctx context.Context) error {
	if err := app.server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
