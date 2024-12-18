# Port Releaser

Un outil simple pour libérer les ports occupés.

## Installation

```bash
go install github.com/codenamev/port-releaser@latest
```

## Utilisation

```bash
port-releaser -p 3000
```

## Paramètres

- `-p` : Le numéro du port à libérer

## Fonctionnement

1. Trouve le processus utilisant le port spécifié en utilisant la commande `netstat`
2. Termine le processus en utilisant la commande `taskkill`
3. Si le processus ne peut pas être terminé normalement, utilise le paramètre `-f` pour forcer la terminaison

## Précautions

- Les privilèges administrateur sont nécessaires pour terminer certains processus
- À utiliser avec précaution, en s'assurant que le processus à terminer n'est pas un processus système critique

## Licence

Ce projet est sous licence Creative Commons Attribution-NonCommercial 4.0 International (CC BY-NC 4.0).

Vous êtes libre de :
- Partager — copier et redistribuer le matériel sous n'importe quel format
- Adapter — remixer, transformer et créer à partir du matériel

Selon les conditions suivantes :
- Attribution — Vous devez créditer l'œuvre, intégrer un lien vers la licence et indiquer si des modifications ont été effectuées
- Pas d'Utilisation Commerciale — Vous n'êtes pas autorisé à faire un usage commercial de cette œuvre sans l'autorisation explicite de l'auteur

Plus d'informations : https://creativecommons.org/licenses/by-nc/4.0/deed.fr
