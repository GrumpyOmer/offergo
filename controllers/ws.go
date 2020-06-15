package controllers

type WebsocketController struct {
	baseController
}

func (w *WebsocketController) Index() {
	w.TplName = "ws.tpl"
}
