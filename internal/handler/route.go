package handler

func (h *Handler) setupRoutes(prefix string) {
	apiGroup := h.gin.Group(prefix)
	{
		apiGroup.POST("/getCacheInfo", h.getCacheInfo)
		apiGroup.POST("/getSystemInfo", h.getNetworkInfo)
		apiGroup.POST("/reloadCache", h.reloadCache)
		apiGroup.GET("/ws", h.wsOpen)

		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("/register", h.userRegister)
			userGroup.POST("/login", h.userLogin)
		}

		applicationGroup := apiGroup.Group("/application")
		{
			applicationGroup.POST("/create", h.applicationCreate)
			applicationGroup.POST("/delete", h.applicationDelete)
			applicationGroup.POST("/list", h.applicationList)
		}

		announcementGroup := apiGroup.Group("/announcement")
		{
			announcementGroup.POST("/list", h.announcementList)
			announcementGroup.POST("/announce", h.announcementAnnounce)
			announcementGroup.POST("/last", h.announcementLast)
			announcementGroup.POST("/delete", h.announcementDelete)
		}

		signGroup := apiGroup.Group("/sign")
		{
			signGroup.POST("", h.signCreate)

			forSignGroup := signGroup.Group("/:SignSessionID")
			{
				forSignGroup.GET("", h.signWait)
				forSignGroup.POST("/complete", h.signComplete)
				forSignGroup.GET("/info", h.signInfo)
			}

		}

		myselfGroup := apiGroup.Group("/myself")
		{
			myselfGroup.POST("/permission/list", h.myselfPermissionList)
		}
	}

}
