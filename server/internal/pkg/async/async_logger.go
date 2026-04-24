package async

import (
	"context"
	"log/slog"
	"sync"
	"time"

	modelSysMonitor "server/internal/model/sysMonitor"
	"server/internal/service/sysMonitor"
)

const (
	// 缓冲通道大小
	bufferSize = 1000
	// 批量插入大小
	batchSize = 100
	// 刷新间隔
	flushInterval = 5 * time.Second
)

// AsyncOperationLogger 异步操作日志记录器
type AsyncOperationLogger struct {
	buffer   chan *modelSysMonitor.OperationLogModel
	wg       sync.WaitGroup
	ctx      context.Context
	cancel   context.CancelFunc
	service  *sysMonitor.OperationLogService
	stopOnce sync.Once
}

var (
	instance *AsyncOperationLogger
	once     sync.Once
)

// GetAsyncLogger 获取单例实例
func GetAsyncLogger() *AsyncOperationLogger {
	once.Do(func() {
		ctx, cancel := context.WithCancel(context.Background())
		instance = &AsyncOperationLogger{
			buffer:  make(chan *modelSysMonitor.OperationLogModel, bufferSize),
			ctx:     ctx,
			cancel:  cancel,
			service: sysMonitor.NewOperationLogService(),
		}
		instance.start()
	})
	return instance
}

// start 启动后台工作协程
func (a *AsyncOperationLogger) start() {
	a.wg.Add(1)
	go a.worker()
}

// worker 后台工作协程，批量处理日志
func (a *AsyncOperationLogger) worker() {
	defer a.wg.Done()

	ticker := time.NewTicker(flushInterval)
	defer ticker.Stop()

	batch := make([]*modelSysMonitor.OperationLogModel, 0, batchSize)

	for {
		select {
		case <-a.ctx.Done():
			// 处理剩余日志
			for len(a.buffer) > 0 {
				select {
				case log := <-a.buffer:
					batch = append(batch, log)
					if len(batch) >= batchSize {
						a.saveBatch(batch)
						batch = batch[:0]
					}
				default:
				}
			}
			if len(batch) > 0 {
				a.saveBatch(batch)
			}
			return

		case log := <-a.buffer:
			batch = append(batch, log)
			if len(batch) >= batchSize {
				a.saveBatch(batch)
				batch = batch[:0]
			}

		case <-ticker.C:
			if len(batch) > 0 {
				a.saveBatch(batch)
				batch = batch[:0]
			}
		}
	}
}

// saveBatch 批量保存日志
func (a *AsyncOperationLogger) saveBatch(logs []*modelSysMonitor.OperationLogModel) {
	if len(logs) == 0 {
		return
	}

	// 逐个保存（如果 repository 支持批量插入，可以优化）
	for _, log := range logs {
		if err := a.service.Create(log); err != nil {
			slog.Error("async save operation log failed", "error", err)
		}
	}
}

// Log 提交日志到异步队列
func (a *AsyncOperationLogger) Log(log *modelSysMonitor.OperationLogModel) {
	select {
	case a.buffer <- log:
		// 成功放入队列
	default:
		// 队列已满，丢弃日志并记录警告
		slog.Warn("async operation log buffer is full, dropping log",
			"path", log.Path,
			"method", log.Method)
	}
}

// Stop 优雅停止，等待所有日志处理完成
func (a *AsyncOperationLogger) Stop() {
	a.stopOnce.Do(func() {
		a.cancel()
		a.wg.Wait()
		close(a.buffer)
	})
}

// BufferLen 获取当前缓冲区长度（用于监控）
func (a *AsyncOperationLogger) BufferLen() int {
	return len(a.buffer)
}
