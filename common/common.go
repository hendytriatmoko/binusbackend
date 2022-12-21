package common

import (
	"encoding/json"
	"os"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

// Configuration stores setting values
type Configuration struct {
	Port                string `json:"port"`
	EnableGinConsoleLog bool   `json:"enableGinConsoleLog"`
	EnableGinFileLog    bool   `json:"enableGinFileLog"`

	LogFilename   string `json:"logFilename"`
	LogMaxSize    int    `json:"logMaxSize"`
	LogMaxBackups int    `json:"logMaxBackups"`
	LogMaxAge     int    `json:"logMaxAge"`

	DbAddrsWarnaCRM       string `json:"db_addrs_warnacrm"`
	DbAddrsWarnaCRMPublic string `json:"db_addrs_warnacrm_public"`
	DbNameWarnaCRM        string `json:"db_name_warnacrm"`
	DbUsernameWarnaCRM    string `json:"db_username_warnacrm"`
	DbPasswordWarnaCRM    string `json:"db_password_warnacrm"`
	DBPortWarnaCRM        string `json:"db_port_warnacrm"`

	JwtSecretPassword string `json:"jwtSecretPassword"`
	Issuer            string `json:"issuer"`

	//ServiceHelper  string `json:"service_helper"`
	//ServiceImage   string `json:"service_image"`
	//ServicePublish string `json:"service_publish"`
	//ServiceMaster  string `json:"service_master"`
	//ServiceLog     string `json:"service_log"`
	//
	//UrlAuth string `json:"url_auth"`

	SMTP_HOST string `json:"smtp_host"`
	SMTP_PORT int    `json:"smtp_port"`
	EMAIL     string `json:"email"`
	PASSWORD  string `json:"password"`
}

// Config shares the global configuration
var (
	Config *Configuration
)

const (
	StatusErrCreate      = "Gagal Menyimpan"
	StatusErrDelete      = "Gagal Menghapus"
	StatusErrUpdate      = "Gagal Mengedit"
	StatusSuccCreate     = "Berhasil Menyimpan"
	StatusSukses         = "Sukses"
	StatusGetNoData      = "Tidak ada data"
	StatusGetExistedData = "Sudah ada data"
	Error                = "Error"
)

// LoadConfig loads configuration from the config file
func LoadConfig() error {
	// Filename is the path to the json config file
	file, err := os.Open("config/config.json")
	if err != nil {
		return err
	}

	Config = new(Configuration)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	// Setting Service Logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   Config.LogFilename,
		MaxSize:    Config.LogMaxSize,    // megabytes after which new file is created
		MaxBackups: Config.LogMaxBackups, // number of backups
		MaxAge:     Config.LogMaxAge,     // days
	})
	log.SetLevel(log.DebugLevel)

	// log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}
