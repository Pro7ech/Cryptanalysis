package linear

import(
	"fmt"
)


func P_t(x [8]uint8) (w [8]uint8){

	z0, z1, z2, z3, z4, z5, z6, z7 := x[0], x[1], x[2], x[3], x[4], x[5], x[6], x[7]

	w[0] = z1 ^ z2 ^ z4 ^ z5 ^ z6
    w[1] = z1 ^ z2 ^ z3 ^ z6 ^ z7
    w[2] = z0 ^ z3 ^ z4 ^ z6 ^ z7
    w[3] = z0 ^ z1 ^ z3 ^ z4 ^ z5
    w[4] = z0 ^ z1 ^ z2 ^ z5 ^ z6
    w[5] = z2 ^ z3 ^ z5 ^ z6 ^ z7
    w[6] = z0 ^ z2 ^ z3 ^ z4 ^ z7
    w[7] = z0 ^ z1 ^ z4 ^ z5 ^ z7

    return
}

type mask struct{
	value [8]uint8
	lp [8] float64
	total_lp float64
}


type path_finder struct {
	mask [][]uint8

	sbox_threshold float64
	linear_threshold float64

	rounds int

	linear_diffusion func([8]uint8)([8]uint8)

	sbox [][256]uint8

	blockbyte int
}


func (self *path_finder) find_masks() {


	for j := 0 ; j < self.blockbyte ; j++ {
		fmt.Printf("Computing mask for byte %d \n", self.blockbyte - j - 1)

		for i := 0 ; i < 256 ; i++ {
			masks := new(mask)
			masks.value[j] = uint8(i)
			masks.lp = 1.0

			self.compute_LP_SPN([][]*mask{[]*mask{masks}}, self.rounds)
		}
	}
}


func (self *path_finder) compute_LP_SPN(masks [][]*mask, level int){
	level -= 1
	if level >= 0 {
		current_mask := masks[-len(masks)-1][0].value
		//lp := masks[-len(masks)-1][0].lp

		//new_masks := self.return_best_masks(self.linear_diffusion(current_mask))
	}
}

func (self *path_finder) return_best_masks(current_mask [8]uint8) (masks []*mask){

	var T [8][]*linear_probability

	counter := 0

	for i := range current_mask {
		if current_mask[i] != 0x00 {
			x := targeted_LP_OUT(current_mask[i], self.sbox[i], self.sbox_threshold)

			if len(x) != 0 {
				T[i] = x
				counter += 1
			}else{
				T[i] = new(linear_probability)
				T[i].lpFlo = 1
			}
		}
	}

	if counter > 0 {

		mask = make([]*mask, 8)

		for a := range T[0] {
			for b := range T[1] {
				for c := range T[2] {
					for d := range T[3] {
						for e := range T[4] {
							for f := range T[5] {
								for g := range T[6] {
									for h := range T[7] {

										mask.value[0] = a
										mask.value[1] = b
										mask.value[2] = c
										mask.value[3] = d
										mask.value[4] = e
										mask.value[5] = f
										mask.value[6] = g
										mask.value[7] = h
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return 
}