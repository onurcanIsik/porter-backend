package main

import (
	"porter/pkg/database"
)

func main() {

	database, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	//secretKey := os.Getenv("SECRET_KEY")
	//jwtManager := jwt.NewJWTManager(secretKey)

	//r := router.SetupRoutes()

	//log.Println("Server is running on port 8080")
	//if err := r.Listen(":8080"); err != nil {
	//	log.Fatalf("Failed to start server: %v", err)
	//}

}
