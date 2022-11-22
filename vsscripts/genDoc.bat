cd vsscripts
start cmd "ServerHandler" /MIN /c startServer.bat
timeout /t 2 /nobreak
start cmd "StartBrowser" /MIN /c openDocs.bat %1