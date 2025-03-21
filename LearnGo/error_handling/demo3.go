package errorhandling

import "fmt"

type BorderException struct {
	message   string
	parameter int
}

func (b BorderException) Error() string {
	return fmt.Sprintf("%d %s", b.parameter, b.message)

}

func TahminEt2(tahmin int) (string, error) {
	if tahmin > 100 || tahmin < 1 {
		return "", &BorderException{"Sınırların dışında kaldı", tahmin}
	}
	return "bildiniz", nil
}
