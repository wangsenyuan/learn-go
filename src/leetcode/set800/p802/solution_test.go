package p802

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, graph [][]int, expect []int) {
	res := eventualSafeNodes(graph)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("sample %v, expect %v, but got %v", graph, expect, res)
	}
}

func TestSample1(t *testing.T) {
	graph := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	res := []int{2, 4, 5, 6}
	runSample(t, graph, res)
}

func TestSample2(t *testing.T) {
	graph := [][]int{{}, {0, 2, 3, 4}, {3}, {4}, {}}
	res := []int{0, 1, 2, 3, 4}
	runSample(t, graph, res)
}

func TestSample3(t *testing.T) {
	graph := [][]int{{1, 3, 7, 9}, {1, 8}, {7}, {5, 6, 7, 8}, {5, 6, 7, 8}, {8}, {7, 8, 9}, {8}, {9}, {}}
	res := []int{2, 3, 4, 5, 6, 7, 8, 9}
	runSample(t, graph, res)
}

func TestSample4(t *testing.T) {
	graph := [][]int{{62}, {17, 24, 62, 155, 160, 168}, {8, 25, 33, 87, 145, 183}, {192}, {47, 59, 64, 79, 143, 159}, {58, 62, 87, 88, 121, 166}, {164}, {107, 120}, {38, 46, 113, 157}, {8, 30, 104, 143, 162, 179}, {42, 69}, {30, 43, 63, 95, 122, 158}, {79, 136, 197}, {50, 52, 109, 125, 139, 155}, {19, 22, 71, 108, 158, 191}, {141}, {}, {80, 108, 131, 141, 166, 171}, {54}, {}, {57, 108, 116, 138, 166, 186}, {}, {116, 131, 162}, {29, 67, 90, 92, 186}, {129}, {46, 54, 150, 174, 176}, {37, 76, 99, 167}, {119, 173}, {166}, {83, 127}, {61, 69, 114}, {33, 67, 185}, {93, 177, 198}, {49, 67, 119, 127, 132, 188}, {46, 98, 125, 127}, {79, 84}, {83, 94}, {119}, {86, 156}, {103, 133, 171}, {80}, {}, {141, 188}, {68, 90, 127, 145, 149, 190}, {78, 99, 116}, {131, 138, 139, 186}, {77, 101, 131, 137, 145, 150}, {83, 99}, {}, {86, 166}, {}, {126, 140, 146, 171}, {}, {104, 106, 111, 142, 146, 193}, {64, 65, 101, 103}, {32, 67, 69, 71, 103, 157}, {7, 57, 82, 151}, {131, 181}, {120}, {63, 68, 71, 84, 91, 113}, {139}, {88, 120, 149, 158}, {103, 150, 153, 154, 173}, {81, 147, 190, 195}, {93, 163, 182}, {}, {132}, {}, {139, 148, 164}, {83, 145, 148, 152, 182, 192}, {}, {26, 185}, {84, 95, 110, 145, 160, 181}, {77, 129, 131, 154, 197}, {194}, {92, 180}, {80, 136, 140}, {83, 127, 137}, {}, {82, 89, 104, 130, 165, 188}, {110, 147, 156}, {92, 113, 155, 176, 190}, {123, 126, 139, 157, 189}, {}, {91, 101, 125, 130, 150, 176}, {125}, {124, 127, 198}, {106, 135, 147}, {}, {116, 126, 145, 152, 193, 195}, {117, 130, 134, 181, 185}, {150}, {}, {101, 112, 146, 159, 173}, {27, 111, 123, 137, 170, 182, 186}, {109, 110, 124, 142}, {184}, {104, 105, 144, 150, 155, 156}, {127}, {}, {}, {106, 145, 164, 176}, {120, 144, 151, 155, 187, 199}, {130, 180}, {108, 147}, {124, 138, 159, 161, 173}, {146}, {138, 181}, {}, {190}, {134, 181}, {}, {122, 156, 167, 168, 169, 187}, {120, 135, 161, 162, 169, 185}, {192}, {147, 173, 181, 195}, {}, {}, {136, 143, 161, 176, 181}, {136, 146, 147, 149, 165, 168}, {}, {136, 137, 169, 179, 188, 198}, {151}, {131, 136, 144, 170}, {31, 168}, {141, 150, 158, 159, 179}, {157, 177}, {132, 178, 183}, {170, 178, 189}, {132, 157, 185}, {143, 169, 175}, {160}, {134, 154, 195}, {137, 146, 151, 156, 174, 199}, {145, 151, 180}, {149, 160, 189, 191, 198}, {150, 162, 168, 181, 188}, {155, 157, 168, 184, 189, 192}, {157, 169, 176, 182, 187, 194}, {148, 151, 161, 184, 186}, {176, 185, 198}, {154, 159, 161, 173, 175, 182}, {}, {32, 151, 156, 159, 176, 180, 184}, {}, {151, 185, 196}, {162, 165, 180, 185, 191, 194}, {162, 169}, {156, 173, 179, 184, 199}, {151, 158, 165, 185, 197}, {154, 161, 169, 188}, {197}, {153, 167, 197}, {159, 190, 191}, {155, 160, 161, 172, 178, 186}, {199}, {180, 193, 198}, {165, 182, 190, 192, 198}, {}, {}, {161, 185, 199}, {169, 179, 181, 184, 187}, {169, 170, 174, 176, 193}, {169, 170, 174, 190, 199}, {184}, {169, 173, 191, 196}, {170, 177, 178, 181, 198}, {195, 197, 199}, {182}, {175, 181, 186, 194}, {177, 179, 184, 187, 192, 195}, {173, 188, 197, 199}, {173, 177, 178, 184}, {165, 189}, {}, {188}, {178, 184, 192, 196, 197}, {}, {186, 187, 188, 191, 198}, {183, 189, 191}, {184, 190, 195, 196, 199}, {182, 184, 189, 195, 198, 199}, {189, 191, 192}, {187}, {185, 189, 191, 193, 195}, {189, 191, 193, 195, 196}, {189, 191, 194}, {189, 192, 193, 196}, {191, 192, 195, 196, 198}, {190, 191, 192, 194, 196, 197}, {35, 191, 194, 196, 197, 198}, {194, 195, 196, 198}, {193, 194, 195, 196, 197, 199}, {53, 197, 198, 199}, {195, 196, 197, 198, 199}, {}, {197, 198, 199}, {198, 199}, {199}, {}}
	res := []int{2, 3, 4, 5, 6, 7, 8, 9}
	runSample(t, graph, res)
}
