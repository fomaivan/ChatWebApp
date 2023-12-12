FROM node:14-alpine as builder

WORKDIR /app

COPY package.json ./
COPY yarn.lock ./


# RUN yarn install --frozen-lockfile
RUN yarn install --ignore-engines
COPY . .
CMD ["yarn", "start"]