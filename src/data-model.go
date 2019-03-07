package main

import (
	"os"
)

import "github.com/whakapapa/gtkgo"

// global data structure
// for GoRenamer


// parameters and signals for UI interaction
type pMainApp struct {
	toggleOptions		*ToggleButton	// bAppOptions
	buttonUndo			*Button			// bAppUndo
	buttonRedo			*ButtonNew		// bAppRedo
	buttonRefresh		*ButtonNew		// bAppRefresh
	menuSelAll			*MenuShell		// mAppSelAll
	menuSelNone			*MenuShell		// mAppSelNone
	menuLoadNames		*MenuShell		// mAppLoadNames
	menuCopy				*MenuShell		// mAppCopy
	menuCut				*MenuShell		// mAppCut
	menuPaste			*MenuShell		// mAppPaste
	menuPrefs			*MenuShell		// mAppPreferences
	menuAbout			*MenuShell		// mAppAbout
	buttonPreview		*Button			// bAppPreview
	buttonClear			*Button			// bAppClear
	buttonRename		*Button			// bAppRename
	comboDirFile		*ComboBoxText	// cbDirFile
	entrySelPat			*Entry			// eSelPattern
	checkRecurse		*CheckButton	// cAppRecursive
	checkExten			*CheckButton	// cAppKeepExten
	checkPreview		*CheckButton	// cAppPreview
}

type tSecPattern struct {
	buttonSaveOrig		*Button			// bPatternSaveOrig
	buttonSaveDest		*Button			// bPatternSaveDest
	buttonEditOrig		*Button			// bPatternEditOrig
	buttonEditDest		*Button			// bPatternEditDest
	comboOrig			*CombBox			// cbPatternOrig
	comboDest			*CombBox			// cbPatternDest
}

type tSecSubstitution struct {
	checkSpaces			*CheckButton	// cSubSpaces
	checkReplace		*CheckButton	// cSubReplace
	checkCapital		*CheckButton	// cSubCapital
	checkAccents		*CheckButton	// cSubAccents
	comboSpaces			*ComboBoxText	// cbSubSpaces
	comboCapital		*ComboBoxText	// cbSubCapital
	entryOrig			*Entry			// eSubOrig
	entryDest			*Entry			// eSubDest
	checkFixSym			*CheckButton	// cSubFixSymbols
}

type tSecInsDel struct {
	checkInsert			*CheckButton	// cInsert
	checkDelete			*CheckButton	// cDelete
	checkDelEnd			*CheckButton	// cDelEnd
	entryInsert			*Entry			// eInsertText
	spinInsPos			*SpinButton		// sInsertPos
	spinDelFrom			*SpinButton		// sDeleteTo
	spinDelTo			*SpinButton		// sDeleteFrom
}

type tSecManual struct {
	entryManual			*Entry			// eManual
}

type tSecImage struct {
	comboOrig			*ComboBoxText	// cbImageOrig
	comboDest			*ComboBoxText	// cbImageDest
	buttonSaveOrig		*Button			// bImageSaveOrig
	buttonSaveDest		*Button			// bImageSaveDest
	buttonEditOrig		*Button			// bImageEditOrig
	buttonEditDest		*Button			// bImageEditDest
}

type tAddPattern struct {
	buttonOK				*Button			// bAddPatternOK
	buttonCancel		*Button			// bAddPatternCancel
	ePattern				*Entry			// eAddPattern
}

type tPatterns struct {
	buttonClose			*Button			// bPatternClose
	buttonAdd			*Button			// bPatternAdd
	buttonRemove		*Button			// bPatterRemove
	buttonEdit			*Button			// bPatternEdit
	buttonUp				*Button			// bPatternUp
	buttonDown			*Button			// bPatternDown
}


// program options
type tDirFile int

const (
	DIRS = 0 + iota
	FILES
	BOTH
)

type tOptions struct {
	// list options
	visible	bool			// are options visible?
	show		tDirFile		// what is shown (DirFile)
	extens	bool			// keep extensions
	preview	bool			// auto preview

	// directories
	dirRoot		string	// root directory
	dirAct		string	// active directory
}

func (o *tOptions) {
	// code to load last settings from gconf
}
