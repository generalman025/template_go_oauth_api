# Start from latest image
FROM golang:latest

# Configure the repo url so we can configure our work directory
ENV REPO_URL=github.com/generalman025/template_go_oauth_api

# Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

# /app/src/github.com/generalman025/template_go_oauth_api/src

# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src

COPY src $WORKPATH

WORKDIR $WORKPATH

# RUN echo "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig

RUN go build -o oauth-api .

# Expose port 8080 to the world:
EXPOSE 8080

CMD [ "./oauth-api" ]