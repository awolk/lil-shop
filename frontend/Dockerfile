FROM node:14.4.0-alpine

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./

RUN npm install

COPY . ./

CMD ["npm", "start"]