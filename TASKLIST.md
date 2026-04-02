## Цель

backend + frontend

Общее состояние, без сессий, без хранения данных.

### TASKS

[ ] Сделать API
[ ] Нарисовать ебаный фронт. Основа - спиздить с etappi, дополнительно - расширенная статистика по 1 человеку, типа урона по структурам. Возможно показ логов по таймлайну.
[ ] Маст хев надо указывать скрытые баффы - бонус новичка, гад гифт, бонус смертника.

## API

TODO для запроса Filter уже надо знать что запрашивать (Имена игроков, тип урона, т.д.)

```go
type Service interface {
    Init(fs.FS) error

    GetGames() ([]Game, error)
    Filter(gameID int, filters map[string]Filter) (map[string]any, error)
}

type Timeline struct {
    From, To time.Time
}

type Filter interface {
    Apply(line LogLine) error
    GetResult() (any, error)
}


type FilterRequest struct {
    GameID int
    Filters map[string]*FilterStruct
}

type Quer

```