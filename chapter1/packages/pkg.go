package packages

import (
	"fmt"
	"github.com/google/uuid"
)

func ShowPackage() {

	newUUID, err := uuid.NewUUID()
	fmt.Printf("Calling Internal Package.\n")
	fmt.Printf("UUID generated: %+v, Error: %+v.\n", newUUID, err)
}
