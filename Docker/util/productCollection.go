package util

// import (
// 	"log"

// 	"github.com/spf13/viper"
// )

// type Product struct {
// 	ID       int    `mapstructure:"id"`
// 	Name     string `mapstructure:"name"`
// 	Quantity int    `mapstructure:"quantity"`
// }

// type Collection struct {
// 	ProductCollection []Product `mapstructure:"product"`
// }

// var PCollection *Collection

// func LoadCollection(path string) (err error) {

// 	//adding app.env variables into struct()

// 	viper.AddConfigPath(path)

// 	viper.SetConfigName("products")

// 	viper.SetConfigType("json")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()

// 	if err != nil {

// 		print("Error detected!", err, "V", err.Error())

// 	}

// 	//save values into struct{} or map{} using umarshal

// 	viper.Unmarshal(&PCollection)

// 	log.Print("OK", PCollection)

// 	return

// }

import (
	"log"

	"github.com/spf13/viper"
)

type Product struct {
	ID       int    `mapstructure:"id"`
	Name     string `mapstructure:"name"`
	Quantity int    `mapstructure:"quantity"`
}

//type Collection []Product `json:product`

type Collection struct {
	Listaa []Product `mapstructure:"productlist"`
}

var PCollection *Collection

func LoadCollection(path string) (err error) {
	print("BEBE", path)
	//adding app.env variables into struct()
	viper.AddConfigPath(path)
	viper.SetConfigName("collection")
	viper.SetConfigType("json")

	print("AJUNGE UNDE TREBUIE", path)

	err = viper.ReadInConfig()

	print("AJUNGE UNDE TREBUIE", err)

	if err != nil {
		log.Print("Error detected!", err, "V", err.Error())
	}
	print("CONTINUEE")
	//save values into struct{} or map{} using umarshal
	err = viper.Unmarshal(&PCollection)
	print("UNMARSHELD", err)
	if err != nil {
		log.Print("EROROR")
	}
	log.Print("OKEY", PCollection)

	return
}
