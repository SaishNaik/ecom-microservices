package proxy

import (
	"context"
	"fmt"
	"github.com/SaishNaik/ecom_microservices/product/config"
	gen "github.com/SaishNaik/ecom_microservices/product/proto/gen/go/proto"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func NewHttpGRPCGateway(ctx context.Context, config *config.Config, opts []gwruntime.ServeMuxOption) (http.Handler, error) {
	productEndpoint := fmt.Sprintf("%s:%s", config.GRPC.Host, config.GRPC.Port)
	mux := gwruntime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gen.RegisterProductServiceHandlerFromEndpoint(ctx, mux, productEndpoint, dialOpts)
	return mux, err
}

//func allowCORS(h http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if origin := r.Header.Get("Origin"); origin != "" {
//			w.Header().Set("Access-Control-Allow-Origin", origin)
//			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
//				preflightHandler(w, r)
//
//				return
//			}
//		}
//		h.ServeHTTP(w, r)
//	})
//}

//func preflightHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//
//	headers := []string{"*"}
//	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
//
//	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
//	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
//
//	slog.Info("preflight request", "http_path", r.URL.Path)
//}

//func withLogger(h http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		slog.Info("Run request", "http_method", r.Method, "http_url", r.URL)
//
//		h.ServeHTTP(w, r)
//	})
//}

//func main() {
//	ctx := context.Background()
//
//	ctx, cancel := context.WithCancel(ctx)
//	defer cancel()
//
//	cfg, err := config.NewConfig()
//	if err != nil {
//		glog.Fatalf("Config error: %s", err)
//	}
//
//	// set up logrus
//	logrus.SetFormatter(&logrus.JSONFormatter{})
//	logrus.SetOutput(os.Stdout)
//	logrus.SetLevel(logger.ConvertLogLevel(cfg.Log.Level))
//
//	// integrate Logrus with the slog logger
//	slog.New(logger.NewLogrusHandler(logrus.StandardLogger()))
//
//	mux := http.NewServeMux()
//
//	gw, err := newGateway(ctx, cfg, nil)
//	if err != nil {
//		slog.Error("failed to create a new gateway", err)
//	}
//
//	mux.Handle("/", gw)
//
//	s := &http.Server{
//		Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
//		Handler: allowCORS(withLogger(mux)),
//	}
//
//	go func() {
//		<-ctx.Done()
//		slog.Info("shutting down the http server")
//
//		if err := s.Shutdown(context.Background()); err != nil {
//			slog.Error("failed to shutdown http server", err)
//		}
//	}()
//
//	slog.Info("start listening...", "address", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
//
//	if err := s.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
//		slog.Error("failed to listen and serve", err)
//	}
//}
