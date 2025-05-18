package handler

import (
	"net/http"
	"time"

	"github.com/chhongzh/chz_Base_Backend/internal/cacher"
	"github.com/chhongzh/chz_Base_Backend/internal/service/action"
	"github.com/chhongzh/chz_Base_Backend/internal/service/announcement"
	"github.com/chhongzh/chz_Base_Backend/internal/service/application"
	"github.com/chhongzh/chz_Base_Backend/internal/service/permission"
	"github.com/chhongzh/chz_Base_Backend/internal/service/security"
	"github.com/chhongzh/chz_Base_Backend/internal/service/sign"
	"github.com/chhongzh/chz_Base_Backend/internal/service/user"
	"github.com/chhongzh/chz_Base_Backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler struct {
	db         *gorm.DB
	gin        *gin.Engine
	logger     *zap.Logger
	cacher     *cacher.Cacher
	wsUpgrader *websocket.Upgrader

	// Service
	userService         *user.Service
	securityService     *security.Service
	applicationService  *application.Service
	actionService       *action.Service
	permissionService   *permission.Service
	announcementService *announcement.Service
	signService         *sign.Service

	commit       string
	startAt      time.Time
	maxUserCount int64
	corsOrigins  []string
}

func New(db *gorm.DB, gin *gin.Engine, logger *zap.Logger, commit string, prefix string, cacher *cacher.Cacher, maxUserCount int64, corsOrigins []string, maxSignSessionCount int, maxSignSessionWaitingCount int) *Handler {
	secret, _ := utils.RandomBytes(32)

	h := &Handler{
		db:         db,
		gin:        gin,
		logger:     logger,
		cacher:     cacher,
		wsUpgrader: &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }},

		userService:         user.New(db, logger),
		securityService:     security.New(logger, db, secret),
		applicationService:  application.New(db, logger),
		actionService:       action.New(logger, db),
		permissionService:   permission.New(db, logger),
		announcementService: announcement.NewService(logger, db),
		signService:         sign.New(logger, maxSignSessionCount, maxSignSessionWaitingCount),

		commit:       commit,
		startAt:      time.Now(),
		maxUserCount: maxUserCount,
		corsOrigins:  corsOrigins,
	}

	// 初始化中间件
	h.setupCors()

	// 初始化路由
	h.setupRoutes(prefix)

	return h
}
