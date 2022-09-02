package demographics

import (
	"fmt"
	"log"

	"github.com/kode-magic/go-bowl/ulids"
	_ "github.com/mattn/go-sqlite3"
)

func AddCountry(country Country) (string, error) {
	db := connectDemograpgicsDB()
	country.ID = ulids.GenerateUUID().String()
	stmt, err := db.Prepare("INSERT INTO countries(id, name, continent) VALUES (?, ?, ?);")

	if err != nil {
		return "", err
	}

	result, resultErr := stmt.Exec(country.ID, country.Name, country.Continent)
	if resultErr != nil {
		return "", resultErr
	}

	fmt.Print(result)

	defer stmt.Close()

	return country.Name + " successfully added.", nil
}

func FetchCountries() []Country {
	db := connectDemograpgicsDB()

	rows, err := db.Query("SELECT id, name, continent, created_at, updated_at FROM countries;")

	if err != nil {
		log.Fatal(err)
	}

	countries := make([]Country, 0)

	for rows.Next() {
		country := Country{}
		err = rows.Scan(&country.ID, &country.Name, &country.Continent, &country.CreatedAt, &country.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}

		countries = append(countries, country)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return countries
}
