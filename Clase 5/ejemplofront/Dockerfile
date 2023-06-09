FROM node:18

WORKDIR /frontend

COPY package.json ./

# COPY package-lock.json ./
COPY yarn.lock ./

# RUN npm install
RUN yarn

ENV BACK_HOST localhost

COPY . .

EXPOSE 5173

CMD ["yarn", "dev"]