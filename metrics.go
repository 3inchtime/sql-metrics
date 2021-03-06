package sql_metrics
//
//import (
//	"github.com/prometheus/client_golang/prometheus"
//)
//
//var (
//	// DBDurationsSeconds mysql 调用耗时
//	DBDurationsSeconds *prometheus.HistogramVec
//	// NetPoolHits 命中空闲连接数量
//	NetPoolHits *prometheus.CounterVec
//	// NetPoolMisses 未命中空闲连接数量
//	NetPoolMisses *prometheus.CounterVec
//	// NetPoolTimeouts 获取连接超时总数
//	NetPoolTimeouts *prometheus.CounterVec
//	// NetPoolStale 问题连接总数
//	NetPoolStale *prometheus.CounterVec
//	// NetPoolTotal 连接总数
//	NetPoolTotal *prometheus.GaugeVec
//	// NetPoolIdle 空闲连接总数
//	NetPoolIdle *prometheus.GaugeVec
//	// DBMaxOpenConnections 最大 DB 连接数
//	DBMaxOpenConnections *prometheus.GaugeVec
//	// DBOpenConnections 当前 DB 连接总数
//	DBOpenConnections *prometheus.GaugeVec
//	// DBInUseConnections 在用 DB 连接数
//	DBInUseConnections *prometheus.GaugeVec
//	// DBIdleConnections 空闲 DB 连接数
//	DBIdleConnections *prometheus.GaugeVec
//	// DBWaitCount 从 DB 连接池取不到连接需要等待的总数量
//	DBWaitCount *prometheus.CounterVec
//	// DBMaxIdleClosed 因为 SetMaxIdleConns 而被关闭的连接总数量
//	DBMaxIdleClosed *prometheus.CounterVec
//	// DBMaxLifetimeClosed 因为 SetConnMaxLifetime 而被关闭的连接总数量
//	DBMaxLifetimeClosed *prometheus.CounterVec
//
//)
//
//var defBuckets = []float64{.005, .01, .025, .05, .1, .25, .5, 1}
//
//func init() {
//	AppID := config.Cfg.GetString("appid")
//	if AppID == ""{
//		AppID = "UnknownService"
//	}
//
//	NetPoolHits = prometheus.NewCounterVec(prometheus.CounterOpts{
//		Namespace:   "kayle",
//		Name:        "net_pool_hits",
//		Help:        "net pool hits",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name", "type"})
//	prometheus.MustRegister(NetPoolHits)
//
//	NetPoolMisses = prometheus.NewCounterVec(prometheus.CounterOpts{
//		Namespace:   "kayle",
//		Name:        "net_pool_misses",
//		Help:        "net pool misses",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name", "type"})
//	prometheus.MustRegister(NetPoolMisses)
//
//	NetPoolTimeouts = prometheus.NewCounterVec(prometheus.CounterOpts{
//		Namespace:   "kayle",
//		Name:        "net_pool_timeouts",
//		Help:        "net pool timeouts",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name", "type"})
//	prometheus.MustRegister(NetPoolTimeouts)
//
//	NetPoolStale = prometheus.NewCounterVec(prometheus.CounterOpts{
//		Namespace:   "kayle",
//		Name:        "net_pool_stale",
//		Help:        "net pool stale",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name", "type"})
//	prometheus.MustRegister(NetPoolStale)
//
//	NetPoolTotal = prometheus.NewGaugeVec(prometheus.GaugeOpts{
//		Namespace:   "kayle",
//		Name:        "net_pool_total",
//		Help:        "net pool total",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name", "type"})
//	prometheus.MustRegister(NetPoolTotal)
//
//	NetPoolIdle = prometheus.NewGaugeVec(prometheus.GaugeOpts{
//		Namespace:   "kayle",
//		Name:        "net_pool_idle",
//		Help:        "net pool idle",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name", "type"})
//	prometheus.MustRegister(NetPoolIdle)
//
//	DBMaxOpenConnections = prometheus.NewGaugeVec(prometheus.GaugeOpts{
//		Namespace:   "kayle",
//		Name:        "db_max_open_conns",
//		Help:        "db max open connections",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name"})
//	prometheus.MustRegister(DBMaxOpenConnections)
//
//	DBOpenConnections = prometheus.NewGaugeVec(prometheus.GaugeOpts{
//		Namespace:   "kayle",
//		Name:        "db_open_conns",
//		Help:        "db open connections",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name"})
//	prometheus.MustRegister(DBOpenConnections)
//
//	DBInUseConnections = prometheus.NewGaugeVec(prometheus.GaugeOpts{
//		Namespace:   "kayle",
//		Name:        "db_in_use_conns",
//		Help:        "db in use connections",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name"})
//	prometheus.MustRegister(DBInUseConnections)
//
//	DBIdleConnections = prometheus.NewGaugeVec(prometheus.GaugeOpts{
//		Namespace:   "kayle",
//		Name:        "db_idle_conns",
//		Help:        "db idle connections",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name"})
//	prometheus.MustRegister(DBIdleConnections)
//
//	DBWaitCount = prometheus.NewCounterVec(prometheus.CounterOpts{
//		Namespace:   "kayle",
//		Name:        "db_wait_count",
//		Help:        "db wait count",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name"})
//	prometheus.MustRegister(DBWaitCount)
//
//	DBMaxIdleClosed = prometheus.NewCounterVec(prometheus.CounterOpts{
//		Namespace:   "kayle",
//		Name:        "db_max_idle_closed",
//		Help:        "db max idle closed",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name"})
//	prometheus.MustRegister(DBMaxIdleClosed)
//
//	DBMaxLifetimeClosed = prometheus.NewCounterVec(prometheus.CounterOpts{
//		Namespace:   "kayle",
//		Name:        "db_max_life_time_closed",
//		Help:        "db max life time closed",
//		ConstLabels: map[string]string{"app": AppID},
//	}, []string{"name"})
//	prometheus.MustRegister(DBMaxLifetimeClosed)
//}
