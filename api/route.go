package api

// RouteRule is 路由规则
const RouteRule = "/:endpoint/:name"

// RouteParam is 路由参宿
type RouteParam struct {
	Endpoint string `uri:"endpoint"`
	Name     string `uri:"name"`
}
