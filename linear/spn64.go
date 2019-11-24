package linear

import (
	"crypto/rand"
	"github.com/cryptanalysis/utils"
	"math"
	"sort"
)

var S8_0 = [256]uint8{142, 177, 44, 179, 140, 17, 46, 19, 240, 79, 114, 77, 242, 207, 112, 205, 190, 129, 20, 139, 180, 33, 30, 43, 64, 255, 202, 245, 74, 127, 192, 117, 223, 96, 85, 106, 213, 224, 95, 234, 1, 62, 171, 52, 11, 158, 161, 148, 111, 208, 237, 210, 109, 80, 239, 82, 49, 14, 147, 12, 51, 174, 145, 172, 31, 32, 181, 42, 21, 128, 191, 138, 94, 225, 212, 235, 84, 97, 222, 107, 238, 81, 108, 83, 236, 209, 110, 211, 47, 16, 141, 18, 45, 176, 143, 178, 113, 206, 243, 204, 115, 78, 241, 76, 144, 175, 50, 173, 146, 15, 48, 13, 160, 159, 10, 149, 170, 63, 0, 53, 193, 126, 75, 116, 203, 254, 65, 244, 136, 183, 34, 189, 130, 23, 40, 29, 105, 214, 227, 220, 99, 86, 233, 92, 219, 100, 89, 102, 217, 228, 91, 230, 186, 133, 24, 135, 184, 37, 26, 39, 68, 251, 198, 249, 70, 123, 196, 121, 5, 58, 167, 56, 7, 154, 165, 152, 55, 8, 157, 2, 61, 168, 151, 162, 246, 73, 124, 67, 252, 201, 118, 195, 27, 36, 185, 38, 25, 132, 187, 134, 197, 122, 71, 120, 199, 250, 69, 248, 41, 22, 131, 28, 35, 182, 137, 188, 119, 200, 253, 194, 125, 72, 247, 66, 232, 87, 98, 93, 226, 215, 104, 221, 150, 169, 60, 163, 156, 9, 54, 3, 90, 229, 216, 231, 88, 101, 218, 103, 164, 155, 6, 153, 166, 59, 4, 57}
var S8_1 = [256]uint8{95, 224, 85, 106, 223, 234, 213, 96, 190, 129, 180, 43, 30, 139, 20, 33, 140, 179, 142, 17, 44, 177, 46, 19, 237, 82, 239, 208, 109, 80, 111, 210, 51, 12, 49, 174, 147, 14, 145, 172, 114, 205, 112, 79, 242, 207, 240, 77, 192, 127, 202, 245, 64, 117, 74, 255, 1, 62, 11, 148, 161, 52, 171, 158, 146, 173, 144, 15, 50, 175, 48, 13, 108, 211, 110, 81, 236, 209, 238, 83, 160, 159, 170, 53, 0, 149, 10, 63, 222, 97, 212, 235, 94, 107, 84, 225, 65, 254, 75, 116, 193, 244, 203, 126, 31, 32, 21, 138, 191, 42, 181, 128, 243, 76, 241, 206, 115, 78, 113, 204, 45, 18, 47, 176, 141, 16, 143, 178, 7, 56, 5, 154, 167, 58, 165, 152, 89, 230, 91, 100, 217, 228, 219, 102, 55, 8, 61, 162, 151, 2, 157, 168, 233, 86, 227, 220, 105, 92, 99, 214, 118, 201, 124, 67, 246, 195, 252, 73, 136, 183, 130, 29, 40, 189, 34, 23, 198, 121, 196, 251, 70, 123, 68, 249, 184, 135, 186, 37, 24, 133, 26, 39, 104, 215, 98, 93, 232, 221, 226, 87, 41, 22, 35, 188, 137, 28, 131, 182, 25, 38, 27, 132, 185, 36, 187, 134, 216, 103, 218, 229, 88, 101, 90, 231, 166, 153, 164, 59, 6, 155, 4, 57, 71, 248, 69, 122, 199, 250, 197, 120, 247, 72, 253, 194, 119, 66, 125, 200, 150, 169, 156, 3, 54, 163, 60, 9}
var S8_2 = [256]uint8{41, 182, 35, 28, 137, 188, 131, 22, 232, 215, 226, 93, 104, 221, 98, 87, 90, 101, 88, 231, 218, 103, 216, 229, 27, 132, 25, 38, 187, 134, 185, 36, 197, 250, 199, 120, 69, 248, 71, 122, 164, 59, 166, 153, 4, 57, 6, 155, 150, 9, 156, 163, 54, 3, 60, 169, 119, 72, 125, 194, 247, 66, 253, 200, 219, 228, 217, 102, 91, 230, 89, 100, 5, 154, 7, 56, 165, 152, 167, 58, 105, 86, 99, 220, 233, 92, 227, 214, 55, 168, 61, 2, 151, 162, 157, 8, 136, 23, 130, 189, 40, 29, 34, 183, 246, 201, 252, 67, 118, 195, 124, 73, 186, 37, 184, 135, 26, 39, 24, 133, 68, 123, 70, 249, 196, 121, 198, 251, 238, 209, 236, 83, 110, 211, 108, 81, 144, 15, 146, 173, 48, 13, 50, 175, 94, 97, 84, 235, 222, 107, 212, 225, 160, 63, 170, 149, 0, 53, 10, 159, 31, 128, 21, 42, 191, 138, 181, 32, 193, 254, 203, 116, 65, 244, 75, 126, 47, 176, 45, 18, 143, 178, 141, 16, 113, 78, 115, 204, 241, 76, 243, 206, 190, 33, 180, 139, 30, 43, 20, 129, 223, 224, 213, 106, 95, 234, 85, 96, 111, 80, 109, 210, 239, 82, 237, 208, 142, 17, 140, 179, 46, 19, 44, 177, 240, 207, 242, 77, 112, 205, 114, 79, 49, 174, 51, 12, 145, 172, 147, 14, 1, 158, 11, 52, 161, 148, 171, 62, 64, 127, 74, 245, 192, 117, 202, 255}
var S8_3 = [256]uint8{210, 111, 80, 109, 82, 237, 208, 239, 140, 177, 46, 179, 44, 19, 142, 17, 28, 41, 182, 35, 188, 131, 22, 137, 226, 87, 104, 93, 98, 221, 232, 215, 163, 150, 9, 156, 3, 60, 169, 54, 125, 200, 247, 194, 253, 66, 119, 72, 77, 240, 207, 242, 205, 114, 79, 112, 51, 14, 145, 12, 147, 172, 49, 174, 99, 214, 233, 220, 227, 92, 105, 86, 2, 55, 168, 61, 162, 157, 8, 151, 83, 238, 209, 236, 211, 108, 81, 110, 146, 175, 48, 173, 50, 13, 144, 15, 45, 16, 143, 18, 141, 178, 47, 176, 204, 113, 78, 115, 76, 243, 206, 241, 189, 136, 23, 130, 29, 34, 183, 40, 252, 73, 118, 67, 124, 195, 246, 201, 222, 107, 84, 97, 94, 225, 212, 235, 63, 10, 149, 0, 159, 160, 53, 170, 228, 89, 102, 91, 100, 219, 230, 217, 165, 152, 7, 154, 5, 58, 167, 56, 26, 39, 184, 37, 186, 133, 24, 135, 123, 198, 249, 196, 251, 68, 121, 70, 128, 181, 42, 191, 32, 31, 138, 21, 65, 244, 203, 254, 193, 126, 75, 116, 101, 216, 231, 218, 229, 90, 103, 88, 187, 134, 25, 132, 27, 36, 185, 38, 33, 20, 139, 30, 129, 190, 43, 180, 95, 234, 213, 224, 223, 96, 85, 106, 158, 171, 52, 161, 62, 1, 148, 11, 192, 117, 74, 127, 64, 255, 202, 245, 250, 71, 120, 69, 122, 197, 248, 199, 4, 57, 166, 59, 164, 155, 6, 153}
var S8_4 = [256]uint8{63, 214, 22, 26, 29, 20, 226, 208, 60, 0, 176, 231, 144, 167, 182, 31, 96, 225, 210, 198, 107, 186, 37, 190, 241, 86, 206, 45, 148, 12, 120, 156, 109, 48, 13, 17, 23, 128, 113, 30, 124, 172, 139, 170, 228, 229, 181, 234, 131, 154, 3, 195, 66, 16, 171, 216, 39, 236, 9, 18, 98, 146, 180, 67, 233, 91, 223, 194, 74, 59, 232, 75, 8, 212, 69, 50, 149, 151, 152, 188, 119, 134, 125, 133, 121, 2, 77, 115, 90, 79, 207, 202, 161, 110, 80, 173, 123, 28, 255, 95, 25, 1, 70, 99, 92, 218, 168, 108, 58, 87, 122, 76, 101, 239, 4, 230, 251, 240, 53, 178, 6, 213, 135, 97, 89, 246, 159, 185, 68, 174, 217, 160, 85, 138, 72, 169, 32, 197, 145, 15, 163, 177, 11, 143, 10, 199, 155, 65, 235, 84, 248, 24, 221, 106, 114, 196, 193, 46, 244, 165, 183, 55, 187, 41, 192, 130, 82, 111, 127, 162, 204, 57, 222, 254, 142, 42, 38, 224, 83, 5, 103, 153, 157, 73, 52, 78, 219, 118, 227, 33, 19, 27, 54, 7, 205, 164, 200, 62, 215, 179, 116, 56, 140, 64, 245, 150, 137, 43, 71, 34, 36, 184, 243, 141, 158, 203, 189, 94, 132, 175, 147, 166, 249, 117, 102, 40, 136, 21, 237, 238, 81, 220, 14, 242, 88, 44, 100, 35, 112, 209, 252, 105, 253, 129, 47, 191, 61, 247, 93, 49, 51, 250, 104, 126, 211, 201}

var S8_0_INV = [256]uint8{118, 40, 179, 239, 254, 168, 250, 172, 177, 237, 114, 44, 59, 111, 57, 109, 89, 5, 91, 7, 18, 68, 209, 133, 154, 196, 158, 192, 211, 135, 22, 64, 65, 21, 130, 212, 193, 157, 195, 159, 134, 208, 67, 23, 2, 92, 6, 88, 110, 56, 106, 60, 43, 119, 238, 176, 171, 255, 169, 253, 234, 180, 41, 117, 24, 126, 223, 187, 160, 206, 164, 202, 221, 185, 28, 122, 103, 11, 101, 9, 53, 81, 55, 83, 76, 34, 141, 225, 244, 146, 240, 150, 143, 227, 72, 38, 33, 77, 226, 140, 145, 245, 147, 247, 230, 136, 35, 79, 82, 52, 86, 48, 14, 96, 10, 100, 123, 31, 190, 216, 203, 167, 201, 165, 186, 220, 121, 29, 69, 17, 132, 210, 197, 153, 199, 155, 128, 214, 71, 19, 4, 90, 0, 94, 104, 62, 108, 58, 47, 115, 232, 182, 175, 251, 173, 249, 236, 178, 45, 113, 112, 46, 183, 235, 248, 174, 252, 170, 181, 233, 116, 42, 63, 107, 61, 105, 93, 1, 95, 3, 20, 66, 213, 129, 156, 194, 152, 198, 215, 131, 16, 70, 30, 120, 219, 191, 166, 200, 162, 204, 217, 189, 26, 124, 99, 15, 97, 13, 49, 85, 51, 87, 74, 36, 137, 229, 242, 148, 246, 144, 139, 231, 78, 32, 37, 73, 228, 138, 149, 241, 151, 243, 224, 142, 39, 75, 84, 50, 80, 54, 8, 102, 12, 98, 127, 27, 184, 222, 207, 163, 205, 161, 188, 218, 125, 25}
var S8_1_INV = [256]uint8{84, 56, 149, 251, 230, 130, 228, 128, 145, 255, 86, 58, 33, 71, 37, 67, 125, 19, 121, 23, 14, 106, 201, 175, 188, 208, 190, 210, 205, 171, 12, 104, 105, 15, 174, 202, 213, 187, 209, 191, 172, 200, 109, 11, 20, 120, 22, 122, 70, 34, 68, 32, 61, 83, 252, 144, 129, 231, 133, 227, 254, 146, 57, 87, 52, 96, 245, 163, 182, 234, 180, 232, 241, 167, 54, 98, 113, 47, 117, 43, 29, 75, 25, 79, 94, 2, 153, 199, 220, 136, 222, 138, 157, 195, 92, 0, 7, 89, 194, 158, 139, 221, 143, 217, 192, 156, 3, 93, 72, 28, 74, 30, 42, 118, 40, 116, 99, 53, 160, 244, 239, 177, 235, 181, 162, 246, 103, 49, 111, 9, 170, 206, 211, 189, 215, 185, 168, 204, 107, 13, 16, 124, 18, 126, 66, 38, 64, 36, 59, 85, 248, 148, 135, 225, 131, 229, 250, 150, 63, 81, 80, 60, 147, 253, 226, 134, 224, 132, 151, 249, 82, 62, 39, 65, 35, 69, 123, 21, 127, 17, 10, 110, 207, 169, 184, 212, 186, 214, 203, 173, 8, 108, 48, 100, 243, 165, 178, 238, 176, 236, 247, 161, 50, 102, 119, 41, 115, 45, 27, 77, 31, 73, 90, 6, 159, 193, 216, 140, 218, 142, 155, 197, 88, 4, 1, 95, 198, 154, 141, 219, 137, 223, 196, 152, 5, 91, 76, 24, 78, 26, 46, 114, 44, 112, 101, 51, 164, 240, 233, 183, 237, 179, 166, 242, 97, 55}
var S8_2_INV = [256]uint8{156, 240, 91, 53, 44, 72, 46, 74, 95, 49, 158, 242, 235, 141, 239, 137, 183, 217, 179, 221, 198, 162, 7, 97, 118, 26, 116, 24, 3, 101, 196, 160, 167, 193, 102, 2, 31, 113, 27, 117, 100, 0, 163, 197, 222, 178, 220, 176, 140, 232, 142, 234, 243, 157, 52, 88, 75, 45, 79, 41, 54, 90, 247, 153, 248, 172, 61, 107, 120, 36, 122, 38, 57, 111, 250, 174, 189, 227, 185, 231, 209, 135, 213, 131, 146, 206, 81, 15, 18, 70, 16, 68, 85, 11, 144, 204, 207, 145, 14, 82, 71, 17, 67, 21, 12, 80, 203, 149, 134, 210, 132, 208, 228, 184, 230, 186, 171, 253, 108, 56, 35, 125, 39, 121, 110, 58, 175, 249, 161, 199, 98, 6, 25, 119, 29, 115, 96, 4, 165, 195, 218, 182, 216, 180, 136, 236, 138, 238, 245, 155, 48, 92, 77, 43, 73, 47, 50, 94, 241, 159, 152, 244, 93, 51, 40, 76, 42, 78, 89, 55, 154, 246, 237, 139, 233, 143, 177, 223, 181, 219, 194, 166, 1, 103, 114, 30, 112, 28, 5, 99, 192, 164, 252, 168, 59, 109, 124, 32, 126, 34, 63, 105, 254, 170, 187, 229, 191, 225, 215, 129, 211, 133, 150, 202, 87, 9, 22, 66, 20, 64, 83, 13, 148, 200, 201, 151, 10, 86, 65, 23, 69, 19, 8, 84, 205, 147, 130, 214, 128, 212, 224, 188, 226, 190, 173, 251, 104, 60, 37, 123, 33, 127, 106, 62, 169, 255}
var S8_3_INV = [256]uint8{139, 229, 72, 36, 248, 156, 254, 154, 78, 34, 137, 231, 59, 93, 57, 95, 97, 15, 99, 13, 209, 183, 22, 114, 166, 202, 160, 204, 16, 116, 211, 181, 180, 208, 117, 19, 205, 163, 207, 161, 119, 17, 178, 214, 12, 96, 10, 102, 90, 62, 92, 56, 226, 142, 39, 73, 159, 249, 157, 251, 37, 75, 228, 136, 236, 184, 45, 123, 173, 243, 175, 241, 47, 121, 234, 190, 108, 48, 106, 54, 2, 86, 4, 80, 130, 222, 71, 25, 199, 145, 197, 147, 69, 27, 132, 216, 221, 131, 28, 64, 148, 192, 146, 198, 26, 70, 223, 129, 85, 3, 87, 1, 55, 105, 53, 107, 191, 233, 122, 46, 242, 174, 244, 168, 124, 40, 189, 235, 176, 212, 115, 21, 203, 165, 201, 167, 113, 23, 182, 210, 8, 100, 14, 98, 94, 58, 88, 60, 230, 138, 33, 79, 153, 255, 155, 253, 35, 77, 224, 140, 141, 227, 76, 32, 252, 152, 250, 158, 74, 38, 143, 225, 61, 91, 63, 89, 103, 9, 101, 11, 215, 177, 18, 118, 162, 206, 164, 200, 20, 112, 213, 179, 232, 188, 43, 125, 171, 245, 169, 247, 41, 127, 238, 186, 104, 52, 110, 50, 6, 82, 0, 84, 134, 218, 65, 31, 193, 151, 195, 149, 67, 29, 128, 220, 219, 133, 24, 68, 144, 196, 150, 194, 30, 66, 217, 135, 83, 5, 81, 7, 49, 111, 51, 109, 185, 239, 126, 42, 246, 170, 240, 172, 120, 44, 187, 237}
var S8_4_INV = [256]uint8{9, 101, 85, 50, 114, 179, 120, 193, 72, 58, 144, 142, 29, 34, 232, 139, 53, 35, 59, 190, 5, 227, 2, 36, 151, 100, 3, 191, 97, 4, 39, 15, 136, 189, 209, 237, 210, 22, 176, 56, 225, 163, 175, 207, 235, 27, 157, 244, 33, 249, 75, 250, 184, 118, 192, 161, 201, 171, 108, 69, 8, 246, 197, 0, 203, 147, 52, 63, 128, 74, 102, 208, 134, 183, 68, 71, 111, 86, 185, 89, 94, 230, 166, 178, 149, 132, 25, 109, 234, 124, 88, 65, 104, 248, 217, 99, 16, 123, 60, 103, 236, 112, 224, 180, 252, 241, 153, 20, 107, 32, 93, 167, 238, 38, 154, 87, 200, 223, 187, 80, 30, 84, 110, 96, 40, 82, 253, 168, 37, 243, 165, 48, 218, 83, 81, 122, 226, 206, 133, 42, 202, 213, 174, 143, 12, 138, 61, 220, 28, 76, 205, 77, 78, 181, 49, 146, 31, 182, 214, 126, 131, 92, 169, 140, 195, 159, 221, 13, 106, 135, 43, 54, 41, 95, 129, 219, 10, 141, 119, 199, 62, 46, 14, 160, 211, 127, 21, 162, 79, 216, 23, 245, 164, 156, 67, 51, 155, 137, 19, 145, 196, 255, 91, 215, 170, 194, 26, 90, 7, 239, 18, 254, 73, 121, 1, 198, 55, 130, 105, 186, 231, 152, 172, 66, 177, 17, 6, 188, 44, 45, 115, 11, 70, 64, 47, 148, 57, 228, 229, 113, 117, 24, 233, 212, 158, 204, 125, 247, 150, 222, 251, 116, 240, 242, 173, 98}

func S(x *[8]uint8) {
	x[0] = S8_0[x[0]]
	x[1] = S8_1[x[1]]
	x[2] = S8_2[x[2]]
	x[3] = S8_3[x[3]]
	x[4] = S8_1[x[4]]
	x[5] = S8_2[x[5]]
	x[6] = S8_3[x[6]]
	x[7] = S8_0[x[7]]
}

func SInv(x *[8]uint8) {
	x[0] = S8_0_INV[x[0]]
	x[1] = S8_1_INV[x[1]]
	x[2] = S8_2_INV[x[2]]
	x[3] = S8_3_INV[x[3]]
	x[4] = S8_1_INV[x[4]]
	x[5] = S8_2_INV[x[5]]
	x[6] = S8_3_INV[x[6]]
	x[7] = S8_0_INV[x[7]]
}

func P(x *[8]uint8) {

	z0, z1, z2, z3, z4, z5, z6, z7 := x[0], x[1], x[2], x[3], x[4], x[5], x[6], x[7]

	x[0] = z2 ^ z3 ^ z4 ^ z6 ^ z7
	x[1] = z0 ^ z1 ^ z3 ^ z4 ^ z7
	x[2] = z0 ^ z1 ^ z4 ^ z5 ^ z6
	x[3] = z1 ^ z2 ^ z3 ^ z5 ^ z6
	x[4] = z0 ^ z2 ^ z3 ^ z6 ^ z7
	x[5] = z0 ^ z3 ^ z4 ^ z5 ^ z7
	x[6] = z0 ^ z1 ^ z2 ^ z4 ^ z5
	x[7] = z1 ^ z2 ^ z5 ^ z6 ^ z7
}

func keySchedule(key [8]uint8) (keys [6][8]uint8) {

	var a, b *[8]uint8

	a = &keys[0]
	for i := 0; i < 8; i++ {
		a[i] = key[i]
	}

	for i := 1; i < 6; i++ {

		a = &keys[i-1]
		b = &keys[i]

		for j := 0; j < 8; j++ {
			b[j] = a[j]
		}

		S(b)
		P(b)
		S(b)
	}

	return keys
}

func keyScheduleInv(key [8]uint8) (keys [6][8]uint8) {
	var a, b *[8]uint8

	a = &keys[5]
	for i := 0; i < 8; i++ {
		a[i] = key[i]
	}

	for i := 4; i >= 0; i-- {

		a = &keys[i+1]
		b = &keys[i]

		for j := 0; j < 8; j++ {
			b[j] = a[j]
		}

		SInv(b)
		P(b)
		SInv(b)
	}

	return keys
}

func K(a, b *[8]uint8) {
	for i := 0; i < 8; i++ {
		a[i] ^= b[i]
	}
}

func Encrypt(x *[8]uint8, k [6][8]uint8, y *[8]uint8) {

	for i := 0; i < 8; i++ {
		y[i] = x[i]
	}

	for i := 0; i < 4; i++ {
		K(y, &k[i])
		S(y)
		P(y)
	}

	K(y, &k[4])

	for i := 0; i < 8; i++ {
		y[i] = S8_4[y[i]]
	}

	K(y, &k[5])

}

func Decrypt(x *[8]uint8, k [6][8]uint8, y *[8]uint8) {

	for i := 0; i < 8; i++ {
		y[i] = x[i]
	}

	K(y, &k[5])

	for i := 0; i < 8; i++ {
		y[i] = S8_4_INV[y[i]]
	}

	K(y, &k[4])

	for i := 3; i >= 0; i-- {
		P(y)
		SInv(y)
		K(y, &k[i])
	}
}

func Abs(x int64) int64 {
	return int64(math.Abs(float64(x)))
}

func RecoverKey(k [8]uint8) bool {

	var keys = keySchedule(k)

	var plaintexts, ciphertexts [][8]uint8

	n := 150000

	plaintexts = make([][8]uint8, n)
	ciphertexts = make([][8]uint8, n)

	tmp := make([]byte, 8)

	for i := 0; i < n; i++ {
		rand.Read(tmp)

		for j := 0; j < 8; j++ {
			plaintexts[i][j] = uint8(tmp[j])
		}

		Encrypt(&plaintexts[i], keys, &ciphertexts[i])
	}

	type candidate struct {
		bias int64
		key  uint8
	}

	k_guesses := make([][]*candidate, 8)

	var candidates = make([]*candidate, 256)

	var p, c uint8

	for i := 0; i < 8; i++ {

		// We have two valide linear paths
		var k_0 [256]int64
		var k_1 [256]int64

		for j := 0; j < 256; j++ {
			k_0[j] = -(int64(n) / 2)
			k_1[j] = -(int64(n) / 2)
		}

		for j := 0; j < n; j++ {

			p = plaintexts[j][i]

			for k := 0; k < 256; k++ {

				ki := uint8(k)

				// partial decryption of the byte i
				c = S8_4_INV[ciphertexts[j][i]^ki]

				// Increments when the linear paths hold
				k_0[ki] += int64(utils.P[(p&uint8(0x41))^(c&uint8(0x51))])
				k_1[ki] += int64(utils.P[(p&uint8(0x41))^(c&uint8(0x14))])
			}
		}

		// One has a high probability of occurence, while the other has a high probabily of not occuring.
		// We count the bias with their respective absolute value.
		for j := 0; j < 256; j++ {
			candidates[j] = new(candidate)
			candidates[j].bias = Abs(k_0[j]) + Abs(k_1[j])
			candidates[j].key = uint8(j)
		}

		// Sort keys candidates by best bias
		sort.Slice(candidates, func(i, j int) bool {
			return candidates[i].bias > candidates[j].bias
		})

		k_guesses[i] = make([]*candidate, 5)
		for j := 0; j < 5; j++ {
			k_guesses[i][j] = candidates[j]
		}
	}

	/*
		for i := range k_guesses{
			for j := range k_guesses[i]{
				fmt.Printf("%d %d %02x \n", i, k_guesses[i][j].bias , k_guesses[i][j].key)
			}
			fmt.Println()
		}
	*/

	for k0 := range k_guesses[0] {
		for k1 := range k_guesses[1] {
			for k2 := range k_guesses[2] {
				for k3 := range k_guesses[3] {
					for k4 := range k_guesses[4] {
						for k5 := range k_guesses[5] {
							for k6 := range k_guesses[6] {
								for k7 := range k_guesses[7] {

									keys := keyScheduleInv([8]uint8{k_guesses[0][k0].key, k_guesses[1][k1].key,
										k_guesses[2][k2].key, k_guesses[3][k3].key,
										k_guesses[4][k4].key, k_guesses[5][k5].key,
										k_guesses[6][k6].key, k_guesses[7][k7].key})

									var test [8]uint8

									state := false
									for i := 0; i < 10; i++ {

										Encrypt(&plaintexts[i], keys, &test)

										state = state || equalSlice(ciphertexts[i], test)
									}

									if state {
										return true
									}

								}

							}

						}

					}

				}

			}

		}
	}

	return false
}
