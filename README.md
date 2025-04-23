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

👨‍💻 Автор
Airat Sagitov