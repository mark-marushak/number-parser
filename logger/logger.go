package logger

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type loggerService struct {
	repository Logger
}

func New(repo Logger) Logger {
	return &loggerService{repository: repo}
}

func (l loggerService) Info(args ...interface{}) {
	l.repository.Info(args)
}

func (l loggerService) Error(args ...interface{}) {
	l.repository.Error(args)
}
