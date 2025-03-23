package app

import (
	"os"

	"github.com/Skapar/simple-rest/config"
	"github.com/Skapar/simple-rest/pkg"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

type App struct {
	Log *zap.SugaredLogger
	DB  *gorm.DB
	Cfg *config.Config
}

func NewApp() *App {
	godotenv.Load()
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "trace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
	}
	zlogger := zap.New(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zap.DebugLevel),
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	log := zlogger.Sugar()
	cfg := config.New()
	cfg.Init()

	dbService := &pkg.GormDatabase{}
	db, err := dbService.Connect(cfg, log)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	return &App{
		Log: log,
		DB:  db,
		Cfg: cfg,
	}
}
