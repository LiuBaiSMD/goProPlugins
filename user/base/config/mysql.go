package config

// MysqlConfig mysql 配置 接口
type MysqlConfig interface {
	GetURL() string
	GetDockerURL() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
}

// defaultMysqlConfig mysql 配置
type defaultMysqlConfig struct {
	URL               string `json:"url"`
	DockerUrl         string `json:"docker_url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

// URL mysql 连接
func (m defaultMysqlConfig) GetURL() string {
	return m.URL
}

// docker环境下的mysql连接
func (m defaultMysqlConfig) GetDockerURL() string {
	return m.DockerUrl
}

// Enabled 激活
func (m defaultMysqlConfig) GetEnabled() bool {
	return m.Enable
}

// 闲置连接数
func (m defaultMysqlConfig) GetMaxIdleConnection() int {
	return m.MaxIdleConnection
}

// 打开连接数
func (m defaultMysqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}
