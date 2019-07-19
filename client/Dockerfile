FROM node:10-alpine as build-deps
ENV NODE_ENV production
WORKDIR /usr/src/app
COPY . .
RUN yarn install && yarn run build

FROM nginx:stable-alpine
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=build-deps /usr/src/app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
