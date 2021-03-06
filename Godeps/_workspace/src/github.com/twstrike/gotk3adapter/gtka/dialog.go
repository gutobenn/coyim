package gtka

import (
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/coyim/Godeps/_workspace/src/github.com/twstrike/gotk3adapter/gtki"
)

type dialog struct {
	*window
	internal *gtk.Dialog
}

func wrapDialogSimple(v *gtk.Dialog) *dialog {
	if v == nil {
		return nil
	}
	return &dialog{wrapWindowSimple(&v.Window), v}
}

func wrapDialog(v *gtk.Dialog, e error) (*dialog, error) {
	return wrapDialogSimple(v), e
}

func unwrapDialog(v gtki.Dialog) *gtk.Dialog {
	if v == nil {
		return nil
	}
	return v.(*dialog).internal
}

func (v *dialog) Run() int {
	return v.internal.Run()
}

func (v *dialog) SetDefaultResponse(v1 gtki.ResponseType) {
	v.internal.SetDefaultResponse(gtk.ResponseType(v1))
}
