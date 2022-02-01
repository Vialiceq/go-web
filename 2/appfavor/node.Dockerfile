FROM node:14.16.1-alpine3.13 AS JS_BUILD
COPY webapp /webapp
WORKDIR /webapp
# Expose API port to the outside
#ENV PORT 80
EXPOSE 3000
RUN npm install && npm run build