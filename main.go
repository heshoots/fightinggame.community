package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

var db *sql.DB

func getSceneGeoJSON(w http.ResponseWriter, r *http.Request) {
	var geojson string
	err := db.QueryRow(`	
	  SELECT json_build_object(
		  'type', 'FeatureCollection',
		  'crs',  json_build_object(
			  'type',      'name',
			  'properties', json_build_object(
				  'name', 'EPSG:4326'
			  )
		  ),
		  'features', json_agg(
			  json_build_object(
				  'type',       'Feature',
				  'id',         id, 
				  'geometry',   ST_AsGeoJSON(location)::json,
				  'properties', json_build_object(
					  -- list of fields
					  'name', name,
					  'facebook', facebook,
					  'twitter', twitter,
					  'discord', discord
				  )
			  )
		  )
	  )
	  FROM (select scenes.id, location, name, facebook, discord, twitter from scenes LEFT JOIN contactpoints on contactpoints.scene_id=scenes.id) as scenes;`).Scan(&geojson)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, geojson)

}

func Logger(name string, route string, handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		end := time.Now()
		log.Println(name, route, end.Sub(start))
	}
}

func main() {
	var err error
	connStr := os.Getenv("DB_CONN")
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/scenes", Logger("getSceneGeoJSON", "/scenes", getSceneGeoJSON))
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}
