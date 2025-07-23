#!/bin/bash

# Rama a verificar
BRANCH="master"

# Obtener el último commit local
LOCAL=$(git rev-parse @)

# Obtener el último commit remoto
REMOTE=$(git ls-remote origin -h refs/heads/$BRANCH | cut -f1)

# Comparar
if [ "$LOCAL" != "$REMOTE" ]; then
    echo "Hay cambios en el remoto. Haciendo git pull..."
    git pull
else
    echo "Todo está actualizado."
fi
