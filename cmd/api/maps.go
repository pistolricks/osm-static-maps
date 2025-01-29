package main

import (
	"net/http"
	"osm-static-maps/internal/data"
)

func (app *application) locationMapHandler(w http.ResponseWriter, r *http.Request) {
	geo := &data.Geo{
		Title:    "Home",
		FileName: "map-check-1",
		Lat:      40.641766,
		Lng:      -73.780968,
	}

	app.background(func() {
		data.LocationMap(geo)
	})

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"title":    geo.Title,
			"filename": geo.FileName,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		env := envelope{"error": err.Error()}

		err := app.writeJSON(w, 500, env, nil)
		if err != nil {
			w.WriteHeader(500)
		}
	}
}
