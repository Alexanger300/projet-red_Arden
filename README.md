Projet_RED 

Nom du Projet:ARDEN

                            -->Il n'est pas nécessaire d'installer quoi que ce soit<--

                                            -->Pour démarrer le jeu<--


Ce projet propose un RPG textuel où le joueur crée son personnage, explore et combat des monstres.
Voici la description des fonctions principales, organisées par thématique pour faciliter la compréhension du code et du déroulement du jeu.

Démarrage du jeu

1)          start.DisplayMenu()
Cette fonction affiche le menu principal du jeu. C’est à partir de ce menu que le joueur peut lancer une nouvelle partie, charger une sauvegarde ou quitter.

2)          Introduction() 
  Cette fonction lance la cinématique d’introduction du jeu. Elle prépare le joueur au scénario et mène ensuite à la création du personnage via InitCharacter()

3)          InitCharacter() Character 
    Cette fonction s’occupe de la création du personnage. Elle utilise la structure Character, qui définit toutes les informations du joueur (nom, classe, points de vie, mana, inventaire, etc.). C’est à partir de cette fonction que toutes les données de votre héros sont initialisées et utilisées dans le jeu.

Personnage et progression

1)          GainExp(amount int)
    Cette fonction gère l’expérience gagnée par le personnage, par exemple après avoir vaincu un ennemi. L’expérience est ajoutée au total actuel, et si le seuil requis est atteint, la fonction déclenche automatiquement un passage de niveau (LevelUp). 

2)          LevelUp()
Quand elle est appelée, cette fonction fait monter le personnage d’un niveau. Ses statistiques principales (points de vie, attaque, défense, mana, etc.) augmentent selon la classe choisie au départ (Paladin, Mage, Géant, Guérisseur). De plus, les points de vie et le mana sont entièrement restaurés.

3)          UseItem(item, target)
Cette fonction permet d’utiliser un objet de l’inventaire. Par exemple, si le joueur choisit une potion de soin, elle sera consommée et appliquera son effet directement sur la cible.

4)          UseSkillOnMonster(skillName, target)
Elle sert à activer une compétence du personnage contre un monstre. Selon la compétence choisie, l’effet peut être une attaque magique, un sort de soin ou une autre capacité spéciale. Cette fonction fait le lien entre les compétences apprises par le personnage et leur utilisation concrète dans un combat.

Combat

1)          StartFight(player, enemy)
C’est la fonction principale qui lance un combat entre le joueur et un monstre. Elle affiche l’arrivée de l’ennemi, puis gère tout le déroulement du combat en alternant les tours. Le combat continue tant que les deux personnages (joueur et monstre) sont encore en vie.

2)          playerTurn(player, enemy)
Pendant le tour du joueur, cette fonction propose plusieurs choix : attaquer, utiliser une compétence, ouvrir l’inventaire ou tenter de fuir. Selon la décision du joueur, les actions sont appliquées directement sur l’ennemi ou sur le personnage (par exemple utiliser une potion).

3)          monsterTurn(player, enemy)
C’est le tour de l’ennemi. Le monstre peut effectuer une attaque basique ou une attaque spéciale, plus puissante. Les dégâts sont alors appliqués sur le joueur.

Inventaire

1)          AddItem(name, qty)
Ajoute un objet dans l’inventaire. Si la limite par type ou la limite globale du sac est dépassée, l’ajout échoue. Sinon, l’objet est ajouté et l’état de l’inventaire est enregistré en arrière-plan dans le fichier JSON.

2)          RemoveItem(name, qty)
Retire un objet de l’inventaire. Si la quantité demandée est supérieure à ce que possède le joueur, l’action échoue. Si après retrait il n’en reste plus, l’objet est complètement supprimé de la liste. Le fichier JSON est ensuite mis à jour en arrière-plan pour garder ce changement.

3)          HasItem(name, qty)
Vérifie si le joueur possède au moins une certaine quantité d’un objet. Pour cela, le jeu lit l’inventaire sauvegardé en arrière-plan dans le fichier JSON.

4)          UpgradeBag(slots)
Augmente la capacité maximale du sac. Cela permet de transporter plus d’objets. Le fichier JSON continue à enregistrer automatiquement l’état de l’inventaire élargi.

5)          CountItem(name)
Indique combien d’exemplaires précis d’un objet se trouvent dans l’inventaire. La valeur est obtenue à partir des données stockées en arrière-plan dans le fichier JSON.

6)          ShowInventory()
Affiche le contenu actuel de l’inventaire directement en jeu. Chaque objet est listé avec sa quantité et les limites associées. En arrière-plan, les données proviennent du fichier JSON qui conserve l’état de l’inventaire.

Sauvegarde

1)          SaveGame(state)
Sauvegarde le personnage, son inventaire et son équipement dans un fichier save_slot_X.json. Concrètement, toutes les données du joueur sont d’abord rassemblées dans une structure GameState, qui contient le nom, la classe, l’argent, l’inventaire et l’équipement. Ensuite, on utilise json.MarshalIndent pour transformer cette structure Go en un texte JSON lisible (facile à relire et à comprendre). Enfin, grâce à os.WriteFile, ce texte est écrit dans le fichier correspondant au slot choisi.

2)          LoadGame(slot)
Recharge une sauvegarde depuis un slot choisi (1, 2 ou 3). Le jeu utilise os.ReadFile pour ouvrir le fichier JSON, puis json.Unmarshal pour reconvertir ce texte JSON en une structure GameState utilisable par le programme. Ainsi, toutes les données du joueur (statistiques, inventaire, progression) sont restaurées exactement comme elles étaient au moment de la sauvegarde.

3)          SlotExists(slot)
Vérifie si un fichier de sauvegarde existe pour un slot donné. Cette vérification se fait avec os.Stat, qui permet de savoir si le fichier est présent sur l’ordinateur. Si le fichier n’existe pas, cela signifie que le slot est vide.

4)          DeleteSlot(slot)
Supprime un fichier de sauvegarde. Ici, os.Remove est utilisé pour effacer le fichier correspondant au slot choisi. Une fois supprimé, il n’est plus possible de charger cette sauvegarde.



Il est nécessaire d'aller dans le Terminal et ensuite avec la commande ls, il doit apparaitre le fichier main.go.

Si il y est,alors vous pouvez lancez le jeu en utilisant la commande:

--->go run .\main.go<----

Profitez bien du jeu :)