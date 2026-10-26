# go-reload
# GO-RELOADED

## 1. Présentation du projet

Go-Reloaded est un programme développé en langage Go.

Son objectif est de lire un fichier texte, d'appliquer différentes transformations et d'enregistrer le résultat dans un nouveau fichier.

Le programme permet notamment de modifier les majuscules et les minuscules, de gérer certaines apostrophes, de corriger la ponctuation et d'effectuer des conversions de caractères en binaire et en hexadécimal.

Ce projet permet de mettre en pratique plusieurs notions importantes du langage Go.

Parmi ces notions, on retrouve les fonctions, les boucles, les conditions, les slices, la manipulation des chaînes de caractères et la lecture et l'écriture de fichiers.

Le programme est organisé en plusieurs fonctions afin de séparer les différentes étapes du traitement.

Chaque fonction possède un rôle précis et est appelée lorsque le programme en a besoin.

---

## 2. Objectifs du projet

Les principaux objectifs de ce projet sont :

- Apprendre à manipuler des chaînes de caractères.
- Utiliser des boucles `for` pour parcourir du texte.
- Utiliser des conditions `if` et des instructions `switch`.
- Comprendre le fonctionnement des slices.
- Ajouter et supprimer des éléments avec `append()`.
- Créer et appeler plusieurs fonctions.
- Lire un fichier avec `os.ReadFile()`.
- Écrire dans un fichier avec `os.WriteFile()`.
- Gérer les erreurs.
- Comprendre comment plusieurs fonctions peuvent travailler ensemble.

L'objectif est également de rendre le code plus facile à comprendre en séparant les différentes fonctionnalités.

---

## 3. Organisation du programme

Le programme utilise un fichier principal appelé `main.go`.

Ce fichier contient la fonction `main()`, qui représente le point de départ du programme.

Il contient également plusieurs fonctions permettant de lire les fichiers et de gérer les commandes.

Les autres fonctions peuvent être placées dans différents fichiers `.go` du même dossier.

Ces fonctions sont utilisées pour appliquer les transformations demandées.

L'ensemble des fichiers Go appartient au package `main`.

### Principales fonctions

| Fonction | Rôle |
|---|---|
| `main()` | Organise les étapes du programme |
| `lireEntree()` | Recherche et lit le fichier d'entrée |
| `appliquerCommandes()` | Applique les commandes de transformation |
| `lireCommande()` | Analyse une commande |
| `estCommande()` | Reconnaît les commandes |
| `separerPonctuationCommandes()` | Sépare la ponctuation des commandes |
| `up()` | Transforme un mot en majuscules |
| `low()` | Transforme un mot en minuscules |
| `cap()` | Met la première lettre en majuscule |
| `apostrophe()` | Gère les apostrophes isolées |
| `versBin()` | Convertit les octets en binaire |
| `versHex()` | Convertit les octets en hexadécimal |
| `estPonctuation()` | Reconnaît les signes de ponctuation |
| `corrigerPonctuation()` | Supprime les espaces avant la ponctuation |

---

## 4. Fonctionnement général

Le programme suit plusieurs étapes.

### Étape 1 : lecture du texte

Le programme commence par rechercher le fichier `input.txt`.

Il utilise la fonction `lireEntree()` pour récupérer son contenu.

Si le fichier est introuvable, un message d'erreur est affiché.

### Étape 2 : découpage du texte

Le contenu du fichier est transformé en chaîne de caractères.

La fonction `strings.Fields()` découpe ensuite le texte en plusieurs mots.

Ces mots sont stockés dans une slice de chaînes de caractères.

Cela permet de parcourir et de modifier les éléments individuellement.

### Étape 3 : gestion des apostrophes

La fonction `apostrophe()` est appelée.

Elle permet de rapprocher les apostrophes isolées des mots qu'elles encadrent.

### Étape 4 : application des commandes

La fonction `appliquerCommandes()` recherche les commandes présentes dans le texte.

Elle applique les modifications aux mots concernés.

Les commandes utilisées sont ensuite supprimées de la slice.

### Étape 5 : reconstruction du texte

La fonction `strings.Join()` rassemble les mots.

Un espace est ajouté entre les différents éléments.

### Étape 6 : correction de la ponctuation

La fonction `corrigerPonctuation()` supprime les espaces situés avant certains signes de ponctuation.

### Étape 7 : création du fichier de sortie

Le résultat est affiché dans le terminal avec `fmt.Println()`.

Il est ensuite enregistré dans `output.txt` avec `os.WriteFile()`.

---

## 5. Explication de la fonction main()

La fonction `main()` est exécutée automatiquement au lancement du programme.

Elle permet d'appeler les différentes fonctions dans le bon ordre.

La première instruction importante est :

`texte, cheminEntree, err := lireEntree()`

Cette instruction récupère trois valeurs.

`texte` contient les données du fichier.

`cheminEntree` contient le chemin du fichier utilisé.

`err` permet de vérifier si une erreur s'est produite.

Si une erreur est détectée, le programme affiche un message et s'arrête.

Ensuite, le programme utilise :

`mots := strings.Fields(string(texte))`

Cette instruction transforme le texte en une slice de mots.

La fonction `apostrophe()` est ensuite appelée pour traiter les apostrophes.

Puis `appliquerCommandes()` effectue les transformations demandées.

Après les modifications, `strings.Join(mots, " ")` rassemble les mots.

La ponctuation est ensuite corrigée.

Enfin, `os.WriteFile()` enregistre le résultat dans le fichier de sortie.

---

## 6. Fonction lireEntree()

La fonction `lireEntree()` permet de récupérer le contenu du fichier d'entrée.

Elle retourne trois valeurs :

- Un tableau d'octets `[]byte`.
- Le chemin du fichier utilisé.
- Une éventuelle erreur.

Elle recherche le fichier dans deux emplacements.

Le premier emplacement est `input.txt`.

Le second est `go-reload/input.txt`.

Pour chaque chemin, elle essaie de lire le fichier avec `os.ReadFile()`.

Si la lecture fonctionne, elle retourne immédiatement son contenu.

Si le fichier n'existe pas, elle essaie le chemin suivant.

Si une autre erreur survient, elle la retourne.

Si aucun fichier n'est trouvé, elle retourne un message indiquant que `input.txt` est introuvable.

Cette fonction permet donc de gérer plusieurs emplacements possibles pour le fichier d'entrée.

---

## 7. Fonction appliquerCommandes()

La fonction `appliquerCommandes()` est l'une des fonctions principales du projet.

Elle reçoit une slice de mots.

Elle retourne la slice modifiée ainsi qu'une éventuelle erreur.

Elle commence par appeler `separerPonctuationCommandes()`.

Ensuite, elle parcourt les éléments de la slice avec une boucle `for`.

Pour chaque élément, elle appelle `lireCommande()`.

Cette fonction permet de savoir si l'élément correspond à une commande.

Si aucune commande n'est reconnue, la boucle continue.

Si une commande est reconnue, le programme vérifie que le nombre de mots demandé est disponible.

Il applique ensuite la transformation aux mots concernés.

Un `switch` permet de choisir entre `up()`, `low()` et `cap()`.

Après la transformation, la commande est supprimée avec `append()`.

L'indice de la boucle est décrémenté avec `indexCommande--`.

Cela permet de ne pas sauter un élément après la suppression.

Une deuxième boucle s'occupe des commandes `(hex)` et `(bin)`.

Ces commandes sont appliquées au mot précédent.

Elles sont également supprimées après leur exécution.

---

## 8. Fonction lireCommande()

La fonction `lireCommande()` permet d'analyser les commandes de modification de casse.

Elle reçoit une slice de mots et un indice.

Elle retourne plusieurs informations :

- Le nom de la commande.
- Le nombre de mots concernés.
- Le nombre d'éléments occupés par la commande.
- Un booléen indiquant si la commande est reconnue.
- Une éventuelle erreur.

Elle reconnaît directement les commandes `(up)`, `(low)` et `(cap)`.

Elle peut également analyser les commandes contenant un nombre.

Par exemple : `(up, 3)`.

La fonction utilise `strings.HasPrefix()` pour reconnaître le début de certaines commandes.

Elle utilise `strings.HasSuffix()` pour vérifier leur terminaison.

La fonction `strconv.Atoi()` permet de convertir le nombre écrit en texte vers un entier.

Si le nombre est invalide ou si la commande est mal formée, une erreur est retournée.

Cette fonction permet de centraliser l'analyse des commandes.

---

## 9. Fonction estCommande()

La fonction `estCommande()` vérifie si un mot correspond à une commande reconnue.

Elle retourne un booléen.

Elle utilise un `switch` pour reconnaître les commandes simples.

Les commandes reconnues sont `(up)`, `(low)`, `(cap)`, `(hex)` et `(bin)`.

Elle vérifie également certains formats commençant par `(up,`, `(low,` ou `(cap,`.

Pour cela, elle utilise `strings.HasPrefix()` et `strings.HasSuffix()`.

Cette fonction sert notamment à la séparation de la ponctuation et des commandes.

---

## 10. Fonction separerPonctuationCommandes()

Cette fonction permet de séparer certains signes de ponctuation lorsqu'ils sont collés à une commande.

Elle parcourt chaque mot de la slice.

La fonction `strings.TrimRight()` retire temporairement les signes de ponctuation situés à la fin du mot.

La fonction `strings.TrimPrefix()` permet ensuite de récupérer cette ponctuation.

Si le mot restant correspond à une commande reconnue, les deux éléments sont ajoutés séparément dans une nouvelle slice.

Dans le cas contraire, le mot est conservé sans modification.

Cette étape facilite la reconnaissance de certaines commandes suivies de ponctuation.

---

## 11. Fonction up()

La fonction `up()` transforme toutes les lettres d'un mot en majuscules.

Elle reçoit une chaîne de caractères.

Elle retourne une nouvelle chaîne modifiée.

Elle utilise la fonction `strings.ToUpper()`.

**Exemple :**

Entrée : `bonjour (up)`

Sortie : `BONJOUR`

La fonction peut également être utilisée avec une commande contenant un nombre.

Entrée : `bonjour tout le monde (up, 3)`

Sortie : `BONJOUR TOUT LE MONDE`

---

## 12. Fonction low()

La fonction `low()` transforme toutes les lettres d'un mot en minuscules.

Elle utilise `strings.ToLower()`.

Cette fonction reçoit une chaîne de caractères et retourne sa version en minuscules.

**Exemple :**

Entrée : `BONJOUR (low)`

Sortie : `bonjour`

Elle peut également modifier plusieurs mots grâce à la commande `(low, n)`.

---

## 13. Fonction cap()

La fonction `cap()` permet de mettre la première lettre d'un mot en majuscule.

Les autres lettres sont transformées en minuscules.

Elle commence par utiliser `strings.ToLower()`.

Elle transforme ensuite le mot en `[]rune`.

Cette conversion permet de manipuler les caractères Unicode, notamment les lettres accentuées.

La fonction vérifie si le mot est vide.

Si le mot est vide, elle retourne une chaîne vide.

Sinon, elle utilise `strings.ToUpper()` pour modifier le premier caractère.

Elle rassemble ensuite la première lettre et le reste du mot.

**Exemple :**

Entrée : `bONJOUR (cap)`

Sortie : `Bonjour`

---

## 14. Fonction apostrophe()

La fonction `apostrophe()` permet de gérer les apostrophes isolées utilisées pour encadrer du texte.

Si tu arrive a me sortir le numero de cette ligne je te paie la biere prochaine fois qu'on sort.

Elle reçoit une slice de mots et retourne une slice modifiée.

Elle utilise une variable booléenne appelée `dansApostrophe`.

Cette variable est initialisée à `false`.

Lorsque le programme rencontre une apostrophe isolée, il vérifie la valeur du booléen.

Si `dansApostrophe` est à `false`, l'apostrophe est considérée comme ouvrante.

Elle est alors collée au mot suivant.

Si `dansApostrophe` est à `true`, l'apostrophe est considérée comme fermante.

Elle est alors collée au mot précédent.

La fonction utilise `append()` pour supprimer les éléments devenus inutiles.

Elle utilise également `i--` après certaines suppressions pour adapter l'indice de la boucle.

**Exemple :**

Avant : `Il a dit : ' bonjour tout le monde '`

Après : `Il a dit : 'bonjour tout le monde'`

---

## 15. Fonction versBin()

La fonction `versBin()` reçoit une chaîne de caractères.

Elle initialise une variable `resultat` contenant une chaîne vide.

Elle parcourt ensuite les octets du mot avec une boucle `for`.

Pour chaque octet, elle utilise `fmt.Sprintf("%08b", mot[i])`.

Le format `%08b` permet d'obtenir une représentation binaire sur au moins huit chiffres.

Chaque résultat est ajouté à la chaîne finale.

**Exemple :**

Entrée : `A (bin)`

Sortie : `01000001`

Dans cette version, la fonction transforme les octets d'un mot en binaire.

Elle ne convertit pas encore un nombre binaire vers sa valeur décimale.

---

## 16. Fonction versHex()

La fonction `versHex()` fonctionne de manière similaire à `versBin()`.

Elle reçoit une chaîne de caractères.

Elle initialise une variable `resultat` vide.

Une boucle `for` parcourt les octets du mot.

Pour chaque octet, elle utilise `fmt.Sprintf("%02x", mot[i])`.

Le format `%02x` permet d'obtenir une représentation hexadécimale sur au moins deux chiffres.

Les résultats sont ajoutés les uns à la suite des autres.

**Exemple :**

Entrée : `A (hex)`

Sortie : `41`

Cette fonction réalise actuellement un encodage des octets du texte en hexadécimal.

Elle ne convertit pas encore un nombre hexadécimal en décimal.

---

## 17. Fonction estPonctuation()

La fonction `estPonctuation()` permet de vérifier si un caractère correspond à un signe de ponctuation.

Elle reçoit un caractère de type `rune`.

Elle utilise une instruction `switch`.

Les signes reconnus sont :

- Le point `.`
- La virgule `,`
- Le point d'exclamation `!`
- Le point d'interrogation `?`
- Les deux-points `:`
- Le point-virgule `;`

Si le caractère correspond à l'un de ces signes, la fonction retourne `true`.

Dans les autres cas, elle retourne `false`.

---

## 18. Fonction corrigerPonctuation()

La fonction `corrigerPonctuation()` permet de supprimer les espaces inutiles avant certains signes de ponctuation.

Elle reçoit une chaîne de caractères.

Elle initialise une variable `resultat` vide.

Elle parcourt ensuite le texte avec une boucle `for`.

Lorsqu'elle rencontre un espace, elle vérifie le caractère suivant.

Si ce caractère est un signe de ponctuation reconnu, elle utilise `continue`.

Cela permet d'ignorer l'espace et de ne pas l'ajouter au résultat.

Dans les autres cas, le caractère est ajouté à la chaîne finale.

**Exemple :**

Avant : `Bonjour , comment vas-tu ?`

Après : `Bonjour, comment vas-tu?`

La fonction ne gère pas encore toutes les règles de ponctuation et d'espacement.

---

## 19. Gestion des erreurs

Le programme utilise plusieurs vérifications pour éviter certaines erreurs.

La fonction `lireEntree()` vérifie que le fichier d'entrée existe et peut être lu.

La fonction `lireCommande()` vérifie le format des commandes contenant un nombre.

La fonction `appliquerCommandes()` vérifie que le nombre de mots demandé est disponible.

Si une commande demande plus de mots qu'il n'y en a avant elle, une erreur est retournée.

Le programme vérifie également les erreurs lors de l'écriture du fichier de sortie.

Les messages d'erreur sont affichés dans le terminal avec `fmt.Println()`.

Ces vérifications permettent de mieux comprendre pourquoi certaines opérations ne peuvent pas être exécutées.

---

## 20. Utilisation du programme

Pour utiliser Go-Reloaded, il faut avoir Go installé sur son ordinateur.

Le programme peut être ouvert dans VS Code.

Il faut ensuite créer un fichier `input.txt` contenant le texte à modifier.

Depuis le terminal, il faut se placer dans le dossier contenant les fichiers Go du projet.

La commande suivante permet de lancer le programme :

`go run .`

Le programme recherche automatiquement `input.txt`.

Il applique ensuite les différentes transformations.

Le résultat est affiché dans le terminal.

Un fichier `output.txt` est créé dans le même dossier que le fichier d'entrée.

Si ce fichier existe déjà, son contenu est remplacé.

**Remarque :** dans la version actuelle, les chemins des fichiers sont définis directement dans le code. Le programme ne lit pas les noms des fichiers depuis les arguments de la ligne de commande.

---

## 21. Exemples d'utilisation

### Exemple 1 : majuscules

**Input :**

`bonjour (up) tout le monde`

**Output :**

`BONJOUR tout le monde`

### Exemple 2 : minuscules

**Input :**

`BONJOUR (low) tout le monde`

**Output :**

`bonjour tout le monde`

### Exemple 3 : première lettre en majuscule

**Input :**

`bONJOUR (cap) tout le monde`

**Output :**

`Bonjour tout le monde`

### Exemple 4 : plusieurs mots

**Input :**

`bonjour tout le monde (up, 3)`

**Output :**

`BONJOUR TOUT LE MONDE`

### Exemple 5 : ponctuation

**Input :**

`Bonjour , comment vas-tu ?`

**Output :**

`Bonjour, comment vas-tu?`

### Exemple 6 : conversion hexadécimale actuelle

**Input :**

`A (hex)`

**Output :**

`41`

### Exemple 7 : conversion binaire actuelle

**Input :**

`A (bin)`

**Output :**

`01000001`

---

## 22. Tests à effectuer

Plusieurs tests peuvent être réalisés pour vérifier le fonctionnement du programme.

Il est possible de tester les commandes simples `(up)`, `(low)` et `(cap)`.

Il faut également tester les commandes contenant un nombre.

Les tests peuvent vérifier le comportement des apostrophes isolées.

Il est aussi utile de tester la ponctuation lorsqu'elle est séparée des mots par des espaces.

Les conversions binaires et hexadécimales peuvent être vérifiées avec des caractères simples.

Il est important de tester les cas d'erreur.

Par exemple, lancer le programme sans fichier `input.txt` permet de vérifier le message d'erreur.

Une commande demandant trop de mots permet de vérifier la gestion des indices.

Les commandes mal formées permettent de tester la fonction `lireCommande()`.

Ces tests permettent de repérer les erreurs et d'améliorer progressivement le programme.

---

## 23. Limites actuelles du programme

Certaines fonctionnalités peuvent encore être améliorées.

Les fonctions `versBin()` et `versHex()` encodent actuellement les octets des chaînes de caractères.

Dans les consignes habituelles de Go-Reloaded, ces commandes doivent convertir des nombres binaires ou hexadécimaux en décimal.

La fonction `corrigerPonctuation()` supprime les espaces avant certains signes, mais ne gère pas toutes les règles de ponctuation.

La fonction `apostrophe()` traite principalement les apostrophes isolées servant à encadrer du texte.

Le découpage avec `strings.Fields()` supprime les espacements et retours à la ligne d'origine.

Le programme utilise actuellement des noms de fichiers fixes.

Enfin, certaines commandes mal formées peuvent rester dans le texte si elles ne sont pas reconnues.

Ces limites représentent des possibilités d'amélioration du projet.

---

## 24. Notions de Go utilisées

Ce projet permet de travailler plusieurs notions importantes du langage Go.

### Les variables

Les variables permettent de stocker les données utilisées pendant le traitement.

### Les fonctions

Les fonctions permettent de séparer le programme en plusieurs parties.

### Les boucles

Les boucles `for` permettent de parcourir les mots et les caractères.

### Les conditions

Les conditions `if` permettent de vérifier différentes situations.

### Les switch

Les instructions `switch` permettent de choisir une action selon une valeur.

### Les slices

Les slices permettent de stocker et de modifier une liste de mots.

### La fonction append()

La fonction `append()` permet d'ajouter des éléments et, en combinant plusieurs parties d'une slice, d'en supprimer.

### Les chaînes de caractères

Le package `strings` permet de manipuler les mots et de modifier leur casse.

### Les fichiers

Le package `os` permet de lire et d'écrire des fichiers.

### Les erreurs

Le type `error` permet de signaler les problèmes rencontrés pendant l'exécution.

---

## 25. Conclusion

Le projet Go-Reloaded m'a permis de mettre en pratique plusieurs notions apprises pendant mes premiers cours de programmation en Go.

J'ai appris à lire un fichier texte, à manipuler son contenu et à enregistrer un résultat dans un autre fichier.

J'ai également travaillé sur les boucles, les conditions, les fonctions et les slices.

Le fait de séparer le programme en plusieurs fonctions m'a aidé à mieux comprendre son organisation.

Certaines fonctionnalités peuvent encore être améliorées, notamment les conversions numériques et la gestion complète de la ponctuation.

Ce projet représente une étape importante dans mon apprentissage du langage Go et m'a permis de mieux comprendre comment construire un programme qui traite automatiquement du texte.


## Installation de l'environnement de développement

Pour utiliser et modifier le projet Go-Reloaded, il faut installer trois outils sur son ordinateur :

- **Visual Studio Code (VS Code)** : un éditeur de code qui permet d'écrire, modifier et organiser les fichiers du projet.
- **Go (Golang)** : le langage de programmation utilisé pour développer et exécuter le programme.
- **Git** : un outil de gestion de versions qui permet de suivre les modifications du code et de récupérer des projets depuis un dépôt distant.

Les étapes suivantes expliquent comment installer ces outils sur Windows.

### 1. Installation de Visual Studio Code

**Étape 1 : Télécharger VS Code**

Se rendre sur le site officiel :

https://code.visualstudio.com/

Cliquer sur **Download for Windows** pour télécharger le programme d'installation.

**Étape 2 : Installer VS Code**

- Ouvrir le fichier téléchargé.
- Accepter les conditions d'utilisation.
- Choisir le dossier d'installation ou conserver celui proposé.
- Cocher l'option permettant d'ajouter VS Code au menu contextuel de Windows, si elle est proposée.
- Cliquer sur **Installer**.
- Une fois l'installation terminée, ouvrir Visual Studio Code.

**Étape 3 : Installer l'extension Go**

Pour faciliter la programmation en Go :

1. Ouvrir VS Code.
2. Cliquer sur l'icône **Extensions** dans le menu de gauche.
3. Rechercher `Go`.
4. Installer l'extension officielle **Go**, publiée par l'équipe Go.
5. Accepter l'installation des outils supplémentaires proposés si nécessaire.

Cette extension permet notamment de bénéficier de la coloration syntaxique, de l'autocomplétion et de la détection de certaines erreurs.

### 2. Installation du langage Go

**Étape 1 : Télécharger Go**

Se rendre sur le site officiel :

https://go.dev/dl/

Télécharger la version stable de Go adaptée à Windows.

Pour un ordinateur Windows 64 bits classique, choisir l'installateur Windows AMD64 au format `.msi`.

**Étape 2 : Installer Go**

- Ouvrir le fichier `.msi` téléchargé.
- Suivre les instructions de l'assistant d'installation.
- Conserver les paramètres proposés par défaut.
- Terminer l'installation.

**Étape 3 : Vérifier l'installation**

Ouvrir un nouveau terminal dans VS Code.

Saisir la commande suivante :

`go version`

Si Go est correctement installé, le terminal affiche la version installée.

Par exemple :

`go version go1.xx.x windows/amd64`

Si la commande n'est pas reconnue, fermer puis rouvrir VS Code et vérifier que Go a bien été ajouté au `PATH` de Windows.

### 3. Installation de Git

**Étape 1 : Télécharger Git**

Se rendre sur le site officiel :

https://git-scm.com/downloads/win

Télécharger la version de Git adaptée à Windows.

**Étape 2 : Installer Git**

- Ouvrir le fichier d'installation.
- Suivre les différentes étapes.
- Conserver les options proposées par défaut si aucune configuration particulière n'est nécessaire.
- Terminer l'installation.

**Étape 3 : Vérifier l'installation**

Ouvrir un nouveau terminal dans VS Code.

Saisir :

`git --version`

Si Git est correctement installé, le terminal affiche la version de Git présente sur l'ordinateur.

### 4. Récupérer le projet avec Git

Une fois Git installé, il est possible de récupérer un projet depuis un dépôt distant.

Ouvrir un terminal dans le dossier où le projet doit être enregistré.

Utiliser la commande :

`git clone URL_DU_DEPOT`

Remplacer `URL_DU_DEPOT` par l'adresse réelle du dépôt Git.

Git télécharge alors les fichiers du projet dans un nouveau dossier.

Il suffit ensuite d'ouvrir ce dossier dans VS Code avec **Fichier > Ouvrir un dossier**.

Si le projet est déjà présent sur l'ordinateur, cette étape n'est pas nécessaire.

### 5. Lancer Go-Reloaded

Une fois VS Code, Go et Git installés :

1. Ouvrir le dossier du projet Go-Reloaded dans VS Code.
2. Vérifier que tous les fichiers `.go` nécessaires sont présents.
3. Créer un fichier `input.txt` contenant le texte à modifier.
4. Ouvrir le terminal intégré de VS Code avec **Terminal > Nouveau terminal**.
5. Se placer dans le dossier du programme.
6. Exécuter la commande :

`go run .`

Le programme lit le fichier `input.txt`, applique les transformations et enregistre le résultat dans `output.txt`.

Le texte modifié est également affiché directement dans le terminal.

**Remarque :** Git est utile pour récupérer et gérer le projet, mais il n'est pas nécessaire pour exécuter un programme Go déjà présent sur l'ordinateur.