package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/PoorMercymain/REST-API-work-duration-counter/internal/handler"
	"github.com/PoorMercymain/REST-API-work-duration-counter/internal/repository"
	"github.com/PoorMercymain/REST-API-work-duration-counter/internal/service"
	"github.com/PoorMercymain/REST-API-work-duration-counter/pkg/router"
	"github.com/PoorMercymain/REST-API-work-duration-counter/pkg/server"

	"github.com/julienschmidt/httprouter"
)

func main() {

	db := repository.NewDb()

	or := repository.NewOrder(db)
	orderS := service.NewOrder(or)
	oh := handler.NewOrder(orderS)

	rr := repository.NewReserve(db)
	rs := service.NewReserve(rr)
	rh := handler.NewReserve(rs)

	sr := repository.NewService(db)
	ss := service.NewService(sr)
	sh := handler.NewService(ss)

	ur := repository.NewUser(db)
	us := service.NewUser(ur)
	uh := handler.NewUser(us)

	r := httprouter.New()

	r.POST("/order", router.WrapHandler(oh.Create))
	r.PUT("/order", router.WrapHandler(oh.Update))
	r.DELETE("/order/:id", router.WrapHandler(oh.Delete))
	r.GET("/order/:id", router.WrapHandler(oh.Read))
	r.POST("/order/addservice", router.WrapHandler(oh.AddService))

	r.POST("/reserve", router.WrapHandler(rh.Create))
	r.PUT("/reserve", router.WrapHandler(rh.Update))
	r.DELETE("/reserve/:id", router.WrapHandler(rh.Delete))
	r.GET("/reserve/:id", router.WrapHandler(rh.Read))
	r.POST("/reserve/approverevenue", router.WrapHandler(rh.ApproveRevenue))

	r.POST("/service", router.WrapHandler(sh.Create))
	r.PUT("/service", router.WrapHandler(sh.Update))
	r.DELETE("/service/:id", router.WrapHandler(sh.Delete))
	r.GET("/service/:id", router.WrapHandler(sh.Read))

	r.POST("/user", router.WrapHandler(uh.Create))
	r.PUT("/user", router.WrapHandler(uh.Update))
	r.DELETE("/user/:id", router.WrapHandler(uh.Delete))
	r.GET("/user/:id", router.WrapHandler(uh.Read))
	r.GET("/user/balance/:id", router.WrapHandler(uh.ReadBalance))
	r.POST("/user/reserve", router.WrapHandler(uh.ReserveMoney))
	r.POST("/user/addmoney", router.WrapHandler(uh.AddMoney))

	theServer := server.New("8000", r)

	var err error

	go func() {
		err = theServer.Run()
	}()

	fmt.Println("Server started")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	if err != nil {
		log.Fatalf("Error occured while running server - %s", err.Error())
	}
}
