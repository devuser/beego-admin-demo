package rediscache

type RedisConfig struct {
    Host string `json:"redis_server_host"`
}


func NewRedisConfig(host string) *RedisConfig{
    return &RedisConfig{
        Host: host}
}
