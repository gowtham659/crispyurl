FROM node:alpine
LABEL author="gowthamthota"

WORKDIR /app

COPY package*.json .

RUN npm install

COPY . . 

RUN npm install -g serve

RUN npm run build

EXPOSE 3000
CMD [ "serve","-s","build" ]


