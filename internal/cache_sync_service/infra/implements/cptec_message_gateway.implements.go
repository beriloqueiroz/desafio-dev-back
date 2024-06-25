package implements

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type CptecMessageGateway struct{}

func NewCptecMessageGateway() *CptecMessageGateway {
	return &CptecMessageGateway{}
}

type Cidades struct {
	XMLName xml.Name `xml:"cidades"`
	Text    string   `xml:",chardata"`
	Cidade  []struct {
		Text string `xml:",chardata"`
		Nome string `xml:"nome"`
		Uf   string `xml:"uf"`
		ID   int64  `xml:"id"`
	} `xml:"cidade"`
}

type PrevisaoCidade struct {
	XMLName     xml.Name `xml:"cidade"`
	Text        string   `xml:",chardata"`
	Nome        string   `xml:"nome"`
	Uf          string   `xml:"uf"`
	Atualizacao string   `xml:"atualizacao"`
	Previsao    []struct {
		Text   string `xml:",chardata"`
		Dia    string `xml:"dia"`
		Tempo  string `xml:"tempo"`
		Maxima string `xml:"maxima"`
		Minima string `xml:"minima"`
		Iuv    string `xml:"iuv"`
	} `xml:"previsao"`
}

func (c *CptecMessageGateway) MessageByLocation(ctx context.Context, city string, state string) (string, error) {
	cityId, err := GetCityID(city, state)
	if err != nil {
		return "", err
	}
	climate, err := GetCityClimate(cityId)
	if err != nil {
		return "", err
	}
	return climate, nil
}

func GetCityClimate(id int64) (string, error) {
	resp, err := http.DefaultClient.Get(fmt.Sprintf("http://servicos.cptec.inpe.br/XML/cidade/%d/previsao.xml", id))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var cidades Cidades
	if err := xml.NewDecoder(resp.Body).Decode(&cidades); err != nil {
		return "", err
	}
	jsonStr, err := json.Marshal(cidades)
	if err != nil {
		return "", err
	}
	return string(jsonStr), nil
}

func GetCityID(city string, state string) (int64, error) {
	resp, err := http.DefaultClient.Get("http://servicos.cptec.inpe.br/XML/listaCidades?city=" + city)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	var cidades Cidades
	if err := xml.NewDecoder(resp.Body).Decode(&cidades); err != nil {
		return 0, err
	}
	var cityId int64 = 0
	for _, cidade := range cidades.Cidade {
		if strings.ToLower(cidade.Uf) == strings.ToLower(state) && strings.ToLower(cidade.Nome) == strings.ToLower(city) {
			cityId = cidade.ID
			break
		}
	}
	if cityId == 0 {
		return 0, errors.New("city not found")
	}
	return cityId, nil
}
