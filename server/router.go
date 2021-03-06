package server

import (
	"github.com/baetyl/baetyl-cloud/common"
	"github.com/baetyl/baetyl-go/log"
	"github.com/gin-gonic/gin"
)

// InitRoute init router
func (s *AdminServer) InitRoute() {
	s.router.NoRoute(noRouteHandler)
	s.router.NoMethod(noMethodHandler)
	s.router.GET("/health", health)

	s.router.Use(requestIDHandler)
	s.router.Use(loggerHandler)
	s.router.Use(s.authHandler)
	v1 := s.router.Group("v1")
	{
		configs := v1.Group("/configs")
		configs.GET("/:name", common.Wrapper(s.api.GetConfig))
		configs.PUT("/:name", common.Wrapper(s.api.UpdateConfig))
		configs.DELETE("/:name", common.Wrapper(s.api.DeleteConfig))
		configs.POST("", common.Wrapper(s.api.CreateConfig))
		configs.GET("", common.Wrapper(s.api.ListConfig))
		configs.GET("/:name/apps", common.Wrapper(s.api.GetAppByConfig))
	}
	{
		registry := v1.Group("/registries")
		registry.GET("/:name", common.Wrapper(s.api.GetRegistry))
		registry.PUT("/:name", common.Wrapper(s.api.UpdateRegistry))
		registry.POST("/:name/refresh", common.Wrapper(s.api.RefreshRegistryPassword))
		registry.DELETE("/:name", common.Wrapper(s.api.DeleteRegistry))
		registry.POST("", common.Wrapper(s.api.CreateRegistry))
		registry.GET("", common.Wrapper(s.api.ListRegistry))
		registry.GET("/:name/apps", common.Wrapper(s.api.GetAppByRegistry))
	}
	{
		configs := v1.Group("/secrets")
		configs.GET("/:name", common.Wrapper(s.api.GetSecret))
		configs.PUT("/:name", common.Wrapper(s.api.UpdateSecret))
		configs.DELETE("/:name", common.Wrapper(s.api.DeleteSecret))
		configs.POST("", common.Wrapper(s.api.CreateSecret))
		configs.GET("", common.Wrapper(s.api.ListSecret))
		configs.GET("/:name/apps", common.Wrapper(s.api.GetAppBySecret))
	}
	{
		nodes := v1.Group("/nodes")
		nodes.GET("/:name", common.Wrapper(s.api.GetNode))
		nodes.GET("/:name/apps", common.Wrapper(s.api.GetAppByNode))
		nodes.GET("/:name/stats", common.Wrapper(s.api.GetNodeStats))
		nodes.PUT("/:name", common.Wrapper(s.api.UpdateNode))
		nodes.DELETE("/:name", common.Wrapper(s.api.DeleteNode))
		nodes.POST("", s.nodeQuotaHandler, common.Wrapper(s.api.CreateNode))
		nodes.GET("", common.Wrapper(s.api.ListNode))
		nodes.GET("/:name/deploys", common.Wrapper(s.api.GetNodeDeployHistory))
		nodes.GET("/:name/init", common.Wrapper(s.api.GenInitCmdFromNode))
	}
	{
		apps := v1.Group("/apps")
		apps.GET("/:name", common.Wrapper(s.api.GetApplication))
		apps.PUT("/:name", common.Wrapper(s.api.UpdateApplication))
		apps.DELETE("/:name", common.Wrapper(s.api.DeleteApplication))
		apps.POST("", common.Wrapper(s.api.CreateApplication))
		apps.GET("", common.Wrapper(s.api.ListApplication))
	}
	{
		register := v1.Group("/register")
		register.GET("", common.Wrapper(s.api.ListBatch))
		register.POST("", common.Wrapper(s.api.CreateBatch))
		register.PUT("/:batchName", common.Wrapper(s.api.UpdateBatch))
		register.DELETE("/:batchName", common.Wrapper(s.api.DeleteBatch))
		register.GET("/:batchName", common.Wrapper(s.api.GetBatch))
		register.GET("/:batchName/init", common.Wrapper(s.api.GenInitCmdFromBatch))

		register.GET("/:batchName/download", common.WrapperRaw(s.api.DownloadRecords))
		register.POST("/:batchName/generate", common.Wrapper(s.api.GenRecordRandom))

		register.POST("/:batchName/record", common.Wrapper(s.api.CreateRecord))
		register.PUT("/:batchName/record/:recordName", common.Wrapper(s.api.UpdateRecord))
		register.DELETE("/:batchName/record/:recordName", common.Wrapper(s.api.DeleteRecord))
		register.GET("/:batchName/record/:recordName", common.Wrapper(s.api.GetRecord))
		register.GET("/:batchName/record", common.Wrapper(s.api.ListRecord))
	}
	{
		callback := v1.Group("/callback")
		callback.POST("", common.Wrapper(s.api.CreateCallback))
		callback.PUT("/:callbackName", common.Wrapper(s.api.UpdateCallback))
		callback.DELETE("/:callbackName", common.Wrapper(s.api.DeleteCallback))
		callback.GET("/:callbackName", common.Wrapper(s.api.GetCallback))

	}
	{
		namespace := v1.Group("/namespace")
		namespace.POST("", common.Wrapper(s.api.CreateNamespace))
		namespace.GET("", common.Wrapper(s.api.GetNamespace))
		namespace.DELETE("", common.Wrapper(s.api.DeleteNamespace))
	}
	{
		sysconfig := v1.Group("/sysconfig")
		sysconfig.GET("/:type/:key", common.Wrapper(s.api.GetSysConfig))
		sysconfig.GET("/:type", common.Wrapper(s.api.ListSysConfig))

	}
	{
		function := v1.Group("/functions")
		function.GET("", common.Wrapper(s.api.ListFunctionSources))
		if len(s.cfg.Plugin.Functions) != 0 {
			function.GET("/:source/functions", common.Wrapper(s.api.ListFunctions))
			function.GET("/:source/functions/:name/versions", common.Wrapper(s.api.ListFunctionVersions))
			function.POST("/:source/functions/:name/versions/:version", common.Wrapper(s.api.ImportFunction))
		}
	}
	{
		objects := v1.Group("/objects")
		objects.GET("", common.Wrapper(s.api.ListObjectSources))
		if len(s.cfg.Plugin.Objects) != 0 {
			objects.GET("/:source/buckets", common.Wrapper(s.api.ListBuckets))
			objects.GET("/:source/buckets/:bucket/objects", common.Wrapper(s.api.ListBucketObjects))
		}
	}
}

// GetRoute get router
func (s *AdminServer) GetRoute() *gin.Engine {
	return s.router
}

// auth handler
func (s *AdminServer) authHandler(c *gin.Context) {
	cc := common.NewContext(c)
	err := s.auth.Authenticate(cc)
	if err != nil {
		log.L().Error("request authenticate failed",
			log.Any(cc.GetTrace()),
			log.Any("namespace", cc.GetNamespace()),
			log.Any("authorization", c.Request.Header.Get("Authorization")),
			log.Error(err))
		common.PopulateFailedResponse(cc, common.Error(common.ErrRequestAccessDenied), true)
	}
}

// access manager handler
func (s *AdminServer) nodeQuotaHandler(c *gin.Context) {
	cc := common.NewContext(c)
	namespace := cc.GetNamespace()
	if err := s.license.CheckQuota(namespace, s.api.NodeNumberCollector); err != nil {
		log.L().Error("iotcore checkquota failed",
			log.Any(cc.GetTrace()),
			log.Any("namespace", cc.GetNamespace()),
			log.Error(err))
		common.PopulateFailedResponse(cc, err, true)
	}
}

// GetRoute get router
func (s *NodeServer) GetRoute() *gin.Engine {
	return s.router
}

func (s *NodeServer) InitRoute() {
	s.router.NoRoute(noRouteHandler)
	s.router.NoMethod(noMethodHandler)
	s.router.GET("/health", health)

	s.router.Use(requestIDHandler)
	s.router.Use(loggerHandler)
	if s.server.TLSConfig == nil {
		HeaderCommonName = s.cfg.NodeServer.CommonName
		s.router.Use(extractNodeCommonNameFromHeader)
	} else {
		s.router.Use(extractNodeCommonNameFromCert)
	}
	v1 := s.router.Group("v1")
	{
		node := v1.Group("/sync")
		node.POST("/report", common.Wrapper(s.api.Report))
		node.POST("/desire", common.Wrapper(s.api.Desire))
	}
}

// GetRoute get router
func (s *ActiveServer) GetRoute() *gin.Engine {
	return s.router
}

func (s *ActiveServer) InitRoute() {
	s.router.NoRoute(noRouteHandler)
	s.router.NoMethod(noMethodHandler)
	s.router.GET("/health", health)

	s.router.Use(requestIDHandler)
	s.router.Use(loggerHandler)
	v1 := s.router.Group("v1")
	{
		active := v1.Group("/active")
		active.POST("/", common.Wrapper(s.api.Active))
		active.GET("/:resource", common.WrapperRaw(s.api.GetResource))
	}
}
