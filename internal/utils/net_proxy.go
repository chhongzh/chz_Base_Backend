package utils

import (
	"net"
	"time"

	"go.uber.org/zap"
)

type ListenerProxy struct {
	proxyTo net.Listener
	logger  *zap.Logger

	aliveCount uint64
}

type ConnProxy struct {
	proxyTo net.Conn
	logger  *zap.Logger

	totalWriteByteCount uint64
	totalReadByteCount  uint64
}

// Type Guard
var _ net.Listener = (*ListenerProxy)(nil)
var _ net.Conn = (*ConnProxy)(nil)

func ProxyListener(oldListener net.Listener, logger *zap.Logger) *ListenerProxy {
	return &ListenerProxy{
		proxyTo: oldListener,
		logger:  logger,
	}
}

func ProxyConn(oldConn net.Conn, logger *zap.Logger) *ConnProxy {
	return &ConnProxy{
		proxyTo: oldConn,
		logger:  logger,
	}
}

// ConnProxy methods

func (c *ConnProxy) Read(b []byte) (n int, err error) {
	n, err = c.proxyTo.Read(b)
	if err != nil {
		c.logger.Error("Read failed", zap.Error(err))
	}

	// 增加计数器
	c.totalReadByteCount += uint64(n)

	return
}

func (c *ConnProxy) Write(b []byte) (n int, err error) {
	n, err = c.proxyTo.Write(b)
	if err != nil {
		c.logger.Error("Write failed", zap.Error(err))
	}

	// 增加计数器
	c.totalWriteByteCount += uint64(n)

	return
}

func (c *ConnProxy) Close() error {
	err := c.proxyTo.Close()
	if err != nil {
		c.logger.Error("Close failed", zap.Error(err))
	}
	return err
}

func (c *ConnProxy) RemoteAddr() net.Addr {
	addr := c.proxyTo.RemoteAddr()
	return addr
}

func (c *ConnProxy) SetDeadline(t time.Time) error {
	err := c.proxyTo.SetDeadline(t)
	if err != nil {
		c.logger.Error("SetDeadline failed", zap.Error(err))
	}
	return err
}

func (c *ConnProxy) SetReadDeadline(t time.Time) error {
	err := c.proxyTo.SetReadDeadline(t)
	if err != nil {
		c.logger.Error("SetReadDeadline failed", zap.Error(err))
	}
	return err
}

func (c *ConnProxy) SetWriteDeadline(t time.Time) error {
	err := c.proxyTo.SetWriteDeadline(t)
	if err != nil {
		c.logger.Error("SetWriteDeadline failed", zap.Error(err))
	}
	return err
}

func (c *ConnProxy) LocalAddr() net.Addr {
	return c.proxyTo.LocalAddr()
}

func (c *ConnProxy) GetTotalWriteByteCount() uint64 {
	return c.totalWriteByteCount
}

func (c *ConnProxy) GetTotalReadByteCount() uint64 {
	return c.totalReadByteCount
}

// ListenerProxy methods

func (l *ListenerProxy) Accept() (net.Conn, error) {
	conn, err := l.proxyTo.Accept()
	if err != nil {
		l.logger.Error("Accept failed", zap.Error(err))
		return nil, err
	}

	// 计数信息
	l.aliveCount++

	proxyConn := ProxyConn(conn, l.logger.Named("Proxy Conn for \""+conn.RemoteAddr().String()+"\""))
	return proxyConn, nil
}

func (l *ListenerProxy) Close() error {
	// 输出
	l.logger.Info("Conn closed.")

	err := l.proxyTo.Close()
	if err != nil {
		l.logger.Error("Close failed", zap.Error(err))
	}

	// 计数信息
	l.aliveCount--

	return err
}

func (l *ListenerProxy) Addr() net.Addr {
	return l.proxyTo.Addr()
}

func (l *ListenerProxy) GetAliveCount() uint64 {
	return l.aliveCount
}
