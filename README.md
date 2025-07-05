# react-go-oracle-app
# 🔗 URL Shortener App

A full-stack URL Shortener application that allows users to input long URLs and generate short, shareable links. Built using **ReactJS** (frontend), **Go Fiber** (backend), and **Oracle Database** (storage).

---

## 🖥️ Tech Stack

### Frontend (React)
- **ReactJS** – UI framework
- **Formik** – Form handling and validation
- **React Router DOM** – Page navigation
- **Axios** – API calls to backend

### Backend (Go)
- **Go Fiber** – Lightweight web framework for routing and middleware
- **go-ora** – Oracle database driver for Go
- **Oracle DB** – Persistent storage for original URLs, shortcodes, and shortened URLs

---

## 📐 Architecture Overview

1. **User inputs** a long URL in the frontend form.
2. **React frontend** sends the URL to the Go backend using Axios.
3. The **Go backend** generates a unique shortcode, saves the mapping to Oracle DB, and returns the shortened URL.
4. The **short URL** can be accessed to trigger a **redirect** to the original URL using the stored mapping.

---

## 🚀 Features

- Convert long URLs into short, shareable links
- Unique shortcode generation
- Redirect to original URL when short link is visited
- Oracle DB-backed persistence
- Clean, responsive UI

---

## 📁 Project Structure
- url-shortener/
    - client/ # React frontend
        - public/
          - index.html
        - src/
          - forminterface
            - styles
              - beforestyles.css
              - afterstyles.css
          - PageIndex.jsx
          - AfterShortened.jsx
          - BeforeShortened.jsx
        - index.js
      - package-lock.json
      - package.json
  - server
    - database
      - db.go
    - handlers
      - handler.go
    - model
      - url_model.go
    - go.mod
    - go.sum
    - main.go
    - .gitignore
    - README.md
