# Use an official Node.js image as the base
FROM node:18

# Set the working directory inside the container
WORKDIR /app

# Copy the package.json and install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy the rest of the frontend files
COPY . ./

# Expose the frontend development server port
EXPOSE 3000

# Run the frontend development server
CMD ["npm", "start"]
