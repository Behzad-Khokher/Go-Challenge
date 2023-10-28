package main

import (
	"fmt"
	"log"

	pg "github.com/Behzad-Khokher/Go-Challenge/store/postgres"
	"github.com/Behzad-Khokher/Go-Challenge/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=behzad password=123 dbname=challenge port=5432 sslmode=disable"

	// Initialize GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// // Initialize ComponentStore with db connection
	// componentStore := pg.NewComponentStore(db)

	// endpointUsecase := usecase.NewEndpointUsecase(componentStore)

	// createEndpointReq := usecase.CreateEndpointReq{
	// 	Name:   "Another test endpoint for testing 1",
	// 	URL:    "http://localhost",
	// 	Method: "GET",
	// }
	// errdb := endpointUsecase.Create(createEndpointReq)
	// if errdb != nil {
	// 	log.Fatalf("Failed to create endpoint: %v", errdb)
	// }
	// Branch
	branchStore := pg.NewBranchStore(db)
	branchUsecase := usecase.NewBranchUsecase(branchStore)
	if err := branchUsecase.Create(usecase.CreateBranchReq{Name: "master"}); err != nil {
		log.Fatalf("Failed to create branch: %v", err)
	}
	fmt.Println("Branch 'master' created!")

	// Commits
	commitStore := pg.NewCommitStore(db)
	commitUsecase := usecase.NewCommitUsecase(commitStore)
	commitNames := []string{"Initial commit", "Add README", "Fix bug"}
	for _, name := range commitNames {
		if err := commitUsecase.Create(usecase.CreateCommitReq{Name: name, BranchID: 1}); err != nil {
			log.Fatalf("Failed to create commit '%s': %v", name, err)
		}
	}
	fmt.Println("Commits created!")

	// Components
	componentStore := pg.NewComponentStore(db) // Assuming you have this in place
	componentUsecase := usecase.NewEndpointUsecase(componentStore)
	components := []usecase.CreateEndpointReq{
		{Name: "Login", URL: "/login", Method: "POST"},
		{Name: "Logout", URL: "/logout", Method: "POST"},
		{Name: "Profile", URL: "/profile", Method: "GET"},
	}
	for _, component := range components {
		if err := componentUsecase.Create(component); err != nil {
			log.Fatalf("Failed to create component '%s': %v", component.Name, err)
		}
	}
	fmt.Println("Components created!")

	// Add components to a commit (to the first commit for simplicity)
	commitComponentStore := pg.NewCommitComponentStore(db)
	commitComponentUsecase := usecase.NewCommitComponentUsecase(commitComponentStore)
	if err := commitComponentUsecase.AddComponentsToCommit(usecase.AddComponentsToCommitReq{
		CommitID:     1,
		ComponentIDs: []int{1, 2, 3},
	}); err != nil {
		log.Fatalf("Failed to add components to commit: %v", err)
	}
	fmt.Println("Components added to commit!")

	latestCommit, err := commitUsecase.GetByBranchID(1) // Assuming 1 is the ID for master branch
	if err != nil || len(latestCommit) == 0 {
		log.Fatalf("Failed to retrieve the latest commit for 'master' branch: %v", err)
	}
	latestCommitID := latestCommit[len(latestCommit)-1].ID

	// Get components associated with the latest commit
	commitComponents, err := commitComponentUsecase.GetCommitsByComponentID(latestCommitID)
	if err != nil {
		log.Fatalf("Failed to retrieve components for the latest commit: %v", err)
	}

	// Display components
	fmt.Println("Components in the latest commit of 'master' branch:")
	for _, cc := range commitComponents {
		component, err := componentUsecase.GetByID(cc.ComponentID) // Assuming you have GetByID method in the component usecase
		if err != nil {
			log.Printf("Failed to retrieve component with ID %d: %v", cc.ComponentID, err)
			continue
		}
		fmt.Printf("Name: %s, URL: %s, Method: %s\n", component.Name, component.URL, component.Method)
	}

	// // Testing: Create a new Component
	// newComponent := &domain.Component{
	// 	Type: domain.EndpointComponent,
	// 	Name: "TestEndpoint2",
	// 	Data: []byte(`{"ID":2, "Name":"TestEndpoint2", "URL":"http://localhost", "Method":"GET"}`),
	// }

	// // Create
	// err = componentStore.Create(newComponent)
	// if err != nil {
	// 	log.Fatalf("Failed to create component: %v", err)
	// }

	// fmt.Println("Successfully created a new component!")

	// // Retrieve
	// component, err := componentStore.GetByID(newComponent.ID)
	// if err != nil {
	// 	log.Fatalf("Failed to retrieve component: %v", err)
	// }
	// fmt.Printf("Retrieved component: %+v\n", component)

}
