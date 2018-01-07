FROM node:9.3-alpine AS base
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
ENV NODE_ENV=production

FROM base AS builder
COPY . .
RUN yarn install --prod && \
    mkdir /tmp/app && \
    mv node_modules /tmp/app && \
    mv package.json /tmp/app && \
    mv server.js /tmp/app

FROM base AS release
COPY --from=builder /tmp/app /usr/src/app
CMD ["npm","run","start"]