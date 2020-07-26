// Package main launches the calculator example directly
package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type draggableBox struct {
	widget.Box
}

var _ fyne.Draggable = (*draggableBox)(nil)
var dragging = false
var draggingCard fyne.CanvasObject
var draggingCardPos = 0

func (db *draggableBox) DragEnd() {
	dragging = false
	draggingCard = nil
}

func (db *draggableBox) Dragged(e *fyne.DragEvent) {
	if dragging == false {
		dragging = true

		for i, child := range db.Children {
			if e.Position.X > child.Position().X && e.Position.X < child.Position().X+child.Size().Width {
				if e.Position.Y > child.Position().Y && e.Position.Y < child.Position().Y+child.Size().Height {
					draggingCard = child
					draggingCardPos = i
					break
				}
			}
		}
		if draggingCard == nil {
			return
		}
	}
	var currentChild fyne.CanvasObject
	var currentChildPos = 0
	for i, child := range db.Children {
		if e.Position.X > child.Position().X && e.Position.X < child.Position().X+child.Size().Width {
			if e.Position.Y > child.Position().Y && e.Position.Y < child.Position().Y+child.Size().Height {
				currentChild = child
				currentChildPos = i
				break
			}
		}
	}
	if currentChild == nil || currentChild == draggingCard {
		return
	}
	if db.Horizontal == true {
		if currentChild.Position().X > draggingCard.Position().X {
			// Moving Right
			if e.Position.X > currentChild.Position().X+currentChild.Size().Width/2 {
				// Swap
				db.Children[draggingCardPos], db.Children[currentChildPos] =
					db.Children[currentChildPos], db.Children[draggingCardPos]
				draggingCardPos = currentChildPos
				db.Refresh()
			}
		} else {
			// Moving Let
			if e.Position.X < currentChild.Position().X+currentChild.Size().Width/2 {
				// Swap
				db.Children[draggingCardPos], db.Children[currentChildPos] =
					db.Children[currentChildPos], db.Children[draggingCardPos]
				draggingCardPos = currentChildPos
				db.Refresh()
			}
		}
	} else {
		if currentChild.Position().Y > draggingCard.Position().Y {
			// Moving Down
			if e.Position.Y > currentChild.Position().Y+currentChild.Size().Height/2 {
				// Swap
				db.Children[draggingCardPos], db.Children[currentChildPos] =
					db.Children[currentChildPos], db.Children[draggingCardPos]
				draggingCardPos = currentChildPos
				db.Refresh()
			}
		} else {
			// Moving Up
			if e.Position.Y < currentChild.Position().Y+currentChild.Size().Height/2 {
				// Swap
				db.Children[draggingCardPos], db.Children[currentChildPos] =
					db.Children[currentChildPos], db.Children[draggingCardPos]
				draggingCardPos = currentChildPos
				db.Refresh()
			}
		}
	}
}

func newDraggableBox() *draggableBox {
	db := &draggableBox{}
	db.ExtendBaseWidget(db)

	return db
}

func main() {
	app := app.New()

	w := app.NewWindow("Fyne Demo")

	db := newDraggableBox()
	for i := 0; i < 10; i++ {
		text := fmt.Sprintf("Card Number: %d", i)
		label := widget.NewLabel(text)
		label.Show()
		db.Append(label)
	}
	db.Show()

	mainHbox := widget.NewHBox(db)

	mainContainer := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil), mainHbox)
	w.SetContent(mainContainer)
	w.Resize(fyne.NewSize(1024, 576))

	w.ShowAndRun()
}
