package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
	"sfdbhw/config"
	"sfdbhw/entity"
	"sfdbhw/service"
	"sfdbhw/storage"
	"sfdbhw/storage/postgres"
)

func main() {
	// создаем подключение к БД
	db := postgres.NewPostgres(config.Config.Dsn)

	storages := storage.NewStorages(db)
	m, err := migrate.New(
		"file://././storage/postgres",
		config.Config.Dsn,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		fmt.Println(err)
	}
	services := service.NewServices(storages)
	for {

		fmt.Println(`Введите номер действия:
1. Добавить задачу
2. Получить все задачи
3. Получить задачи по автору
4. Получить задачи по лейблу
5. Обновить задачу
6. Удалить задачу
7. Выход`)
		var n int
		fmt.Scanln(&n)
		switch n {
		case 1:
			raw := `
		{
  			"title": "Fix login issue",
  			"content": "Users are unable to log in.",
  			"authorID": 1,
  			"assignedID": 2,
  			"opened": 1633027200,
  			"closed": 1633113600,
  			"label": [
    			{
      				"id": 2
    			}
  			]
		}
	`
			ctx := context.Background()
			task := entity.Task{}
			err = json.Unmarshal([]byte(raw), &task)
			if err != nil {
				fmt.Println(err)
			}
			_, err = services.NewTask(ctx, task)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			var tasks []entity.Task
			ctx := context.Background()
			tasks, err = services.AllTasks(ctx)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(tasks)
		case 3:
			var tasks []entity.Task
			var id int
			fmt.Print("Введите id автора: ")
			fmt.Scanln(&id)
			ctx := context.Background()
			tasks, err = services.TaskByAuthor(ctx, id)
			fmt.Println(tasks)
		case 4:
			var tasks []entity.Task
			var id int
			fmt.Print("Введите id лейбла: ")
			fmt.Scanln(&id)
			ctx := context.Background()
			tasks, err = services.TaskByLabel(ctx, id)
			fmt.Println(tasks)
		case 5:
			raw := `
		{	
			"id": 1,
			"title": "new title",
  			"content": "new content",
  			"authorID": 2,
  			"assignedID": 3,
  			"opened": 1633027200,
  			"closed": 1633113600,
  			"label": [
    			{
      				"id": 3
    			}
  			]
		}
	`
			ctx := context.Background()
			task := entity.Task{}
			err = json.Unmarshal([]byte(raw), &task)
			if err != nil {
				fmt.Println(err)
			}
			err = services.UpdateTask(ctx, task)
			if err != nil {
				fmt.Println(err)
			}
		case 6:
			var id int
			fmt.Print("Введите id задачи: ")
			fmt.Scanln(&id)
			ctx := context.Background()
			err = services.DeleteTask(ctx, id)
			if err != nil {
				fmt.Println(err)
			}
		case 7:
			os.Exit(0)
		default:
			os.Exit(0)

		}

	}
}
