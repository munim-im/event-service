# Event Service

## Steps taken
* Initialize go modules using `go mod init event-service`.
* Initialize git at the very first so that we can use the VCS.
* Install `gin` package for http web framework.
* Install `gorm` for ORM and install gorm postgres driver as we will be using postgresql as database.
```shell
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres 
```
* Start with the data model designing as that is the core part of service. I like to put all the models in a 
single folder if number of models are few in number. If we require more models, then I try to keep them in folders so that
I can see all the models with their related and supported models too.  