# Idea

FileRain is a self-hosted cloud alternative. It's meant for file backup and quick retrieval via a web interface. The
goal is to keep good performance and small resource footprint.

# Tech stack

- Golang + Templ
- Postgres
- HTMX
- Developer sweat and tears
- Flutter (for mobile/desktop app)

This app will be server-side rendered. There will be a single source of truth, which is the back-end. Partial reloading
will be done using HTMX only if necessary. It would be nice to leverage all the things browsers give us out of the box.

# Roadmap

### MVP:

- Setup infrastructure
    - PostgreSQL
    - Templ support
    - Static files
    - Migrations
    - 400/500 error pages
- Authentication essentials:
    - Sign up
    - Sign in
- File uploads/downloads via a Web Browser
- File preview *(images/videos, using HTML img/video tags)*

### V1

- Thumbnail generation for:
    - Images
    - Videos
- Authentication nice to have:
    - Forgotten password
    - Session management
    - Manage other users as administrator

### V2

- [WebDav](https://datatracker.ietf.org/doc/html/rfc4918) support
- App for iOS/Android to automatically upload media from Gallery

# Prerequisites

- Golang 1.23.4
- Docker 27.5.1 + Docker Compose

# Running locally

```shell
docker compose up -d
templ generate --watch --proxy="http://localhost:80" --cmd="go run ."
```