package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	accessToken := "TokenAnda"
	owner := "NamaOwner"
	repo := "LokasiRepo"
	file := "Nama File"

	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/main/%s", owner, repo, file)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Gagal melakukan request:", err)
		return
	}

	req.Header.Add("Authorization", "token "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Gagal melakukan request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Gagal mengambil %s: %s\n", file, resp.Status)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Gagal membaca respon body:", err)
		return
	}

	err = ioutil.WriteFile(file, body, 0644)
	if err != nil {
		fmt.Println("Gagal menulis file:", err)
		return
	}

	notif := fmt.Sprintf("%s sukses mendownload", file)
	fmt.Println(notif)
}
