install:
	go1.17 build -o yealink-click2dial .
	cp yealink-click2dial ~/bin/yealink-click2dial
	rm yealink-click2dial
