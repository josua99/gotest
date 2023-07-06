package main

import (
	"log"

	"github.com/josua99/gotest/internal/api"
)

func main() {
/*
	client := http.Client{}

	resp, err := client.Post("https://webhook.servicesforfree.com/993f714f-eebe-465c-b123-96bc7cc50753", "application/json", nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))

	os.Exit(0)
	*/



	apiServer := api.NewServer(nil)

	//ctx, _ := context.WithTimeout(context.Background(), (time.Second*3))
	//go func ()  {
		if err := apiServer.Start(); err != nil {
			log.Fatalln(err)
		}
	//}()
	
	//<-ctx.Done()
}
