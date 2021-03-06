package views

import (
	. "github.com/fishedee/web"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type MusicListTable struct {
	*widgets.QTableView
	Model
	parent   widgets.QWidget_ITF
	model    *gui.QStandardItemModel
	curIndex int
}

type MusicListContextAction struct {
	Name   string
	Action func(actionRows []int)
}

func NewMusicListTable(parent widgets.QWidget_ITF) *MusicListTable {
	musicListTable := MusicListTable{}
	InitModel(&musicListTable)
	musicListTable.init(parent)
	return &musicListTable
}

func (this *MusicListTable) init(parent widgets.QWidget_ITF) {
	this.parent = parent
	this.QTableView = widgets.NewQTableView(parent)
	this.model = gui.NewQStandardItemModel(parent)
	this.model.SetColumnCount(3)
	this.SetFixedWidth(300)
	this.SetShowGrid(false)
	this.SetWordWrap(false)
	this.SetMouseTracking(true)

	this.VerticalHeader().SetSectionResizeMode(widgets.QHeaderView__ResizeToContents)
	this.VerticalHeader().SetSectionsClickable(false)

	this.HorizontalHeader().SetSectionResizeMode(widgets.QHeaderView__Stretch)
	this.HorizontalHeader().SetSectionsClickable(false)
	this.HorizontalHeader().Hide()
	this.HorizontalHeader().SetStyleSheet(`
        selection-background-color:lightblue;
    `)

	this.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	this.SetSelectionMode(widgets.QAbstractItemView__ExtendedSelection)

	this.SetFocusPolicy(core.Qt__NoFocus)
	this.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	this.SetStyleSheet(`
        selection-background-color:lightblue;
    `)

	this.SetModel(this.model)
	this.curIndex = -1
}

func (this *MusicListTable) AddSong(title string, artist string, timeString string) {
	song := []*gui.QStandardItem{
		gui.NewQStandardItem2(title),
		gui.NewQStandardItem2(artist),
		gui.NewQStandardItem2(timeString),
	}
	song[2].SetTextAlignment(core.Qt__AlignCenter)
	this.model.AppendRow(song)
}

func (this *MusicListTable) DelSong(index int) {
	this.model.RemoveRow(index, core.NewQModelIndex())
	if this.curIndex > index {
		this.curIndex--
	}
}

func (this *MusicListTable) ClearSong() {
	this.model.Clear()
}

func (this *MusicListTable) SetDoubleClickListener(handler func(index int)) {
	this.ConnectMouseDoubleClickEvent(func(event *gui.QMouseEvent) {
		if event.Button() == core.Qt__LeftButton {
			pos := event.Pos()
			index := this.IndexAt(pos).Row()
			handler(index)
		}
	})
}

func (this *MusicListTable) ActiveIndex(index int) {
	if this.curIndex != -1 {
		color := gui.NewQColor3(0, 0, 0, 255)
		brush := gui.NewQBrush3(color, 0)
		this.model.Item(this.curIndex, 0).SetForeground(brush)
		this.model.Item(this.curIndex, 1).SetForeground(brush)
		this.model.Item(this.curIndex, 2).SetForeground(brush)
	}
	if index != -1 {
		color2 := gui.NewQColor3(255, 0, 0, 255)
		brush2 := gui.NewQBrush3(color2, 0)
		this.model.Item(index, 0).SetForeground(brush2)
		this.model.Item(index, 1).SetForeground(brush2)
		this.model.Item(index, 2).SetForeground(brush2)
	}
	this.curIndex = index
}

func (this *MusicListTable) getContextMenu(actions []MusicListContextAction) *widgets.QMenu {
	contextMenu := widgets.NewQMenu(this.parent)
	for _, singleAction := range actions {
		singleActionAction := singleAction.Action
		if singleAction.Name != "" {
			action := widgets.NewQAction2(singleAction.Name, this.parent)
			action.ConnectTriggered(func(checked bool) {
				selectedIndexs := this.SelectionModel().SelectedRows(0)
				selectedRows := []int{}
				for _, singleIndex := range selectedIndexs {
					row := this.model.ItemFromIndex(singleIndex).Row()
					selectedRows = append(selectedRows, row)
				}
				singleActionAction(selectedRows)
			})
			contextMenu.QWidget.AddAction(action)
		} else {
			contextMenu.AddSeparator()
		}

	}
	return contextMenu
}

func (this *MusicListTable) SetContextMenuListener(handler func(index int) []MusicListContextAction) {
	this.ConnectContextMenuEvent(func(event *gui.QContextMenuEvent) {
		pos := event.Pos()
		index := this.IndexAt(pos).Row()
		actions := handler(index)
		contextMenu := this.getContextMenu(actions)
		contextMenu.Exec2(event.GlobalPos(), nil)
	})
}
