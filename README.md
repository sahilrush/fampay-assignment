# YouTube Fetcher
<img width="1470" alt="Screenshot 2025-01-20 at 10 57 51 AM" src="https://github.com/user-attachments/assets/8067c492-14df-4144-b6f9-43194ceff0fe" />


A full-stack application that fetches YouTube videos based on search queries and stores them in a PostgreSQL database. Built with the Gin framework for the backend API and Vite for the frontend.

Video Link - [https://www.floik.com/flos/trr6/ttsn/f644e590.html?show-author=true](https://www.floik.com/flos/trr6/ttsn/f644e590.html?show-author=true)

## Features
- Periodic YouTube video fetching based on configured search queries
- REST API to retrieve stored videos
- Pagination support for video listing
- Background worker for automatic video fetching
- PostgreSQL database for persistent storage
- Frontend built with Vite and pnpm

## Setup

### Clone the Repository
```bash
git clone https://github.com/sahilrush/fampay-assignment.git
cd fampay-assignment
git checkout master
```

### Backend Setup

#### Navigate to Backend Directory
```bash
cd be
```

#### Set Up Environment Variables
Refer to the `.env.example` file in the `be` directory for the required variables and create a `.env` file with appropriate values.

#### Start the Backend
```bash
docker-compose up
```
The backend application will:
- Connect to the PostgreSQL database
- Run database migrations
- Start the background video fetcher
- Start the HTTP server on port `8080`

### Frontend Setup

#### Navigate to Frontend Directory
```bash
cd fe
```

#### Install Dependencies
```bash
pnpm install
```

#### Start the Development Server
```bash
pnpm dev
```
The development server will start at [http://localhost:3000](http://localhost:3000).

#### Build for Production
```bash
pnpm build
```
The production build will be available in the `dist` directory.

## Access the API

### List Videos (with Pagination)
```bash
curl "http://localhost:8080/videos?page=1&limit=7"
```
This will hit the backend to fetch paginated results.

## API Endpoints

### GET /videos
Lists stored videos with pagination support.

#### Query Parameters:
- `page` (default: 1): Page number
- `limit` (default: 10): Number of items per page
