package main

import (
	"fmt"
	"grpctest/pb"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewCalcServiceClient(conn)

	//gin is used to create API endpoints to call the grpc function
	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid First Number"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Second Number"})
			return
		}

		req := &pb.ClientRequest{FirstNum: int64(a), SecondNum: int64(b)}

		if response, err := client.Add(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		}
	})

	g.GET("/mult/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid First Number"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Second Number"})
			return
		}

		req := &pb.ClientRequest{FirstNum: int64(a), SecondNum: int64(b)}

		if response, err := client.Multiply(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		}

	})

	if err := g.Run(":4040"); err != nil {
		log.Fatal("Failed to run server : %v", err)
	}

}
