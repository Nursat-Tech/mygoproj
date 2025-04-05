---
tags:
  - golang
  - rest-api
  - rest
---

# Футбол API

Бұл жоба **REST API** ұсынатын футболшылардың деректерін басқару жүйесін құрайды. Осы API арқылы футболшылар туралы мәліметтерді құру, оқу, жаңарту және жоюға болады. Жоба **Gin** және **GORM** фреймворктарын қолданады, PostgreSQL дерекқорымен жұмыс істейді.

## Негізгі компоненттер

1. **Gin** — HTTP серверлері мен API жасау үшін жеңіл және жоғары өнімділікті Go фреймворкы.
2. **GORM** — PostgreSQL және басқа да дерекқорлармен жұмыс істеуге мүмкіндік беретін ORM (Object-Relational Mapping) кітапханасы.

## Орнату және іске қосу

### 1. Зависимостарды орнату

Жұмысты бастамас бұрын, барлық тәуелділіктерді орнату керек. Бұл үшін келесі команданы орындаңыз:

```sh
go mod tidy
2. Docker арқылы іске қосу
Проектті іске қосу үшін Docker және Docker Compose қолдануды ұсынамыз.

Қажетті контейнерлерді тартып, жобаны құрыңыз:

sh
Копировать
Редактировать
docker-compose up --build
Сервер іске қосылған соң, ол http://localhost:8080 мекен-жайында қолжетімді болады.

3. Дерекқорды баптау
PostgreSQL дерекқоры автоматты түрде контейнер арқылы құрылып, дерекқорды іске қосады.

Gin көмегімен сервер мысалы
Футболшылармен жұмыс істейтін негізгі серверді жасаймыз:

go
Копировать
Редактировать
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourproject/internal/handler"
	"github.com/yourusername/yourproject/internal/repository"
	"github.com/yourusername/yourproject/internal/service"
	"github.com/yourusername/yourproject/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "postgres://postgres:51177477@localhost:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Дерекқорға қосылу қатесі:", err)
	}

	err = db.AutoMigrate(&models.Player{})
	if err != nil {
		log.Fatal("Миграция қатесі:", err)
	}

	playerRepo := repository.NewPlayerRepository(db)
	playerService := service.NewPlayerService(playerRepo)
	playerHandler := handler.NewPlayerHandler(playerService)

	r := gin.Default()
	r.GET("/players", playerHandler.GetPlayers)
	r.GET("/players/:id", playerHandler.GetPlayer)
	r.POST("/players", playerHandler.CreatePlayer)
	r.PUT("/players/:id", playerHandler.UpdatePlayer)
	r.DELETE("/players/:id", playerHandler.DeletePlayer)

	r.Run(":8080") // Серверді 8080 портында іске қосу
}
API маршруты
API келесі операцияларды қолдайды:

GET /players — барлық футболшылардың тізімін алу.

GET /players/:id — белгілі бір футболшы туралы мәліметтерді алу.

POST /players — жаңа футболшы қосу.

PUT /players/:id — футболшы туралы мәліметтерді жаңарту.

DELETE /players/:id — футболшыны жою.

Мысалы, жаңа футболшы қосу сұранысы:

bash
Копировать
Редактировать
POST /players
{
  "name": "Lionel Messi",
  "position": "Forward",
  "team": "Paris Saint-Germain",
  "nationality": "Argentina"
}
GORM көмегімен жұмыс
1. Дерекқорға қосылу
GORM арқылы PostgreSQL дерекқорына қосылу мысалы:

go
Копировать
Редактировать
package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "postgres://postgres:51177477@localhost:5432/postgres?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Дерекқорға қосылу қатесі:", err)
	}

	log.Println("Қосылу сәтті аяқталды!")
}
2. Футболшы моделін анықтау
GORM моделін анықтау үшін структураны қолданамыз:

go
Копировать
Редактировать
type Player struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string
	Position  string
	Team      string
	Nationality string
}
3. Миграциялар және кестелер жасау
Кестені жасау үшін AutoMigrate әдісін қолданамыз:

go
Копировать
Редактировать
db.AutoMigrate(&Player{})
4. CRUD операциялары
Футболшыны құру:

go
Копировать
Редактировать
player := Player{Name: "Cristiano Ronaldo", Position: "Forward", Team: "Al Nassr", Nationality: "Portugal"}
db.Create(&player)
Деректерді оқу:

go
Копировать
Редактировать
var player Player
db.First(&player, 1) // ID=1 футболшысын табу
Деректерді жаңарту:

go
Копировать
Редактировать
db.Model(&player).Update("Team", "Manchester United")
Деректерді жою:

go
Копировать
Редактировать
db.Delete(&player)
5. Қатынас (1-мен 1, 1-мен көп, көп-мен көп)
Футболшылар мен олардың матчтары арасында "бірнеше матчтар бір футболшыға" деген қатынасты көрсету:

go
Копировать
Редактировать
type Match struct {
	ID        uint
	Opponent  string
	PlayerID  uint
}

type Player struct {
	ID     uint
	Name   string
	Matches []Match `gorm:"foreignKey:PlayerID"`
}
Қорытынды
Бұл жоба REST API арқылы футболшылар туралы деректерді басқаруға арналған. Gin және GORM фреймворктарын қолдана отырып, сіз футболшыларды басқаруға арналған тиімді және жеңіл API жасай аласыз. Жобаға қосымша функционалдар қосу арқылы оны кеңейтуге болады, мысалы, аутентификация, авторизация, қателерді өңдеу және басқа да мүмкіндіктер.

Копировать
Редактировать

Бұл README сіздің жобаңыздың құрылымын, API мүмкіндіктерін және оны іске қосу бойынша нұсқауларды қамтиды