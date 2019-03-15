package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"google.golang.org/grpc"

	aquifers "github.com/stephenhillier/groundwater/aquifers/proto/aquifers"
	wells "github.com/stephenhillier/groundwater/wells/proto/wells"
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

	aq.Wells, err = api.wells.FindWells(context.Background(), &wells.WellSearchRequest{Aquifer: int32(id)})

	if err != nil {
		log.Println(err)
		// todo: decide what else to do with this error (currently, there will simply be no additional wells data attached
		// to the aquifer response)
	}

	render.JSON(w, req, aq)
}

func main() {
	router := chi.NewRouter()

	// Start Aquifers client
	aqConn, err := grpc.Dial("aquifers:7778", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer aqConn.Close()
	aquiferClient := aquifers.NewAquiferServiceClient(aqConn)

	log.Print("Aquifers client ready")

	// Start Wells client
	wellConn, err := grpc.Dial("wells:7777", grpc.WithInsecure())
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
			r.Post("/", api.CreateAquifer)
			r.Get("/{id}", api.GetAquifer)
		})
		r.Route("/wells", func(r chi.Router) {
			r.Post("/", api.wells.CreateWell)
		})
	})

	// Start HTTP API
	log.Print("HTTP API listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", api.router))
}
