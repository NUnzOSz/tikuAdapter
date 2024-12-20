package configs

import (
	"tikuAdapter/internal/search"

	"github.com/elastic/go-elasticsearch/v8"
)

// Config 所有的配置文件
type Config struct {
	Limit             LimitConfig          `yaml:"limit"`
	API               []search.API         `yaml:"api"`
	Elasticsearch     elasticsearch.Config `yaml:"elasticsearch"`
	RecordEmptyAnswer bool                 `yaml:"recordEmptyAnswer"`
	Mysql             string               `yaml:"mysql"`
	OSS               OSSConfig            `yaml:"oss"`
	Plat              []PlatConfig         `yaml:"plat"`
}

// LimitConfig 限流配置
type LimitConfig struct {
	Enable   bool   `yaml:"enable"`
	Duration uint   `yaml:"duration"`
	Requests uint64 `yaml:"requests"`
}

// OSSConfig 阿里云oss配置
type OSSConfig struct {
	EndPoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
}

// PlatConfig 自定义的平台类型
type PlatConfig struct {
	Label string `yaml:"label"`
	Value string `yaml:"value"`
}
