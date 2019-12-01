package controller

import (
	"net/http"
	"net/url"

	"github.com/cortago/cortago/controller"
	"github.com/cortago/cortago/router"
)

type SecondController struct {
	controller.Controller
}

var Second = router.New()

func init(){
	secondController := SecondController{}
	secondController.Name = "Second Controller"
	Second.Handle("/index", secondController.Index)
	Second.Handle("/list", secondController.List)
}

func (c *SecondController) Index(w http.ResponseWriter, r *http.Request, params url.Values) {
	c.Initiliaze(w, r)
	c.RenderText(c.Name + " index is working")
	c.Terminate()
}

func (c *SecondController) List(w http.ResponseWriter, r *http.Request, params url.Values) {
	c.Initiliaze(w, r)
	c.RenderText(c.Name + " list is working")
	c.Terminate()
}
