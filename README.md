<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Thanks again! Now go create something AMAZING! :D
***
***
***
*** To avoid retyping too much info. Do a search and replace for the following:
*** github_username, repo_name, twitter_handle, email, project_title, project_description
-->

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

<!-- PROJECT LOGO & TITLE -->
<p align="center">
  <h1 align="center">Guest Check</h1>
  <p align="center">
    <b>Modern microservice for guest check and time recording, built with Go, REST, gRPC, Kafka, and Kubernetes.</b><br>
    <i>Fast, scalable, and cloud-native solution for hospitality and retail environments.</i>
    <br />
    <a href="#getting-started"><strong>Get Started »</strong></a>
    ·
    <a href="https://github.com/patricksferraz/pinned-guest-check/issues">Report Bug</a>
    ·
    <a href="https://github.com/patricksferraz/pinned-guest-check/issues">Request Feature</a>
  </p>
</p>

---

<!-- BADGES -->
<p align="center">
  <img alt="Go version" src="https://img.shields.io/badge/Go-1.18+-00ADD8?logo=go">
  <img alt="Docker" src="https://img.shields.io/badge/Docker-ready-blue?logo=docker">
  <img alt="Kubernetes" src="https://img.shields.io/badge/Kubernetes-ready-326ce5?logo=kubernetes">
  <img alt="License" src="https://img.shields.io/badge/license-MIT-green">
  <img alt="PRs Welcome" src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square">
</p>

---

## 🚀 About the Project

**Guest Check** is a cloud-native microservice designed for fast, reliable, and scalable guest check and time recording operations. It provides a robust API for managing guest checks, items, and payments, supporting both REST and gRPC protocols, and integrates seamlessly with Kafka for event-driven architectures.

- **Use case:** Ideal for restaurants, hotels, and retail environments needing modern, distributed guest check management.
- **Why Guest Check?**
  - Cloud-native, scalable, and production-ready
  - Built with Go for performance and reliability
  - Easy integration via REST, gRPC, and Kafka
  - Ready for Docker and Kubernetes deployments

---

## 🧩 Features

- Full CRUD for guest checks and items
- RESTful API and gRPC support
- Event-driven with Kafka (consume and produce)
- PostgreSQL and MongoDB support via GORM
- OpenAPI/Swagger documentation
- Cloud-native: Docker & Kubernetes ready
- Secure: API key authentication
- Extensible and easy to contribute

---

## 🏗️ Architecture & Tech Stack

- **Language:** Go (Golang)
- **Frameworks:** Fiber (REST), Cobra (CLI), GORM (ORM)
- **Messaging:** Kafka (Confluent)
- **API:** REST (Fiber), gRPC
- **Database:** PostgreSQL, MongoDB
- **Containerization:** Docker, Docker Compose
- **Orchestration:** Kubernetes (k8s)
- **Docs:** Swagger/OpenAPI

---

## ⚡ Quickstart

### Prerequisites
- Go 1.18+
- Docker & Docker Compose
- Kubernetes cluster (local or cloud)
- [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
- [Helm](https://helm.sh/) (optional)

### Local Development (Docker Compose)
```sh
# Clone the repository
$ git clone https://github.com/patricksferraz/pinned-guest-check.git
$ cd pinned-guest-check

# Copy and edit environment variables
$ cp .env.example .env

# Start all services
$ make up

# Run tests
$ make gtest
```

### Kubernetes Deployment
1. Fill in `k8s/.env` using `k8s/.env.example`.
2. Create secrets:
   ```sh
   kubectl create secret generic guest-check-secret --from-env-file k8s/.env
   ```
3. (Optional) Create Docker registry secret:
   ```sh
   kubectl create secret docker-registry regsecret \
     --docker-server=$DOCKER_REGISTRY_SERVER \
     --docker-username=$DOCKER_USER \
     --docker-password=$DOCKER_PASSWORD \
     --docker-email=$DOCKER_EMAIL
   ```
4. Deploy all resources:
   ```sh
   kubectl apply -f ./k8s
   ```

---

## 📚 API Overview

### REST Endpoints (Base: `/api/v1`)
- `POST   /guest-checks` — Create a new guest check
- `GET    /guest-checks` — List/search guest checks
- `GET    /guest-checks/{guest_check_id}` — Get guest check by ID
- `POST   /guest-checks/{guest_check_id}/wait-payment` — Mark as waiting payment
- `POST   /guest-checks/{guest_check_id}/cancel` — Cancel guest check
- `POST   /guest-checks/{guest_check_id}/pay` — Pay guest check
- `POST   /guest-checks/{guest_check_id}/items` — Add item to guest check
- `GET    /guest-checks/{guest_check_id}/items/{item_id}` — Get item details
- `POST   /guest-checks/{guest_check_id}/items/{item_id}/cancel` — Cancel item
- `POST   /guest-checks/{guest_check_id}/items/{item_id}/prepare` — Mark item as preparing
- `POST   /guest-checks/{guest_check_id}/items/{item_id}/ready` — Mark item as ready
- `POST   /guest-checks/{guest_check_id}/items/{item_id}/forward` — Forward item
- `POST   /guest-checks/{guest_check_id}/items/{item_id}/deliver` — Deliver item

> Full Swagger docs available at `/api/v1/swagger/index.html`

### Kafka Topics Consumed
- `OPEN_GUEST_CHECK` — Open a new guest check
- `NEW_GUEST` — Register a new guest
- `NEW_PLACE` — Register a new place
- `NEW_EMPLOYEE` — Register a new employee
- `NEW_MENU_ITEM` — Register a new menu item
- `UPDATE_MENU_ITEM` — Update menu item

---

## 🛠️ Built With
- [Go](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [Cobra](https://github.com/spf13/cobra)
- [GORM](https://gorm.io/)
- [Kafka](https://kafka.apache.org/)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/)
- [Swagger](https://swagger.io/)

---

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**!

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📬 Contact & Community

- Project Link: [https://github.com/patricksferraz/pinned-guest-check](https://github.com/patricksferraz/pinned-guest-check)

---

## 🏷️ SEO & Keywords

<sub>
Guest check, time recording, microservice, Go, Golang, REST API, gRPC, Kafka, Kubernetes, Docker, hospitality, retail, open source, cloud-native, scalable, event-driven, POS, restaurant, hotel, guest management, order management, payment, API, backend, distributed systems, openapi, swagger, fiber, gorm, cobra, postgres, mongodb, devops, CI/CD, modern architecture, best practices.
</sub>

---

<details>
<summary>Meta Description (for search engines)</summary>
Guest Check is a modern, cloud-native microservice for guest check and time recording, built with Go, REST, gRPC, Kafka, and Kubernetes. Fast, scalable, and ready for hospitality and retail environments. Open source and easy to integrate.
</details>
