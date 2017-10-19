package logs

import (
	"context"
	"io"

	"github.com/fnproject/fn/api/models"
	"github.com/pkg/errors"
)

type mock struct {
	Logs map[string]io.Reader
}

func NewMock() models.LogStore {
	return &mock{make(map[string]io.Reader)}
}

func (m *mock) InsertLog(ctx context.Context, appName, callID string, callLog io.Reader) error {
	m.Logs[callID] = callLog
	return nil
}

func (m *mock) GetLog(ctx context.Context, appName, callID string) (io.Reader, error) {
	logEntry := m.Logs[callID]
	if logEntry == nil {
		return nil, errors.New("Call log not found")
	}

	return logEntry, nil
}

func (m *mock) DeleteLog(ctx context.Context, appName, callID string) error {
	delete(m.Logs, callID)
	return nil
}
