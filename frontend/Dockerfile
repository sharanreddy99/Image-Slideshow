FROM node:12.22.1-alpine

WORKDIR /app

COPY package* ./
RUN npm install

COPY ./start.sh ../
RUN ["chmod","+x","../start.sh"]

COPY . ./

EXPOSE 3000

CMD ["../start.sh"]

