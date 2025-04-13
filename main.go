package main

import (
	"cv-api/internal/handler"
	"cv-api/internal/repository"
	"cv-api/internal/usecase"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Domain struct {
		Name string `yaml:"name"`
		SSL  bool   `yaml:"ssl"`
	} `yaml:"domain"`
	Paths struct {
		Static    string `yaml:"static"`
		Templates string `yaml:"templates"`
	} `yaml:"paths"`
}

func main() {
	// Chargement de la configuration
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Erreur lors de la lecture du fichier de configuration: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("Erreur lors du parsing de la configuration: %v", err)
	}

	// Initialisation des composants
	cvRepo := repository.NewCVRepository()
	cvUsecase := usecase.NewCVUsecase(cvRepo)
	cvHandler := handler.NewCVHandler(cvUsecase)

	// Configuration des routes
	http.HandleFunc("/", cvHandler.ServeHTML)
	http.HandleFunc("/cv", cvHandler.GetCV)     // Route pour le JavaScript
	http.HandleFunc("/api/cv", cvHandler.GetCV) // Route API alternative

	// Configuration du serveur de fichiers statiques
	fs := http.FileServer(http.Dir(config.Paths.Static))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Configuration de robots.txt
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("User-agent: *\nAllow: /\nSitemap: https://" + config.Domain.Name + "/sitemap.xml"))
	})

	// Configuration de sitemap.xml
	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
   <url>
      <loc>https://` + config.Domain.Name + `/</loc>
      <lastmod>2024-04-13</lastmod>
      <changefreq>monthly</changefreq>
      <priority>1.0</priority>
   </url>
</urlset>`))
	})

	// Démarrage du serveur
	addr := config.Server.Host + ":" + strconv.Itoa(config.Server.Port)
	log.Printf("Serveur démarré sur %s", addr)
	log.Printf("Domaine configuré: %s (SSL: %v)", config.Domain.Name, config.Domain.SSL)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur: %v", err)
	}
}
