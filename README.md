# go-reload
Explication du projet :

Le programme lit un fichier input.txt, récupère son contenu et applique différentes modifications comme la mise en majuscules ou minuscules, la gestion des apostrophes, la ponctuation et les conversions binaires et hexadécimales.

Une fois les modifications effectuées, le programme enregistre le texte corrigé dans un nouveau fichier appelé output.txt.

## Explication des fonctions

### 1. Fonction up()

La fonction `up()` permet de transformer toutes les lettres d'un mot en majuscules grâce à `strings.ToUpper()`.

**Exemple :** `bonjour (up)` → `BONJOUR`

### 2. Fonction low()

La fonction `low()` permet de transformer toutes les lettres d'un mot en minuscules grâce à `strings.ToLower()`.

**Exemple :** `BONJOUR (low)` → `bonjour`

### 3. Fonction cap()

La fonction `cap()` permet de mettre la première lettre d'un mot en majuscule et toutes les autres en minuscules.

Elle utilise `strings.ToLower()` pour mettre le mot en minuscules, puis `strings.ToUpper()` pour mettre la première lettre en majuscule.

La fonction utilise également `[]rune` pour gérer les caractères accentués.

**Exemple :** `bONJOUR (cap)` → `Bonjour`

### 4. Fonction apostrophe() -->

La fonction apostrophe() permet de supprimer les espaces inutiles autour des apostrophes utilisées pour encadrer une citation.

Elle parcourt les mots du texte avec une boucle for et utilise une variable booléenne pour savoir si l'apostrophe rencontrée est ouvrante ou fermante.

Elle utilise ensuite append() pour regrouper les éléments et supprimer les apostrophes isolées de la slice après les avoir rattachées aux mots.

Exemple :

Avant : Il a dit : ' bonjour tout le monde '

Après : Il a dit : 'bonjour tout le monde'

### 5. Fonction versBin()

La fonction versBin() permet de convertir chaque octet d'un mot en binaire.

Elle utilise une boucle for pour parcourir le mot, puis fmt.Sprintf("%08b", ...) pour convertir chaque octet en une représentation binaire sur 8 bits.

Exemple : A (bin) → 01000001

### 6. Fonction versHex()

La fonction versHex() permet de convertir chaque octet d'un mot en hexadécimal.

Elle utilise une boucle for pour parcourir le mot, puis fmt.Sprintf("%02x", ...) pour convertir chaque octet en hexadécimal.

Exemple : A (hex) → 41

### 7. Fonction estPonctuation()

La fonction estPonctuation() permet de vérifier si un caractère est un signe de ponctuation.

Elle utilise un switch pour reconnaître les caractères . , ! ? : ; et retourne true ou false.

### 8. Fonction corrigerPonctuation()

La fonction corrigerPonctuation() permet de supprimer les espaces placés avant les signes de ponctuation.

Elle parcourt le texte avec une boucle for et vérifie si un espace est suivi d'un signe de ponctuation. Si c'est le cas, elle ignore cet espace grâce à continue.

Exemple :

Avant : Bonjour , comment vas-tu ?

Après : Bonjour, comment vas-tu?

## Utilisation du programme

Pour utiliser Go-Reloaded, il faut avoir Go installé sur son ordinateur.

1. Ouvrir le projet dans VS Code.
2. Créer un fichier `input.txt` contenant le texte à modifier.
3. Ouvrir le terminal dans le dossier du projet.
4. Exécuter la commande `go run .`

Le programme lit automatiquement le fichier `input.txt`, applique les modifications et enregistre le résultat dans un fichier `output.txt`.

Le résultat est également affiché dans le terminal.