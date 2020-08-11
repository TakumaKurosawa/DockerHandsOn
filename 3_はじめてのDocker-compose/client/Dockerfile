FROM node:10.15.1-alpine as builder
WORKDIR /app

COPY ./package.json /app

RUN yarn

COPY . /app
RUN yarn build

FROM node:10.15.1-alpine
WORKDIR /app
COPY --from=builder /app /app
CMD ["yarn", "start"]