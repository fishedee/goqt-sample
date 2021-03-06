package views

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type TopToolFrame struct {
	*widgets.QFrame
	Model
	parent       widgets.QWidget_ITF
	isDraging    bool
	closeButton  *PushButton
	miniButton   *PushButton
	moveListener func(x int, y int)
	dragPosition *core.QPoint
}

func NewTopToolFrame(parent widgets.QWidget_ITF) *TopToolFrame {
	topToolFrame := TopToolFrame{}
	InitModel(&topToolFrame)
	topToolFrame.init(parent)
	return &topToolFrame
}

func (this *TopToolFrame) init(parent widgets.QWidget_ITF) {
	this.QFrame = widgets.NewQFrame(parent, 0)
	this.parent = parent
	this.SetObjectName("frame")
	this.SetStyleSheet(`
		QFrame#frame{
			border-image:url(res/box.png);
			background-repeat:no-repeat;
		}
	`)
	this.isDraging = false
	this.SetGeometry2(0, 0, 800, 60)
	this.addRobotLogo()
	this.addButtons()
	this.ConnectMousePressEvent(this.pressEvent)
	this.ConnectMouseReleaseEvent(this.releaseEvent)
	this.ConnectMouseMoveEvent(this.moveEvent)
}

func (this *TopToolFrame) addRobotLogo() {
	btn := widgets.NewQPushButton(this.parent)
	btn.SetObjectName("btnSpecial")
	btn.SetStyleSheet(`
		QPushButton#btnSpecial {
        border-image: url(res/robot_1.png);
        background-repeat: no-repeat;
        }
        QPushButton#btnSpecial:hover {
        border-image: url(res/robot_2.png);
        background-repeat: no-repeat;
        }
        QPushButton#btnSpecial:pressed {
        border-image: url(res/robot_3.png);
        background-repeat: no-repeat;
        }
	`)
	btn.SetGeometry2(20, 0, 67, 60)
}

func (this *TopToolFrame) addButtons() {
	this.closeButton = NewPushButton(this)
	this.closeButton.LoadPixmap("res/close.png")
	this.closeButton.SetGeometry2(770, 10, 16, 16)

	this.miniButton = NewPushButton(this)
	this.miniButton.LoadPixmap("res/mini.png")
	this.miniButton.SetGeometry2(740, 10, 16, 16)
}

func (this *TopToolFrame) SetCloseListener(listener func()) {
	this.closeButton.SetClickListener(listener)
}

func (this *TopToolFrame) SetMiniListener(listener func()) {
	this.miniButton.SetClickListener(listener)
}

func (this *TopToolFrame) pressEvent(event *gui.QMouseEvent) {
	this.isDraging = true
	this.dragPosition = event.GlobalPos()
}

func (this *TopToolFrame) releaseEvent(event *gui.QMouseEvent) {
	this.isDraging = false
}

func (this *TopToolFrame) moveEvent(event *gui.QMouseEvent) {
	if this.isDraging {
		nowDragPosition := event.GlobalPos()
		xMove := nowDragPosition.X() - this.dragPosition.X()
		yMove := nowDragPosition.Y() - this.dragPosition.Y()
		this.dragPosition = nowDragPosition
		this.moveListener(xMove, yMove)
	}
}

func (this *TopToolFrame) SetMoveListener(listener func(x int, y int)) {
	this.moveListener = listener
}
