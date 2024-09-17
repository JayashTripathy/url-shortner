package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		dest, pathFound :=  pathToUrls[path];
		if  pathFound {
			fmt.Println("redirecting to ", dest)
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

type pathUrl struct {
	Path string `json:"path"`
	URL string `json:"url"`
}

func jsonHandler(jsonData []byte, fallback http.Handler, defaultPaths ...map[string]string) http.HandlerFunc  {
	
	pathUrls, err := parseJSON(jsonData)

	if(err != nil) {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	extractedPathToUrls := buildMap(pathUrls)
	if len(defaultPaths) > 0 && defaultPaths[0] != nil {		
		for path, url := range defaultPaths[0] {
			extractedPathToUrls[path] = url
		}
	}

	return MapHandler(extractedPathToUrls, fallback)

}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathStruc := map[string]string{}

	for _, pathUrl := range pathUrls{
		pathStruc[pathUrl.Path] = pathUrl.URL
	}

	return pathStruc
}


func parseJSON(jsonData []byte) ([]pathUrl, error) {
	pathUrls := []pathUrl{}
	err := json.Unmarshal(jsonData, &pathUrls)

	if(err != nil) {
		return nil ,err
	}

	return pathUrls, nil
}