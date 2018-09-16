package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	aquifers "github.com/stephenhillier/groundwater/aquifers/proto/aquifers"

	wells "github.com/stephenhillier/groundwater/wells/proto/wells"
	"google.golang.org/grpc"
)

// API represents the server API and holds clients for various services
type API struct {
	router   *chi.Mux
	aquifers aquifers.AquiferServiceClient
	wells    wells.WellServiceClient
}

// GetAquifer retrieves an aquifer from the Aquifers service (by ID)
// and writes it to JSON
func (api *API) GetAquifer(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), 404)
		return
	}

	aq, err := api.aquifers.GetAquifer(context.Background(), &aquifers.SingleAquiferRequest{Id: int32(id)})
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	render.JSON(w, req, aq)
}

func main() {
	router := chi.NewRouter()

	// Start Aquifers client
	aqConn, err := grpc.Dial("127.0.0.1:7778", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer aqConn.Close()
	aquiferClient := aquifers.NewAquiferServiceClient(aqConn)

	log.Print("Aquifers client ready")

	// Start Wells client
	wellConn, err := grpc.Dial("127.0.0.1:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer wellConn.Close()
	wellClient := wells.NewWellServiceClient(wellConn)

	log.Print("Wells client ready")

	// Create an instance of the API with a router and the service clients
	api := &API{
		router,
		aquiferClient,
		wellClient,
	}

	// Set middleware
	api.router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	// Define routes
	api.router.Route("/v1", func(r chi.Router) {
		r.Route("/aquifers", func(r chi.Router) {
			r.Get("/{id}", api.GetAquifer)
		})
	})

	// Start HTTP API
	log.Print("HTTP API listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", api.router))
}
