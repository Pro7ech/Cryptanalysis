package linear

import(
	"github.com/cryptanalysis/utils"
	"sort"
)


type linear_probability struct {
	in uint64
	out uint64
	lpInt int64
	lpFlo float64
}

func compute_linear_probability(sbox []uint64) (result []*linear_probability){

	var tmp int64

	sbox_size := uint64(len(sbox))

	result = make([]*linear_probability, (sbox_size-1)*(sbox_size-1))


	index := 0

	for i := uint64(1) ; i < sbox_size ; i++ {

		for j := uint64(1) ; j < sbox_size ; j++ {

			tmp = - int64((sbox_size>>1))

			for k := uint64(0) ; k < sbox_size; k++ {
				tmp += int64(utils.Parity((i & k) ^ (sbox[k] & j)))
			}

			result[index] = new(linear_probability)
			result[index].in = i
			result[index].out = j
			result[index].lpInt = tmp

			index++
		}
	}

	process_result(result, sbox_size)

	return 
}

func process_result(result []*linear_probability, sbox_size uint64){

	sort.Slice(result, func(i, j int) bool {
		return result[i].lpInt > result[j].lpInt
	})

	var tmp float64

	for i := range result {
		tmp = float64(2*result[i].lpInt) / float64(sbox_size)
		result[i].lpFlo = tmp * tmp
	}
}