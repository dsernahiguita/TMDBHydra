FROM node:latest
# Add Maintainer Info
LABEL maintainer="Diana Lucia Serna <dserna@gmx.net>"

# Set the Current Working Directory inside the container
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY . /usr/src/app

RUN npm install --silent
RUN npm install react-scripts@3.4.3 -g --silent

EXPOSE 3000 3001

CMD ["npm", "start"]
