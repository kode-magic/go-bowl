package demographics

import "github.com/kode-magic/go-bowl/ulids"

func AddDistrict(district District) (string, error) {
	db := connectDemograpgicsDB()
	district.ID = ulids.GenerateUUID().String()
	stmt, err := db.Prepare("INSERT INTO regions(id, name, region_id) VALUES (?, ?, ?);")

	if err != nil {
		return "", err
	}

	_, resultErr := stmt.Exec(district.ID, district.Name, district.RegionId)
	if resultErr != nil {
		return "", resultErr
	}

	defer stmt.Close()

	return district.Name + " successfully added", nil
}
