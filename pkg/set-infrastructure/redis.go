package setinfrastructure

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
	modtype "github.com/synzofficial/nsd-synz-sharelib-api/pkg/mod-type"
)

/*
	Example: local.env
	REDIS_ADDR="0.0.0.0:6379"
	REDIS_DB="0"
	REDIS_USERNAME=""
	REDIS_PASSWORD=""
	REDIS_READ_TIMEOUT="10s"
	REDIS_WRITE_TIMEOUT="10s"
	REDIS_CONTEXT_TIMEOUT_ENABLED="true"
	REDIS_POOL_FIFO="false"
	REDIS_POOL_SIZE="10"
	REDIS_POOL_TIMEOUT="30s"
	REDIS_MIN_IDLE_CONNS="30"
	REDIS_MAX_ACTIVE_CONNS="0"
	REDIS_CONN_MAX_IDLE_TIME="30m"
	REDIS_CONN_MAX_LIFETIME="0m"
*/

type RedisConfig struct {
	// host:port address.
	Addr string `envconfig:"REDIS_ADDR" description:"" example:"host:port" default:"localhost:6379"`
	// Database to be selected after connecting to the server.
	DB int `envconfig:"REDIS_DB" description:"Database to be selected after connecting to the server." default:"0"`
	// Use the specified Username to authenticate the current connection
	// with one of the connections defined in the ACL list when connecting
	// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
	Username string `envconfig:"REDIS_USERNAME" description:"Use the specified Username to authenticate"`
	// Optional password. Must match the password specified in the
	// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
	// or the User Password when connecting to a Redis 6.0 instance, or greater,
	// that is using the Redis ACL system.
	Password string `envconfig:"REDIS_PASSWORD" description:"Use the specified Password to authenticate"`
	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Supported values:
	//   - `0` - default timeout (3 seconds).
	//   - `-1` - no timeout (block indefinitely).
	//   - `-2` - disables SetReadDeadline calls completely.
	ReadTimeout modtype.Duration `envconfig:"REDIS_READ_TIMEOUT" description:"Timeout for socket read. [0=default timeout(3 seconds), -1=no timeout, -2=disables SetReadDeadline calls completely]" default:"3s"`
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.  Supported values:
	//   - `0` - default timeout (3 seconds).
	//   - `-1` - no timeout (block indefinitely).
	//   - `-2` - disables SetWriteDeadline calls completely.
	WriteTimeout modtype.Duration `envconfig:"REDIS_WRITE_TIMEOUT" description:"Timeout for socket write. [0=default timeout(3 seconds), -1=no timeout, -2=disables SetReadDeadline calls completely]" default:"3s"`
	// ContextTimeoutEnabled controls whether the client respects context timeouts and deadlines.
	ContextTimeoutEnabled bool `envconfig:"REDIS_CONTEXT_TIMEOUT_ENABLED" description:"ContextTimeoutEnabled controls whether the client respects context timeouts and deadlines." default:"true"`
	// Type of connection pool.
	// true for FIFO pool, false for LIFO pool.
	// Note that FIFO has slightly higher overhead compared to LIFO,
	// but it helps closing idle connections faster reducing the pool size.
	PoolFIFO bool `envconfig:"REDIS_POOL_FIFO" description:"true for FIFO pool, false for LIFO pool Note that FIFO has slightly higher overhead compared to LIFO, but it helps closing idle connections faster reducing the pool size."`
	// Base number of socket connections.
	// Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
	// If there is not enough connections in the pool, new connections will be allocated in excess of PoolSize,
	// you can limit it through MaxActiveConns
	PoolSize int `envconfig:"REDIS_POOL_SIZE" description:"Base number of socket connections. Default is 10 connections per every available CPU as reported by runtime." default:"10"`
	// Amount of time client waits for connection if all connections
	// are busy before returning an error.
	// Default is ReadTimeout + 1 second.
	PoolTimeout modtype.Duration `envconfig:"REDIS_POOL_TIMEOUT" description:"Amount of time client waits for connection if all connections are busy before returning an error." default:"30s"`
	// Minimum number of idle connections which is useful when establishing
	// new connection is slow.
	// Default is 0. the idle connections are not closed by default.
	MinIdleConns int `envconfig:"REDIS_MIN_IDLE_CONNS" description:"Amount of time client waits for connection if all connections are busy before returning an error." default:"0"`
	// Maximum number of connections allocated by the pool at a given time.
	// When zero, there is no limit on the number of connections in the pool.
	MaxActiveConns int `envconfig:"REDIS_MAX_ACTIVE_CONNS" default:""`
	// ConnMaxIdleTime is the maximum amount of time a connection may be idle.
	// Should be less than server's timeout.
	//
	// Expired connections may be closed lazily before reuse.
	// If d <= 0, connections are not closed due to a connection's idle time.
	//
	// Default is 30 minutes. -1 disables idle timeout check.
	ConnMaxIdleTime modtype.Duration `envconfig:"REDIS_CONN_MAX_IDLE_TIME" default:"30m"`
	// ConnMaxLifetime is the maximum amount of time a connection may be reused.
	//
	// Expired connections may be closed lazily before reuse.
	// If <= 0, connections are not closed due to a connection's age.
	//
	// Default is to not close idle connections.
	ConnMaxLifetime modtype.Duration `envconfig:"REDIS_CONN_MAX_LIFETIME"`
}

func NewRedis(ctx context.Context, cfg RedisConfig) (*redis.Client, error) {
	opt := redis.Options{
		Addr:                  cfg.Addr,
		DB:                    cfg.DB,
		Username:              cfg.Username,
		Password:              cfg.Password,
		ReadTimeout:           time.Duration(cfg.ReadTimeout),
		WriteTimeout:          time.Duration(cfg.WriteTimeout),
		ContextTimeoutEnabled: cfg.ContextTimeoutEnabled,
		PoolFIFO:              cfg.PoolFIFO,
		PoolSize:              cfg.PoolSize,
		PoolTimeout:           time.Duration(cfg.PoolTimeout),
		MinIdleConns:          cfg.MinIdleConns,
		MaxActiveConns:        cfg.MaxActiveConns,
		ConnMaxIdleTime:       time.Duration(cfg.ConnMaxIdleTime),
		ConnMaxLifetime:       time.Duration(cfg.ConnMaxLifetime),
	}
	rc := redis.NewClient(&opt)
	if err := rc.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rc, nil
}
