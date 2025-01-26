# Question Search Application

A high-performance search application built with Go, gRPC, MongoDB, and React. This application efficiently handles and searches through 100,000+ question entries, supporting multiple question types including MCQ and ANAGRAM formats.

## 🌐 Live Demo
- Frontend: [https://speakx-query-search-assignment.onrender.com/](https://speakx-query-search-assignment.onrender.com/)
- Backend API: [https://querysearch.onrender.com](https://querysearch.onrender.com)

## 🏗️ Architecture

### Backend
- Go (Golang) for high-performance server
- gRPC for efficient client-server communication
- MongoDB for scalable data storage
- RESTful API endpoints

### Frontend
- React.js with modern hooks
- Custom CSS for styling
- Pagination for handling large result sets
- Real-time search functionality

## 🗄️ Database Structure

MongoDB collection with 100,000+ entries supporting two question types:

### MCQ Format
```json
{
    "_id": ObjectId,
    "type": "MCQ",
    "title": "Question text",
    "options": [
        {
            "text": "Option text",
            "isCorrectAnswer": boolean
        }
    ]
}
```

### ANAGRAM Format
```json
{
    "_id": ObjectId,
    "type": "ANAGRAM",
    "title": "Question text",
    "blocks": [
        {
            "text": "Block text",
            "showInOption": boolean,
            "isAnswer": boolean
        }
    ],
    "solution": "Solution text"
}
```

## 🚀 Getting Started

### Prerequisites
- Go 1.16 or higher
- Node.js 14.x or higher
- MongoDB
- Git

### Backend Setup

1. Clone the repository
```bash
git clone https://github.com/your-username/question-search-app.git
cd question-search-app/backend
```

2. Install Go dependencies
```bash
go mod download
go mod tidy
```

3. Set up environment variables (create `.env` file)
```env
PORT=8080
MONGODB_URI=your_mongodb_connection_string
DB_NAME=questionsDB
```

4. Run the backend server
```bash
go run main.go
```

The server will start on:
- HTTP: `http://localhost:8080`
- gRPC: `localhost:50051`

### Frontend Setup

1. Navigate to frontend directory
```bash
cd ../frontend
```

2. Install dependencies
```bash
npm install
```

3. Set up environment variables
```env
# .env.development
REACT_APP_BACKEND_URL=http://localhost:8080

# .env.production
REACT_APP_BACKEND_URL=https://querysearch.onrender.com
```

4. Run development server
```bash
npm start
```

5. Build for production
```bash
npm run build
```

## 📝 API Endpoints

### Search Questions
```http
GET /search?query=your_search_term
```

Response:
```json
[
    {
        "ID": "string",
        "Title": "string",
        "Type": "MCQ|ANAGRAM",
        "options": [...],  // for MCQ
        "blocks": [...],   // for ANAGRAM
        "solution": "string"  // for ANAGRAM
    }
]
```

## 🔧 Technology Stack

### Backend
- Go
- gRPC
- MongoDB Go Driver
- Gorilla Handlers (CORS)

### Frontend
- React
- Custom CSS
- Fetch API
- React Hooks

## 📊 Performance

- Handles 100,000+ question entries
- Efficient search with MongoDB indexing
- Pagination with 25 items per page
- Response time < 500ms for search queries

## 🚀 Deployment

### Backend Deployment (Render)
```yaml
services:
  - type: web
    name: question-search-backend
    env: go
    buildCommand: go build -o app
    startCommand: ./app
```

### Frontend Deployment (Render)
```yaml
services:
  - type: static
    name: question-search-frontend
    buildCommand: npm install && npm run build
    publishDir: build
    envVars:
      - key: REACT_APP_BACKEND_URL
        value: https://querysearch.onrender.com
```

## 🎨 CSS Structure

The project uses a modular CSS approach with separate files for components:

```css
/* index.css */
.min-h-screen {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

/* SearchBar styles */
.search-container {
  padding: 1.5rem;
  background: white;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

/* ResultsList styles */
.results-container {
  padding: 1.5rem;
}

.question-card {
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  margin-bottom: 1rem;
  padding: 1rem;
}

/* Pagination styles */
.pagination {
  display: flex;
  justify-content: center;
  gap: 1rem;
  margin-top: 2rem;
}
```

## 🔐 Security

- CORS configuration
- Rate limiting
- Input validation
- Secure MongoDB connection

## 🛠️ Development Tools

- VS Code
- MongoDB Compass
- MongoDB Atlas
- Postman
- Git
- Render

## 📈 Future Improvements

- [ ] Add authentication
- [ ] Implement caching
- [ ] Add more question types
- [ ] Add analytics dashboard

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## 👤 Author

**Prashant Bhargava**
- GitHub: [@bhargava-prashant](https://github.com/bhargava-prashant)

## 🙏 Acknowledgments

- MongoDB Team for excellent documentation
- Go community for gRPC support
- React community for hooks implementation

## 📞 Support

For support, call +919068520396, email prashantbhargava365@gmail.com or raise an issue in the repository.

---
⚡️ Built with ❤️ by Prashant Bhargava
