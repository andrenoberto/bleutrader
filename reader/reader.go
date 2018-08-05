package reader

import (
	"encoding/hex"
	"log"
	"github.com/st3v/translator/google"
	"Bleu/environment"
	"fmt"
	"Bleu/packages"
	"net/http"
	"encoding/json"
)

type EtherScanResponse struct {
	Jsonrpc float64   `json:"jsonrpc,string"`
	Id      int64     `json:"id,string"`
	Result  EtherScan `json:"result"`
}

type EtherScan struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	To               string `json:"to"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

func GetMessageByHash(hash string) {
	uri := environment.EtherScanRootUrl + "?module=proxy&action=eth_getTransactionByHash"
	uri += "&txhash=" + hash
	uri += "&apikey=" + environment.EtherScanApiKey
	req, err := http.NewRequest("GET", uri, nil)
	packages.ErrorHandler(err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	packages.ErrorHandler(err)
	defer resp.Body.Close()
	var responseJson EtherScanResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&responseJson)
	translateHex(responseJson.Result.Input[2:])
}

func translateHex(message string) {
	fmt.Println("########################################################################")
	fmt.Println("# Translating...")
	fmt.Println("########################################################################")
	// Decoding message
	decoded, err := hex.DecodeString(message)
	if err != nil {
		log.Fatal(err)
	}
	decodedText := string(decoded)
	// Translate
	paragraphs := make([]string, 0)
	translatedParagraphs := make([]string, 0)
	init := 0
	// Split paragraphs
	for index := range decodedText {
		if decodedText[index] == 10 {
			paragraphs = append(paragraphs, decodedText[init:index])
			init = index
		}
	}
	// Add last paragraph
	paragraphs = append(paragraphs, decodedText[init:])
	translator := google.NewTranslator(environment.GoogleTranslateApiKey)
	languageCode, err := translator.Detect(paragraphs[0])
	if err != nil {
		log.Panicf("Error detecting language: %s", err.Error())
	}
	// Translating
	for index := range paragraphs {
		translation, _ := translator.Translate(paragraphs[index], languageCode, "pt")
		translatedParagraphs = append(translatedParagraphs, translation)
	}
	// Print result
	packages.ClearScreen()
	for index := range translatedParagraphs {
		fmt.Printf("%s\n", translatedParagraphs[index])
	}
}
