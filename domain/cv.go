package domain

type CV struct {
	Nom         string   `json:"nom"`
	Prenom      string   `json:"prenom"`
	Poste       string   `json:"poste"`
	Description string   `json:"description"`
	Experiences []string `json:"experiences"`
	Formations  []string `json:"formations"`
	Competences []string `json:"competences"`
	SoftSkills  []string `json:"softSkills"`
}
