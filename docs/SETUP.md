# AstroCMS Setup Guide

## 1. Project Structure
- Created directories: `frontend/`, `backend/`, `cms/`, `builder/`, `deploy/`, `docs/`
- Each directory contains a README describing its purpose and tech stack.

## 2. Frontend Setup
- Bootstrapped with Vite + React + TypeScript.
- Installed and configured TailwindCSS v3.x, PostCSS, and Autoprefixer.
- Added Zustand (state), Headless UI (components), and Framer Motion (animation).
- Tailwind config scans `src/` and `index.html`.
- Main CSS imports Tailwind directives.

## 3. Backend Setup
- Installed Go via Homebrew.
- Initialized Go module in `backend/`.
- Created a basic HTTP server (`main.go`).
- Installed PostgreSQL via Homebrew and started the service.
- Created a development database: `astrocms_dev`.
- Added dependencies: `pq` (PostgreSQL driver), `golang-jwt/jwt` (JWT), `bcrypt` (password hashing).
- Created a `users` table for authentication.
- Implemented `/register` (user registration with hashed passwords) and `/login` (JWT authentication) endpoints.

## 4. Directus CMS Setup
- Cleaned the `cms/` directory for a fresh install.
- Created a PostgreSQL database: `directus`.
- Initialized a new Directus project in `cms/` using PostgreSQL (no SSL).
- Created the first admin user for Directus.

## 5. Next Steps

### Immediate
- [ ] Start Directus (`npx directus start`) and complete initial admin setup.
- [ ] Document Directus admin URL and credentials (in a secure place).

### Integration
- [ ] Integrate Directus with the Go backend:
    - Set up API calls from Go to Directus (REST or GraphQL)
    - Add endpoints in Go for content management via Directus
    - Implement authentication/authorization flow between Go and Directus
- [ ] Document integration patterns and best practices.

### Astro Builder
- [ ] Initialize Astro project in `builder/`.
- [ ] Set up integration with Directus for dynamic content.
- [ ] Scaffold a basic site template and test content fetching.

### General
- [ ] Add environment variable management for all services.
- [ ] Write developer onboarding documentation.
- [ ] Set up CI/CD for all components.

---

> For detailed commands and troubleshooting, see the main README and this file. Update this document as you progress. 