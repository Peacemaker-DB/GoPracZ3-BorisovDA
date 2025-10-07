package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pz3-http/internal/api"
	"pz3-http/internal/storage"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	store := storage.NewMemoryStore()
	h := api.NewHandlers(store)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		api.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Коллекция
	mux.HandleFunc("GET /tasks", h.ListTasks)
	mux.HandleFunc("POST /tasks", h.CreateTask)
	// Элемент
	mux.HandleFunc("GET /tasks/{id}", h.GetTask)
	mux.HandleFunc("PATCH /tasks/{id}", h.PatchTask)
	mux.HandleFunc("DELETE /tasks/{id}", h.DELETETask)

	// Подключаем логирование
	handler := api.CORS(api.Logging(mux))

	addr := ":" + port
	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	// Канал для ошибок сервера
	serverErr := make(chan error, 1)

	go func() {
		log.Println("listening on", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
		close(serverErr)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		log.Printf("received signal: %s, Плавное отключение сервера", sig)
	case err := <-serverErr:
		log.Printf("server error: %v, Отключение сервера", err)
	}

	// Создаем контекст с таймаутом для Завершения
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Выключаем сервер
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Плавное отключение сервера ошибка: %v", err)
	}

	log.Println("Сервер плавно отключен")
}
