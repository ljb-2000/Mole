package routers

import (
	"github.com/astaxie/beego"
	"github.com/zssky/Mole/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// Vitess
	beego.Router("/api/keyspaces", &controllers.KeyspacesController{})
	beego.Router("/api/vtctl", &controllers.VtctlController{})
	beego.Router("/api/shards", &controllers.ShardsController{})
	beego.Router("/api/tablets", &controllers.TabletsController{})
	beego.Router("/api/schema", &controllers.SchemaController{})

	// K8s
	beego.Router("/api/namespaces/{namespace}/pods", &controllers.PodController{})
	beego.Router("/api/namespaces/{namespace}/services", &controllers.ServicesController{})
	beego.Router("/api/namespaces/{namespace}/replicationcontrollers", &controllers.RcController{})
	beego.Router("/api/namespaces", &controllers.NamespacesController{})

}
