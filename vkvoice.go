package vkvoice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AudioSpech struct {
	Qid    string `json:"qid"`
	Result struct {
		Texts []struct {
			Text           string  `json:"text"`
			Confidence     float64 `json:"confidence"`
			PunctuatedText string  `json:"punctuated_text"`
		} `json:"texts"`
		PhraseID string `json:"phrase_id"`
	} `json:"result"`
}

func Voice2Text(file string, token string) (string, error) {

	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	req, err := http.NewRequest("POST", "https://voice.mcs.mail.ru/asr", f)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "audio/ogg; codecs=opus")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// responseString := string(responseData)
	// fmt.Println(responseString)
	var ctask AudioSpech

	jsonErr := json.Unmarshal(responseData, &ctask)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return "", jsonErr
	}
	if len(ctask.Result.Texts) ==  0 {
		return "", fmt.Errorf("Fail_speech_voice")
	}
	if ctask.Result.Texts[0].PunctuatedText != "" {
		return ctask.Result.Texts[0].PunctuatedText, nil
	} else {
		return "", err
	}

}
