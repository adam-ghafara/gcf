package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/petapedia/peda"
)

func init() {
	functions.HTTP("HelloHTTP", HelloHTTP)
}

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	// Set header Access-Control-Allow-Origin untuk mengizinkan permintaan dari domain yang spesifik.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Set header Access-Control-Allow-Methods untuk mengizinkan metode HTTP yang diizinkan.
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	// Set header Access-Control-Allow-Headers untuk mengizinkan header yang diizinkan dalam permintaan.
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		// Jika permintaan adalah preflight OPTIONS, Anda hanya perlu mengirimkan header CORS.
		return
	}

	// Tulis respons Anda ke Writer seperti yang Anda lakukan sebelumnya.
	response := peda.GCFHandler("MONGODATA", "geojsonList", "geojsonList")
	fmt.Fprintf(w, response)
}

func GetToken(r *http.Request) string {
	return r.Header.Get("Authorization")
}
