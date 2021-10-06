all: bindata.go

bindata.go: static_web/static static_web/static
	go-bindata -pkg='bindata' -o bindata/bindata.go static/

clean:
	rm bindata/bindata.go
