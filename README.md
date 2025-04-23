# 🕹️ Go Idle Empire

Простая кликер-экономическая игра, написанная на Go, вдохновлённая Royal Cash.  
Полноценный DevOps CI/CD pipeline с автосборкой, деплоем и Docker-контейнерами под капотом.  
Теперь с HTTPS-доступом, Nginx-прокси и автоматическим SSL от Let's Encrypt 🔐

---

## 🚀 Демонстрация (Production)

🔗 https://asyagitov-idle.hostkey.in/status  
🔘 Поддерживается HTTPS через Let's Encrypt

---

## 🧩 Стек технологий

- 📦 Язык: **Go**
- 🐳 Контейнеризация: **Docker**
- 🌐 Reverse Proxy: **Nginx**
- 🔐 SSL: **Certbot + Let's Encrypt**
- 🤖 CI/CD: **GitHub Actions**
- 🗂️ Registry: **DockerHub** (`tercel340/go-idle-empire`)
- ☁️ Сервер: `205.172.58.56` (`asyagitov-idle.hostkey.in`)

---

## 📦 API

| Endpoint         | Method | Описание                 |
|------------------|--------|--------------------------|
| `/status`        | GET    | Получить кол-во очков    |
| `/click`         | POST   | Увеличить счётчик кликов |

---

## 🔧 Локальный запуск

```bash
git clone https://github.com/Airat1998/go-idle-empire.git
cd go-idle-empire
docker-compose up --build
