package pkg

import (
	"fmt"
	"log"
	"net/http"
)

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
