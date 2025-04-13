# API CV en Go

Une API simple en Go pour afficher un CV en ligne.

## Installation

1. Assurez-vous d'avoir Go installé sur votre machine
2. Clonez ce dépôt
3. Exécutez `go mod tidy` pour installer les dépendances

## Utilisation

1. Lancez le serveur avec la commande :
```bash
go run main.go
```

2. Accédez à l'API via :
```
http://localhost:8080/cv
```

## Structure du CV

Le CV est retourné au format JSON avec les champs suivants :
- nom
- prenom
- email
- telephone
- poste
- description
- experiences
- formations
- competences

## Modification du CV

Pour modifier les informations du CV, éditez le fichier `main.go` et modifiez la structure `CV` dans la fonction `handleCV`. 