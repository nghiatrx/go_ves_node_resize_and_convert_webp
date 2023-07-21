FROM node:18-alpine3.18


# Set the working directory inside the container
WORKDIR /usr/src/app

COPY package.json .

# Copy the rest of the application source code
COPY . .

RUN yarn

CMD ["node", "./index.js"]