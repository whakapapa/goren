package main

import "github.com/whakapapa/gtkgo"


func (p *tOptions) diagPrefs() {
	// initialize GTK
	gtk.Init(nil)

	// construct builder
	bPref, err := gtk.BuilderNew()
	if err != nil {
		log.Fatal("Construction builder failed:", err)
	}

	// load GTK design
	err = b.AddFromFile(GoRenGlade)
	if err != nil {
		log.Fatal("Loading GTK design failed:", err)
	}

	// construct prefs window_main
	obj, err := b.GetObject("DiagPrefs")
	if err != nil {
		log.Fatal("Construction main object failed:", err)
	}

	// connect objects and signals
	// create pointer to main window from object and destructor
	app := obj.(*gtk.Window)
	app.Connect("destroy", func() {
		gtk.MainQuit()
	})

	
}
