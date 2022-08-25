package convert

import "strconv"

func Id_Atoi(id string) (*uint, error) {
	castIntId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	castUintId := uint(castIntId)

	return &castUintId, nil
}

func Id_Itoa(id uint) string {
	castId := strconv.Itoa(int(id))

	return castId
}

func Id_Uint(id uint) int {
	return int(id)
}
