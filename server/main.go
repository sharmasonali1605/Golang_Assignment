package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/sharmasonali1605/Golang_Assignment/blogpb"
	"github.com/sharmasonali1605/Golang_Assignment/handler"
	"github.com/sharmasonali1605/Golang_Assignment/repository"
	"github.com/sharmasonali1605/Golang_Assignment/service"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8072"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	repo := repository.NewInMemoryBlogRepository()
	svc := service.NewBlogService(repo)
	grpcServer := grpc.NewServer()

	blogpb.RegisterBlogServiceServer(grpcServer, handler.NewBlogHandler(svc))

	log.Printf("gRPC server started on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start Server : %v", err)
	}

	// jobs := make(chan int, 100)
	// var wg sync.WaitGroup

	// for i := 0; i < 3; i++ {
	// 	wg.Add(1)
	// 	go worker(i, jobs, &wg)
	// }

	// for i := 0; i < 10; i++ {
	// 	jobs <- i
	// }

	// close(jobs)
	// wg.Wait()
	// fmt.Println("All workers completed")

	//largestEle([]int{1, 23, 12, 9, 30, 2, 50}, 2)

}

// func worker(workerId int, jobs chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for job := range jobs {
// 		fmt.Printf("Worker %d started job %d\n", workerId, job)
// 	}
// 	//panic("unimplemented")
// }

// func largestEle(arr []int, k int) {
// 	//[] arr = {1,23,12,9,30,2,50}
// 	for i := 0; i < len(arr); i++ {
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[i] < arr[j] {
// 				arr[i], arr[j] = arr[j], arr[i]
// 			}
// 		}
// 	}

// 	for i := 0; i < k; i++ {
// 		fmt.Print(arr[i], " ")
// 	}

// }

/*

Largest element in array {1,23,12,9,30,2,50} k =3 -> {50,30,23}
*/
