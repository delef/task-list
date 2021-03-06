package router

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/kulti/task-list/server/internal/models"
)

type sprintHandler struct {
	sprintStore sprintStore
	tmplService sprintTemplateService
}

func newSprintHandler(sprintStore sprintStore, tmplService sprintTemplateService) sprintHandler {
	return sprintHandler{
		sprintStore: sprintStore,
		tmplService: tmplService,
	}
}

func (h sprintHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var sprintID string
	sprintID, r.URL.Path = shiftPath(r.URL.Path)

	if sprintID == "" {
		h.handleCreateSprint(w, r)
		return
	}

	var action string
	action, r.URL.Path = shiftPath(r.URL.Path)
	switch action {
	case "":
		h.handleGetTaskList(w, r, sprintID)
	case "add":
		h.handleCreateTaskInSprint(w, r, sprintID)
	default:
		http.NotFound(w, r)
	}
}

func (h sprintHandler) handleCreateSprint(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	jsDecoder := json.NewDecoder(r.Body)

	var opts struct {
		Begin string `json:"begin"`
		End   string `json:"end"`
	}

	err := jsDecoder.Decode(&opts)
	if err != nil {
		httpBadRequest(w, "failed to parse new sprint body", err)
		return
	}

	begin, err := time.Parse("2006-01-02", opts.Begin)
	if err != nil {
		httpBadRequest(w, "failed to parse begin time", err)
		return
	}

	end, err := time.Parse("2006-01-02", opts.End)
	if err != nil {
		httpBadRequest(w, "failed to parse end time", err)
		return
	}

	err = h.sprintStore.NewSprint(r.Context(), begin, end)
	if err != nil {
		httpInternalServerError(w, "failed to create new sprint", err)
		return
	}

	tmpl, err := h.tmplService.Get(r.Context(), begin, end)
	if err != nil {
		zap.L().Warn("failed to get sprint template - skip it", zap.Error(err))
	}

	httpJSON(w, &tmpl)
}

func (h sprintHandler) handleCreateTaskInSprint(w http.ResponseWriter, r *http.Request, sprintID string) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	task, err := h.parseTask(r.Body)
	if err != nil {
		httpBadRequest(w, "failed to parse task body", err)
		return
	}

	id, err := h.sprintStore.CreateTask(r.Context(), task, sprintID)
	if err != nil {
		httpInternalServerError(w, "failed to create task", err)
		return
	}

	task.ID = id
	httpJSON(w, &task)
}

func (h sprintHandler) parseTask(r io.Reader) (models.Task, error) {
	jsDecoder := json.NewDecoder(r)

	var task models.Task
	err := jsDecoder.Decode(&task)
	if err != nil {
		return models.Task{}, err
	}
	if task.Text == "" {
		return models.Task{}, errMissingArgText
	}
	if task.Points == 0 {
		return models.Task{}, errMissingArgPoints
	}

	return task, nil
}

func (h sprintHandler) handleGetTaskList(w http.ResponseWriter, r *http.Request, sprintID string) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	taskList, err := h.sprintStore.ListTasks(r.Context(), sprintID)
	if err != nil {
		httpInternalServerError(w, "failed to get task list from db", err)
		return
	}

	httpJSON(w, &taskList)
}
