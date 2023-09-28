# Utilisez l'image officielle de Go
FROM golang:latest

# Répertoire de travail dans le conteneur
WORKDIR /app

# Copiez tout le contenu local dans le conteneur
COPY . .

# Installez les dépendances
RUN go install
RUN go build .

ENTRYPOINT [ "/app/api" ]