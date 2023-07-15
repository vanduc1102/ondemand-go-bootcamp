package services

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/vanduc1102/ondemand-go-bootcamp/src/models"
)

func TestImportSuccessfully(t *testing.T) {
	rStatus := 200
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(rStatus)
		res.Write([]byte(`
        {
            "count": 1281,
            "next": "https://pokeapi.co/api/v2/pokemon?offset=3&limit=2",
            "previous": "https://pokeapi.co/api/v2/pokemon?offset=0&limit=1",
            "results": [
              {
                "name": "ivysaur",
                "url": "https://pokeapi.co/api/v2/pokemon/2/"
              },
              {
                "name": "venusaur",
                "url": "https://pokeapi.co/api/v2/pokemon/3/"
              }
            ]
          }`))
	}))

	defer func() { testServer.Close() }()

	mockUrl := testServer.URL
	file, err1 := os.CreateTemp("", "ondemand-golang-bootcamp-test-write-*.csv")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer os.Remove(file.Name())
	viper.New()
	viper.Set("POKEMON_API_ENDPOINT", mockUrl+"/")
	viper.Set("CSV_FILE", file.Name())

	_, err := Import(2, 1)

	assert.NoError(t, err)
	content, err := os.ReadFile(viper.GetString("CSV_FILE"))
	assert.NoError(t, err)
	assert.Contains(t, string(content), "3,venusaur")

}

func TestImportCantCallRemote(t *testing.T) {
	rStatus := 400
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(rStatus)
		res.Write([]byte("bad request"))
	}))

	defer func() { testServer.Close() }()

	mockUrl := testServer.URL
	file, err1 := os.CreateTemp("", "ondemand-golang-bootcamp-test-write-*.csv")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer os.Remove(file.Name())
	viper.New()
	viper.Set("POKEMON_API_ENDPOINT", mockUrl+"/")
	viper.Set("CSV_FILE", file.Name())

	_, err := Import(2, 1)

	assert.Error(t, err, "invalid character 'b' looking for beginning of value")

}

func TestFindAll(t *testing.T) {
	file, err1 := os.CreateTemp("", "ondemand-golang-bootcamp-test-write-*.csv")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer os.Remove(file.Name())
	viper.New()
	viper.Set("CSV_FILE", file.Name())

	if _, err := file.Write([]byte("1,poke\n2,mon")); err != nil {
		assert.NoError(t, err, "should write without error")
	}

	pokemonList, err := FindAll()
	assert.NoError(t, err, "should read all record from file")
	assert.ElementsMatch(t, []models.Pokemon{{Id: 1, Name: "poke"}, {Id: 2, Name: "mon"}}, pokemonList, "should contains records")

}

func TestFindAll_Error(t *testing.T) {
	file, err1 := os.CreateTemp("", "ondemand-golang-bootcamp-test-write-*.csv")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer os.Remove(file.Name())
	viper.New()
	viper.Set("CSV_FILE", file.Name())

	if _, err := file.Write([]byte("1,poke\n2s,mon")); err != nil {
		assert.NoError(t, err, "should write without error")
	}

	_, err := FindAll()
	assert.Error(t, err, "id=2s is not a valid integer type")

}

func TestFindById(t *testing.T) {
	file, err1 := os.CreateTemp("", "ondemand-golang-bootcamp-test-write-*.csv")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer os.Remove(file.Name())
	viper.New()
	viper.Set("CSV_FILE", file.Name())

	if _, err := file.Write([]byte("1,poke\n2,mon")); err != nil {
		assert.NoError(t, err, "should write without error")
	}

	pokemon, err := FindById(1)
	assert.NoError(t, err, "should read all record from file")
	assert.Equal(t, models.Pokemon{Id: 1, Name: "poke"}, *pokemon, "should contains a record")

}

func TestFindById_NotFound(t *testing.T) {
	file, err1 := os.CreateTemp("", "ondemand-golang-bootcamp-test-write-*.csv")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer os.Remove(file.Name())
	viper.New()
	viper.Set("CSV_FILE", file.Name())

	if _, err := file.Write([]byte("1,poke\n2,mon")); err != nil {
		assert.NoError(t, err, "should write without error")
	}

	_, err := FindById(100)
	assert.Error(t, err, "pokemon id=100 does not exist")

}
