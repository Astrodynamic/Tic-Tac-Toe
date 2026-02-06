package di

import (
	"context"
	"net/http"
	"tictactoe/internal/repository"
	"tictactoe/internal/service"
	transport "tictactoe/internal/transport/http"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("tictactoe",
		fx.Provide(
			repository.NewStorage,
			repository.NewRepository,
			service.NewService,
			transport.NewHandler,
			func(h *transport.Handler) *http.Server {
				return &http.Server{Addr: ":8080", Handler: h}
			},
		),
		fx.Invoke(func(lc fx.Lifecycle, srv *http.Server) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go srv.ListenAndServe()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return srv.Shutdown(ctx)
				},
			})
		}),
	)
}
