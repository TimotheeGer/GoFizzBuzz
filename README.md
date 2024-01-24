# Lancement du Serveur
Pour lancer le serveur, suivez ces étapes :

## Prérequis :
- Assurez-vous d'avoir Docker et Docker Compose installés sur votre machine.

## Cloner le Répertoire :
- Clonez le répertoire du projet sur votre machine locale.

## Lancer le Serveur :
- Ouvrez un terminal et naviguez vers le répertoire du projet.
- Exécutez la commande suivante pour démarrer le serveur : "docker-compose up"

# Utilisation de l'API

## Endpoint Fizz-Buzz
- **URL** : `/fizzbuzz`
- **Méthode** : `GET`
- **Paramètres** :
- `int1` : Premier entier pour le remplacement.
- `int2` : Deuxième entier pour le remplacement.
- `limit` : Limite de la séquence numérique.
- `str1` : Chaîne de remplacement pour `int1`.
- `str2` : Chaîne de remplacement pour `int2`.

## Endpoint de Statistiques
- **URL** : `/statistics`
- **Méthode** : `GET`
- **Description** : Retourne les paramètres de la requête la plus fréquente et le nombre de fois effectuée.
