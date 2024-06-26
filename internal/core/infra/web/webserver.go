package web

import (
	_ "github.com/beriloqueiroz/desafio-dev-back/docs/core/swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

//	@title			Swagger Desafio Meli API
//	@version		1.0
//	@description	This is a notification server .
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	berilo.queiroz@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

type HandlerFuncMethod struct {
	HandleFunc http.HandlerFunc
	Method     string
}

type WebServer struct {
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddRoute(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) Start() error {
	mux := http.NewServeMux()
	s.AddRoute("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))
	for path, handler := range s.Handlers {
		mux.Handle(path, handler)
	}
	return http.ListenAndServe(s.WebServerPort, mux)
}

type output struct {
	Message string
}
