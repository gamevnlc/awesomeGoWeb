package config

import (
	"awesomeWeb/internal/models"
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	ErrorLog      *log.Logger
	MailChan      chan models.MailData
}
