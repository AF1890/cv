document.addEventListener('DOMContentLoaded', () => {
    fetch('/cv')
        .then(response => response.json())
        .then(data => {
            // Mise à jour des informations personnelles
            document.getElementById('nom').innerHTML = `${data.prenom} ${data.nom}`;
            document.getElementById('poste').innerHTML = data.poste;
            document.getElementById('description').innerHTML = data.description;

            // Mise à jour des expériences
            const experiencesList = document.getElementById('experiences');
            experiencesList.innerHTML = '';
            data.experiences.forEach(exp => {
                const li = document.createElement('li');
                li.innerHTML = exp;
                experiencesList.appendChild(li);
            });

            // Mise à jour des formations
            const formationsList = document.getElementById('formations');
            formationsList.innerHTML = '';
            data.formations.forEach(formation => {
                const li = document.createElement('li');
                li.innerHTML = formation;
                formationsList.appendChild(li);
            });

            // Mise à jour des compétences techniques
            const competencesList = document.getElementById('competences');
            competencesList.innerHTML = '';
            data.competences.forEach(competence => {
                const li = document.createElement('li');
                li.innerHTML = competence;
                competencesList.appendChild(li);
            });

            // Mise à jour des soft skills
            const softSkillsList = document.getElementById('softSkills');
            softSkillsList.innerHTML = '';
            data.softSkills.forEach(skill => {
                const li = document.createElement('li');
                li.innerHTML = skill;
                softSkillsList.appendChild(li);
            });
        })
        .catch(error => {
            console.error('Erreur lors du chargement du CV:', error);
            document.body.innerHTML = '<div class="error">Erreur lors du chargement du CV. Veuillez réessayer plus tard.</div>';
        });
}); 