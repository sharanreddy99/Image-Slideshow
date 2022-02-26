FROM node:12.22.1-alpine

WORKDIR /app

COPY package* ./
RUN npm install

COPY . ./

EXPOSE 3000

RUN ["chmod","+x","start.sh"]
CMD ["./start.sh"]