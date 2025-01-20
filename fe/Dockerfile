# Use the official Node.js image as a base image
FROM node:18-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the package.json and package-lock.json files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Expose the port that Vite runs on
EXPOSE 5173

# Command to start Vite in development mode
CMD ["npm", "run", "dev", "--", "--host"]
