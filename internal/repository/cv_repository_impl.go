package repository

import (
	"cv-api/domain"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type cvRepository struct{}

type PersonalConfig struct {
	Personal struct {
		Nom         string `yaml:"nom"`
		Prenom      string `yaml:"prenom"`
		Poste       string `yaml:"poste"`
		Description string `yaml:"description"`
	} `yaml:"personal"`
}

func NewCVRepository() CVRepository {
	return &cvRepository{}
}

func (r *cvRepository) GetCV() (*domain.CV, error) {
	// Lecture du fichier de configuration
	configFile, err := os.ReadFile("config/personal.yaml")
	if err != nil {
		return nil, err
	}

	var config PersonalConfig
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		return nil, err
	}

	// Calcul de la durée pour Leboncoin
	// Leboncoin
	startDateLeboncoin := time.Date(2021, time.November, 1, 0, 0, 0, 0, time.UTC)
	durationLeboncoin := calculateDuration(startDateLeboncoin)

	return &domain.CV{
		Nom:         config.Personal.Nom,
		Prenom:      config.Personal.Prenom,
		Poste:       config.Personal.Poste,
		Description: config.Personal.Description,
		Experiences: []string{
			"Développeuse backend GO & PHP/Symfony - <a href=https://www.linkedin.com/company/leboncoin/posts/?feedView=all'>Leboncoin</a> · CDI - nov. 2021 - aujourd'hui · " + durationLeboncoin + " - Paris, France",
			"Développeuse web PHP/Symfony - <a href='https://www.linkedin.com/company/mb-diffusion/'>MB Diffusion</a> · CDI - sept. 2019 - oct. 2021 · 2 ans et 2 mois - Evry, France",
			"Développeuse web PHP/Symfony - <a href='https://www.linkedin.com/company/sismic-developpement/posts/?feedView=all'>Sismic, Développement Web</a> · juin 2017 - nov. 2018 · 1 ans et 6 mois - Lieusaint, France",
			"Chargée de mission projets R&D / Digitaux - <a href='https://www.linkedin.com/company/cap-digital/posts/?feedView=all'>Cap Digital</a> · août 2011 - mai 2017 · 5 ans et 9 mois - Paris, France",

			`Pôle de compétitivité Cap Digital, il est le pôle de compétitivité de la transformation numérique.
Missions :
- Gestion des appels à projets de R&D (ANR, FUI, Investissements d'Avenir, BpiFrance, appels régionaux…)
- Participation au jury international Computer Cooking Contest 2012 à Lyon
- Marchés des entreprises soutenues – Mobilité et Smart Cities, e-Education, e-Santé, Médias-Multimédia, Commerce-Distribution…`,
		},
		Formations: []string{
			"<a href='https://www.linkedin.com/school/wild-code-school/posts/?feedView=all'>Wild Code School</a> - Développeur web php/symfony<br><small>· janv 2017 - mai 2017</small>",
			"<a href='https://www.linkedin.com/search/results/all/?keywords=Amiter%20-%20Upec'>Amiter - Upec</a> - Master 2 (M2), Développement de projets territoriaux<br><small>· sept 2009 - juin 2011</small>",
			"<a href='https://www.linkedin.com/search/results/all/?keywords=Amiter%20-%20Upec'>IUT Sénart Fontainebleau</a> - Licence pro (L3), Métiers du Commerce International<br><small>· sept 2008 - juin 2009</small>",
		},
		Competences: []string{
			"Go <span class='skill-gauge'>●●●●○</span><small class='skill-detail'>· Développement backend, API REST, tests unitaires</small>",
			"PHP/Symfony <span class='skill-gauge'>●●●●○</span><small class='skill-detail'>· Connaissance du Framework Symfony 4, développement et maintenance d'applications existantes, bonnes pratiques</small>",
			"PostgreSQL <span class='skill-gauge'>●●●●○</span><small class='skill-detail'>· Conception de schémas, requêtes complexes, optimisation</small>",
			"Git (gerrit/github) <span class='skill-gauge'>●●●●○</span><small class='skill-detail'>· Workflow complet, gestion de branches, résolution de conflits</small>",
			"AWS <span class='skill-gauge'>●●○○○</span><small class='skill-detail'>· Utilisation des services principaux (S3, EventBridge, SNS/SQS, Cognito), monitoring basique</small>",
			"Datadog <span class='skill-gauge'>●●○○○</span><small class='skill-detail'>· Monotoring alertes et d'applications</small>",
			"Docker <span class='skill-gauge'>●○○○○</span><small class='skill-detail'>· Utilisation basique : gestion des conteneurs (start/stop), exécution de commandes dans un conteneur</small>",
			"JavaScript <span class='skill-gauge'>●○○○○</span><small class='skill-detail'>· Notions de base, manipulation simple du DOM (expérience non récente)</small>",
			"HTML/CSS <span class='skill-gauge'>●○○○○</span><small class='skill-detail'>· Notions de base, modifications simples de mise en page et de style</small>",
		},
		SoftSkills: []string{
			"Méthodologies Agiles <span class='skill-gauge'>●●●●●</span><small class='skill-detail'>· Expérience en tant que Scrum Master, utilisation quotidienne de Jira, animation de cérémonies agiles</small>",
			"Collaboration & Communication <span class='skill-gauge'>●●●●○</span><small class='skill-detail'>· Travail en équipe, code reviews, documentation technique, transmission de connaissances</small>",
			"Outils Collaboratifs <span class='skill-gauge'>●●●○○</span><small class='skill-detail'>· Miro pour la création de schémas et diagrammes, Confluence pour la documentation</small>",
			"Organisation <span class='skill-gauge'>●●●●○</span><small class='skill-detail'>· Gestion des priorités, respect des délais, documentation des processus</small>",
		},
	}, nil
}

func calculateDuration(startDate time.Time) string {
	now := time.Now()
	years := now.Year() - startDate.Year()
	months := int(now.Month() - startDate.Month())

	if months < 0 {
		years--
		months += 12
	}

	if years > 0 {
		if months > 0 {
			return fmt.Sprintf("%d ans %d mois", years, months)
		}
		return fmt.Sprintf("%d ans", years)
	}
	return fmt.Sprintf("%d mois", months)
}
