FROM node:12.13.0-alpine

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 50052
CMD [ "node", "server.js" ]