# GOAsDirect: Modern WordPress Alternative

GOAsDirect is a modern, agency-focused alternative to WordPress, designed for rapid deployment of React, Next.js, and Astro-powered websites and webapps. It leverages a headless architecture with Directus for content management and Go for backend orchestration.

## 📁 Project Structure

```
AstroCMS/
├── frontend/   # React + Vite dashboard (Tauri wrapper)
├── backend/    # Go API server
├── cms/        # Directus instance
├── builder/    # Astro site builder
├── deploy/     # Deployment orchestration
└── docs/       # Documentation
```

## 🏗️ Directory Purpose

- **frontend/**: Source code for the React dashboard, wrapped with Tauri for desktop use. Handles user interface, client management, and site building.
- **backend/**: Go-based API server for authentication, multi-tenant logic, and orchestration between services.
- **cms/**: Directus headless CMS instance for content modeling, editing, and media management.
- **builder/**: Astro-powered static site builder, integrates with Directus for dynamic content.
- **deploy/**: Scripts and configuration for deployment automation (Vercel, Netlify, etc.).
- **docs/**: All project documentation, architecture decisions, and guides.

## 🧩 High-Level Architecture

- **Frontend**: React + Vite + TailwindCSS, Tauri desktop wrapper
- **Backend**: Go API server, connects frontend, Directus, and builder
- **CMS**: Directus for content management
- **Site Builder**: Astro for static/dynamic site generation
- **Deployment**: Automated scripts for modern hosting platforms

## 🚀 Getting Started

1. Clone the repository
2. See each directory's README for setup instructions (coming soon)
3. Follow the roadmap in `docs/` for development steps

---

> This project is in early development. See `docs/` for the latest updates and architecture decisions. 
