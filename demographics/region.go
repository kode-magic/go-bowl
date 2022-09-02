package demographics

import (
	"log"

	"github.com/kode-magic/go-bowl/ulids"
)

func AddCountryRegion(region Region) (string, error) {
	db := connectDemograpgicsDB()
	region.Id = ulids.GenerateUUID().String()
	stmt, err := db.Prepare("INSERT INTO regions(id, name, country_id) VALUES (?, ?, ?);")

	if err != nil {
		return "", err
	}

	_, resultErr := stmt.Exec(region.Id, region.Nmae, region.CountryId)
	if resultErr != nil {
		return "", resultErr
	}

	defer stmt.Close()

	return region.Nmae + " successfully added to country regions", nil
}

func FetchCountryRegions(country string) []Region {
	db := connectDemograpgicsDB()
	rows, err := db.Query("SELECT id, name, country_id, created_at, updated_at FROM regions WHERE country_id = '"+ country +"';")

	if err != nil {
		log.Fatal(err)
	}

	regions := make([]Region, 0)

	for rows.Next() {
		region := Region{}
		err = rows.Scan(&region.Id, &region.Nmae, &region.CountryId, &region.CreatedAt, &region.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}

		regions = append(regions, region)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return regions
}