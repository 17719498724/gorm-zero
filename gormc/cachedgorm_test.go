package gormc

import (
	"context"
	"testing"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mathx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	logx.Disable()
	stat.SetReporter(nil)
}

type sqlitecfg struct {
	DSN string // SQLite DSN
}

func (m *sqlitecfg) Dsn() string {
	return m.DSN
}
func TestGormc_QueryWithExpire(t *testing.T) {

	cfg := sqlitecfg{
		DSN: "file::memory:?cache=shared",
	}
	db, err := gorm.Open(sqlite.Open(cfg.Dsn()), &gorm.Config{})
	if err != nil {
		t.Error(err)
		return
	}
	ccf := cache.CacheConf{
		cache.NodeConf{
			RedisConf: redis.RedisConf{
				Host: "127.0.0.1:6379",
				Pass: "",
			},
			Weight: 100,
		},
	}
	gormc := NewConn(db, ccf)
	var str string
	err = gormc.QueryWithExpireCtx(context.Background(), &str, "any", time.Second*5, func(conn *gorm.DB, v interface{}) error {
		*v.(*string) = "value"
		return nil
	})
	if err != nil {
		t.Error(err)
		return
	}

	err = gormc.QueryWithCallbackExpireCtx(context.Background(), &str, "any", func(conn *gorm.DB, v interface{}) error {
		*v.(*string) = "value"
		return nil
	}, func(i interface{}) time.Duration {
		return time.Second * 5
	})

}

func TestUnstable(t *testing.T) {

	unstable := mathx.NewUnstable(0.1)
	t.Logf("unstable: %v", unstable.AroundDuration(5*time.Minute))
	t.Logf("unstable: %v", unstable.AroundDuration(5*time.Minute))
	t.Logf("unstable: %v", unstable.AroundDuration(5*time.Minute))
	t.Logf("unstable: %v", unstable.AroundDuration(5*time.Minute))
	t.Logf("unstable: %v", unstable.AroundDuration(5*time.Minute))

}
