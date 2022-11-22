In order to serve go documentation compactly,
through the `godoc` command, I need two terminals, one to 
launch the server and one to open up my browser 
and display it.

This is accomplished by a VScode task in `tasks.json`
that calls `genDoc.bat`, which spawns two workers,
`startServer.bat` and `openDocs.bat`.

`startServer.bat` launches the `godoc` server and keeps 
running in the background, while `openDocs.bat` opens
up the system's default browser displaying the docs, 
it then exits.