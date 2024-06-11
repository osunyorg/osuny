# Notes sur la mise en place

```bash
go install github.com/spf13/cobra-cli@latest
```
Il semble qu'il y ait un pb de chemin.

```bash
go init github.com/osunyorg/cli
go mod tidy
go get -u github.com/spf13/cobra@latest
```

L'installation de Cobra-CLI se fait avec 
```bash
go install github.com/spf13/cobra-cli@latest
```

Il a fallu ajouter ça au fichier `.zshrc` pour que le module se charge
```
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```

Ensuite, création du fichier .cobra.yaml
```yaml
author: noesya <contact@noesya.coop>
year: 2024
license: MIT
```
Finalement je ne l'utilise pas (-> supprimé).

Initialisation de l'app
```bash
cobra-cli init --config .cobra.yaml
```

Ajout d'une commande
```bash
cobra-cli add watch --config .cobra.yaml
```

