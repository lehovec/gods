package main

import (
	"fmt"
	"github.com/lehovec/gods"
)


func main() {

	// Client init
	client := gods.New()
	clientStorage := gods.NewMemoryStorage()
	client.SetStorage(clientStorage)
	client.AddClientConnection("client")
	document := "In ancient times cats were worshipped as gods; they have not forgotten this."
	clientStorage.Set(gods.DOCTYPE_DOCUMENT, "client", document)
	fmt.Println("Set client document to: \""+document+"\"")

	// Server init
	server := gods.New()
	serverStorage := gods.NewMemoryStorage()
	server.SetStorage(serverStorage)
	server.AddServerConnection("server")

	// Get client patches
	clientPatches, err := client.GetPatch("client")

	// Apply patch on server
	doc, err := server.ApplyPatch(clientPatches, "server")
	fmt.Println("Server doc: ", doc, err)
	document = "In ancient ages cats were worshipped as gods; they never forgot this."
	serverStorage.Set(gods.DOCTYPE_DOCUMENT, "server", document)
	fmt.Println("Set server document to: \""+document+"\"")

	// Get patches on server
	serverPatches, err := server.GetPatch("server")

	// Apply Patch on client
	document = "In the past cats were worshipped as gods; they have not forgotten this."
	clientStorage.Set(gods.DOCTYPE_DOCUMENT, "client", document)
	fmt.Println("Set client document to: \""+document+"\"")
	doc, err = client.ApplyPatch(serverPatches, "client")
	fmt.Println("Client doc: ", doc, err)

	// Get patch on client
	clientPatches, err = client.GetPatch("client")

	// Apply patch on server
	doc, err = server.ApplyPatch(clientPatches, "server")
	fmt.Println("Server doc: ", doc, err)
}
