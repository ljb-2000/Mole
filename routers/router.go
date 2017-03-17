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
	beego.Router("/api/shards/:keyspace", &controllers.ShardsController{})
	beego.Router("/api/tablets", &controllers.TabletsController{})
	beego.Router("/api/schema/apply", &controllers.SchemaController{}, "post:Apply")

	// K8s
	beego.Router("/api/k8sns/:namespace/pods", &controllers.PodController{})
	beego.Router("/api/k8sns/:namespace/pods/:name", &controllers.PodController{}, "get:Detail;delete:Delete")
	beego.Router("/api/k8sns/:namespace/services", &controllers.ServicesController{})
	beego.Router("/api/k8sns/:namespace/services/:name", &controllers.ServicesController{}, "get:Detail;delete:Delete")
	beego.Router("/api/k8sns/:namespace/rc", &controllers.RcController{})
	beego.Router("/api/k8sns/:namespace/rc/:name", &controllers.RcController{}, "get:Detail;delete:Delete")
	beego.Router("/api/k8sns", &controllers.NamespacesController{})
	beego.Router("/api/k8sns/:name", &controllers.NamespacesController{}, "get:Detail;delete:Delete")

}
