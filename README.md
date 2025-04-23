# 🕹️ Go Idle Empire

Простая кликер-экономическая игра, написанная на Go, вдохновлённая Royal Cash.  
Полноценный DevOps CI/CD pipeline с автосборкой, деплоем и Docker.

---

## 🚀 Демонстрация

Staging-сервер: [http://205.172.58.56:8080/status](http://205.172.58.56:8080/status)

---

## 🧩 Стек

- 📦 Язык: **Go**
- 🐳 Контейнеризация: **Docker**
- 🤖 CI/CD: **GitHub Actions**
- 🚀 Деплой: SSH на staging
- 🗂️ Registry: **DockerHub** (`tercel340/go-idle-empire`)
- ☁️ Сервер: `205.172.58.56`

---

## 📦 API

| Endpoint         | Method | Описание                 |
|------------------|--------|--------------------------|
| `/status`        | GET    | Получить кол-во очков    |
| `/click`         | POST   | Увеличить счётчик кликов |

---

## 🛠 Локальный запуск

```bash
git clone https://github.com/Airat1998/go-idle-empire.git
cd go-idle-empire
docker-compose up --build


⚙️ CI/CD Pipeline
При пуше в main:

🔍 Проверка go vet, go test

🐳 Сборка Docker-образа

📤 Push на DockerHub

🚀 SSH-деплой на сервер (205.172.58.56)

🔐 Secrets (в GitHub)

Secret name	Назначение
DOCKER_USERNAME	логин DockerHub
DOCKER_PASSWORD	токен DockerHub
STAGING_HOST	IP сервера
STAGING_USER	root/другой пользователь
STAGING_KEY	приватный SSH ключ
📌 TODO
 Добавить версионирование Docker-образов

 Healthcheck и алерты

 Подключение через HTTPS + Nginx

 Kubernetes-деплой (в будущем)

👨‍💻 Автор
Airat Sagitov