FROM node:18-alpine3.18


# Set the working directory inside the container
WORKDIR /usr/src/app

COPY package.json /usr/src/app

# Copy the rest of the application source code
COPY . /usr/src/app

RUN yarn

CMD ["node", "./index.js"]