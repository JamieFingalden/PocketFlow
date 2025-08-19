package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config 配置结构体
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
	OCR      OCRConfig      `yaml:"ocr"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Name       string `yaml:"name"`
	Charset    string `yaml:"charset"`
	ParseTime  bool   `yaml:"parseTime"`
	Loc        string `yaml:"loc"`
	Dsn        string `yaml:"-"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}

// OCRConfig OCR服务配置
type OCRConfig struct {
	ServiceURL string `yaml:"service_url"`
}

// LoadConfig 加载配置文件
func LoadConfig() *Config {
	// 尝试不同的配置文件路径
	configPaths := []string{
		"config/config.yaml",
		"../config/config.yaml",
		"../../config/config.yaml",
	}
	
	var data []byte
	var err error
	
	// 尝试读取配置文件
	for _, path := range configPaths {
		data, err = os.ReadFile(path)
		if err == nil {
			break
		}
	}
	
	if err != nil {
		// 如果在所有路径都找不到配置文件，尝试从当前工作目录搜索
		err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == "config.yaml" && filepath.Base(filepath.Dir(path)) == "config" {
				data, err = os.ReadFile(path)
				return err
			}
			return nil
		})
		
		if err != nil {
			log.Fatalf("Failed to read config file: %v", err)
		}
	}

	// 解析YAML
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	// 构建DSN
	config.Database.Dsn = buildDSN(&config.Database)

	return &config
}

// buildDSN 构建数据库连接字符串
func buildDSN(db *DatabaseConfig) string {
	return db.User + ":" + db.Password + "@tcp(" + db.Host + ":" + db.Port + ")/" + db.Name + "?charset=" + db.Charset + "&parseTime=" + 
		map[bool]string{true: "True", false: "False"}[db.ParseTime] + "&loc=" + db.Loc
}