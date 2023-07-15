# Note Taking Application Backend

## Overview
This is the backend server for a note-taking application. It provides a REST API to manage user accounts and notes.

## Installation

Clone the repository:

   ```bash
   git clone https://github.com/BalaGoChainDev/NoteTakingApplication.git
   ```

Navigate to the project directory:
  ```bash
   cd NoteTakingApplication
   ```

## Run the Application

To start the application
   ```bash
   make run
   ```

## API Endpoints

The following are the API endpoints provided by the Note Taking Application Backend:

### Create a new user

**Endpoint**: `POST /signup`

**Request Payload**:
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Request Payload**:
- 200 OK on success
- 400 Bad Request if the request format is invalid

### User Login

**Endpoint**: `POST /login`

**Request Payload**:
```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Request Payload**:
- 200 OK with the session ID (sid) on successful login
- 400 Bad Request if the request format is invalid
- 401 Unauthorized if the username and password don't match

### List all notes

**Endpoint**: `GET /notes`

**Request Payload**:
```json
{
    "sid": "205fd11b-b69c-4158-90e3-ca805cf04348"
}
```

**Request Payload**:
- 200 OK with the list of notes on success
- 400 Bad Request if the request format is invalid
- 401 Unauthorized if the session ID is invalid

### Create a new note

**Endpoint**: `POST /notes`

**Request Payload**:
```json
{
    "sid": "3d89babf-c8cc-4a79-9c81-b92ac301c9b3",
    "note": "SampleNote"
}
```

**Request Payload**:
- 200 OK with the ID (id) of the newly created note
- 400 Bad Request if the request format is invalid
- 401 Unauthorized if the session ID is invalid

### Delete a note

**Endpoint**: `DELETE /notes`

**Request Payload**:
```json
{
    "sid": "3d89babf-c8cc-4a79-9c81-b92ac301c9b3",
    "id":"1"
}
```

**Request Payload**:
- 200 OK on success
- 400 Bad Request if the request format is invalid
- 401 Unauthorized if the session ID is invalid
