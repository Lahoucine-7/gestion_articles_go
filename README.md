# 📰 API de Gestion d'Articles

## 🧠 Description du Projet
Ce projet a pour but de développer une API REST complète dédiée à la gestion d'articles, avec documentation interactive via Swagger et tests d'intégration pour garantir la qualité du service.  
Le projet est réalisé dans un but **d'apprentissage** et **n'est pas destiné à un usage en production**.

## 🚀 Contexte d’Utilisation
Le client est une revue en ligne qui modernise sa gestion éditoriale.  
L'API doit :
- 📋 Permettre aux rédacteurs de publier, modifier et supprimer des articles.
- 🛡️ Fournir une interface sécurisée, fiable et documentée pour l'intégration avec des plateformes front-end.

## 📜 Cahier des Charges
- **CRUD complet sur les articles** :
  - Créer (`POST /articles`)
  - Lire tous (`GET /articles`) ou un article (`GET /articles/{id}`)
  - Mettre à jour (`PUT /articles/{id}`)
  - Supprimer (`DELETE /articles/{id}`)
- **Documentation interactive** générée automatiquement avec Swagger.
- **Tests unitaires et d'intégration** couvrant chaque endpoint.
- **Bonne structure** du code : séparation entre modèles, contrôleurs, routes, base de données.

---

## 🛠️ Stack Technique

- **Langage :** Go (≥ 1.17)
- **Framework :** Gin
- **ORM :** GORM
- **Base de Données :** PostgreSQL
- **Documentation API :** Swagger via swaggo
- **Tests :** Go testing + Testify
- **Gestion de config :** Variables d'environnement via godotenv

---

## ⚙️ Installation et Configuration

### 📥 Cloner le Dépôt

```bash
git clone <URL_DU_DEPOT>
cd gestion_articles_go
```

### 📦 Installer les Dépendances

```bash
go mod tidy
```

### 🛡️ Créer un Fichier `.env`
À la racine du projet :

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=ton_mot_de_passe
DB_NAME=gestion_articles
DB_PORT=5432
```

### 🛠️ Générer la Documentation Swagger

```bash
swag init --generalInfo cmd/main.go --output docs
```

Cela génère la documentation Swagger dans le dossier `docs/`.

---

## 🚀 Lancer l'API

```bash
go run cmd/main.go
```

L'API sera disponible sur :  
```text
http://localhost:8080
```

La documentation Swagger sera disponible sur :  
```text
http://localhost:8080/swagger/index.html
```

---

## 🌐 Principaux Endpoints

### 📝 Créer un Article
- **Méthode :** POST
- **URL :** `/articles`
- **Exemple de Payload :**
  ```json
  {
    "title": "Mon premier article",
    "content": "Voici le contenu de l'article.",
    "author": "Auteur Inconnu",
    "published_date": "2025-04-14T00:00:00Z"
  }
  ```
- **Réponse :** Article créé (201 Created)

### 📚 Récupérer Tous les Articles
- **Méthode :** GET
- **URL :** `/articles`
- **Réponse :** Liste d'articles

### 📖 Récupérer un Article par ID
- **Méthode :** GET
- **URL :** `/articles/{id}`
- **Réponse :** Détail d'un article

### ✏️ Mettre à Jour un Article
- **Méthode :** PUT
- **URL :** `/articles/{id}`
- **Payload :** Même structure que POST
- **Réponse :** Article mis à jour

### ❌ Supprimer un Article
- **Méthode :** DELETE
- **URL :** `/articles/{id}`
- **Réponse :** 204 No Content

---

## 🧪 Exécuter les Tests

Tests d'intégration situés dans le dossier `tests/`.

Pour exécuter les tests :

```bash
go test ./tests/...
```

Les tests couvrent :
- Création d'article
- Récupération de tous les articles
- Récupération d'un article spécifique
- Mise à jour d'un article
- Suppression d'un article

---

## 🔮 Suggestions d'Évolutions Futures

- **Validation Avancée des Champs** (ex : tailles minimales des titres).
- **Gestion des Migrations SQL** au lieu de l'auto-migration GORM.
- **Création d'un Package `configs`** pour centraliser la gestion de la configuration.
- **Ajout d'une Couche `services`** pour isoler la logique métier complexe.
- **Tests de Charge** pour simuler plusieurs utilisateurs.
- **Sécurisation Complète** (authentification, rôles, permissions).

---

## 📄 Licence et Avertissement

Ce projet est un projet personnel réalisé dans un but d’apprentissage.  
Il est fictif et **n'est pas destiné à un usage en production**.

---
