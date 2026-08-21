package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PokiRes struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous string  `json:"previous"`
	Results  []Areas `json:"results"`
}
type Areas struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func makereq(cnf *config, isnext bool) error {
	var url string
	if isnext {
		url = cnf.next
		cnf.prev = cnf.next
	} else {
		url = cnf.prev
		cnf.next = cnf.prev
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var pokiRes PokiRes
	if err := json.Unmarshal(data, &pokiRes); err != nil {
		return err
	}
	for _, a := range pokiRes.Results {
		fmt.Println(a.Name)
	}
	if isnext {
		cnf.next = pokiRes.Next
	} else {
		cnf.prev = pokiRes.Previous
	}
	return nil
}

func getmap(cnf *config) error {
	if err := makereq(cnf, true); err != nil {
		return err
	}
	return nil
}

func getmapb(cnf *config) error {
	if err := makereq(cnf, false); err != nil {
		return err
	}
	return nil
}
