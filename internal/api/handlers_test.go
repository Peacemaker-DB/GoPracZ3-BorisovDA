package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pz3-http/internal/storage"
	"testing"
)

func TestHandlers_CreateTask(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := NewHandlers(store)

	tests := []struct {
		name           string
		payload        interface{}
		expectedStatus int
	}{
		{
			name: "successful creation",
			payload: map[string]interface{}{
				"title": "Test task",
				"done":  false,
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "too short title",
			payload: map[string]interface{}{
				"title": "ab",
				"done":  false,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name: "empty title",
			payload: map[string]interface{}{
				"title": "",
				"done":  false,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/tasks", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			handler.CreateTask(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedStatus == http.StatusCreated {
				var task storage.Task
				if err := json.Unmarshal(rr.Body.Bytes(), &task); err != nil {
					t.Errorf("failed to parse response: %v", err)
				}
				if task.Title != tt.payload.(map[string]interface{})["title"] {
					t.Errorf("expected title %s, got %s", tt.payload.(map[string]interface{})["title"], task.Title)
				}
			}
		})
	}
}

func TestHandlers_ListTasks(t *testing.T) {
	store := storage.NewMemoryStore()
	handler := NewHandlers(store)

	// Добавляем тестовые данные
	handler.Store.Create("Task 1")
	handler.Store.Create("Task 2")
	handler.Store.Create("Different task")

	tests := []struct {
		name           string
		query          string
		expectedCount  int
		expectedStatus int
	}{
		{
			name:           "get all tasks",
			query:          "",
			expectedCount:  3,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "filter tasks",
			query:          "?q=diff",
			expectedCount:  1,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "no matches",
			query:          "?q=nonexistent",
			expectedCount:  0,
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/tasks"+tt.query, nil)
			rr := httptest.NewRecorder()

			handler.ListTasks(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			var tasks []storage.Task
			if err := json.Unmarshal(rr.Body.Bytes(), &tasks); err != nil {
				t.Errorf("failed to parse response: %v", err)
			}

			if len(tasks) != tt.expectedCount {
				t.Errorf("expected %d tasks, got %d", tt.expectedCount, len(tasks))
			}
		})
	}
}
