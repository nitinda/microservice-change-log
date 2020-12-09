package logger

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"
)

// DBLoggerInfo for datbase Info logs
var DBLoggerInfo = logger.New(
	log.New(os.Stdout, "changelog-api - database - ", log.Ldate|log.Ltime|log.Lshortfile), // io writer
	logger.Config{
		SlowThreshold: time.Second, // Slow SQL threshold
		LogLevel:      logger.Info, // Log level
		Colorful:      true,        // Disable color
	},
)

// DBLoggerError for datbase logs
var DBLoggerError = logger.New(
	log.New(os.Stdout, "changelog-api - database - ", log.Ldate|log.Ltime|log.Lshortfile), // io writer
	logger.Config{
		SlowThreshold: time.Second,  // Slow SQL threshold
		LogLevel:      logger.Error, // Log level
		Colorful:      true,         // Disable color
	},
)

// DBLoggerWarn for datbase logs
var DBLoggerWarn = logger.New(
	log.New(os.Stdout, "changelog-api - database - ", log.Ldate|log.Ltime|log.Lshortfile), // io writer
	logger.Config{
		SlowThreshold: time.Second, // Slow SQL threshold
		LogLevel:      logger.Warn, // Log level
		Colorful:      true,        // Disable color
	},
)
