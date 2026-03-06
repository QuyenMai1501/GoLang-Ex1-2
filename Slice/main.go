package main

import "fmt"

func main() {
	var ids []int
	var names []string
	var prices []float64

	ids = []int{1, 2, 3}
	names = []string{"Cà phê đen", "Cà phê sữa", "Trà đào"}
	prices = []float64{28000, 32000, 45000}

	var luaChon int

	for {
		fmt.Println("\n--- QUẢN LÝ QUÁN CÀ PHÊ ---")
		fmt.Println("1. Thêm món mới")
		fmt.Println("2. Sửa món")
		fmt.Println("3. Xóa món")
		fmt.Println("4. Hiển thị menu")
		fmt.Println("0. Thoát")
		fmt.Print("Chọn: ")

		fmt.Scan(&luaChon)

		if luaChon == 0 {
			fmt.Println("Thoát chương trình.")
			break
		}

		switch luaChon {
		case 1:
			themMon(&ids, &names, &prices)
		case 2:
			suaMon(ids, &names, &prices)
		case 3:
			xoaMon(&ids, &names, &prices)
		case 4:
			hienThiMenu(ids, names, prices)
		default:
			fmt.Println("Lựa chọn không hợp lệ!")
		}
	}
}

func hienThiMenu(ids []int, names []string, prices []float64) {
	if len(names) == 0 {
		fmt.Println("Menu hiện đang trống!")
		return
	}

	fmt.Println("\n=== MENU HIỆN TẠI ===")
	fmt.Printf("%-5s %-20s %s\n", "ID", "Tên món", "Giá")
	fmt.Println("-----------------------------")

	for i := 0; i < len(names); i++ {
		fmt.Printf("%-5d %-20s %.0f\n", ids[i], names[i], prices[i])
	}

	fmt.Printf("Tổng cộng: %d món\n", len(names))
}

func themMon(ids *[]int, names *[]string, prices *[]float64) {
	var ten string
	var gia float64

	fmt.Print("Nhập tên món mới: ")
	fmt.Scan(&ten)

	fmt.Print("Nhập giá món mới: ")
	fmt.Scan(&gia)

	var idMoi int
	if len(*ids) == 0 {
		idMoi = 1
	} else {
		idMoi = (*ids)[len(*ids)-1] + 1
	}

	*ids = append(*ids, idMoi)
	*names = append(*names, ten)
	*prices = append(*prices, gia)

	fmt.Printf("Đã thêm: %s (ID: %d)\n", ten, idMoi)
}

func suaMon(ids []int, names *[]string, prices *[]float64) {
	var id int
	fmt.Print("Nhập ID món cần sửa: ")
	fmt.Scan(&id)

	for i := 0; i < len(ids); i++ {
		if ids[i] == id {
			fmt.Printf("Món hiện tại: %s - %.0f đồng\n", (*names)[i], (*prices)[i])

			var tenMoi string
			fmt.Print("Tên mới (0 để giữ nguyên): ")
			fmt.Scan(&tenMoi)
			if tenMoi != "0" {
				(*names)[i] = tenMoi
			}

			var giaMoi float64
			fmt.Print("Giá mới (0 để giữ nguyên): ")
			fmt.Scan(&giaMoi)
			if giaMoi != 0 {
				(*prices)[i] = giaMoi
			}

			fmt.Println("Đã cập nhật")
			return
		}
	}

	fmt.Printf("Không tìm thấy ID %d\n", id)
}

func xoaMon(ids *[]int, names *[]string, prices *[]float64) {
	var id int
	fmt.Print("Nhập ID món cần xóa: ")
	fmt.Scan(&id)

	for i := 0; i < len(*ids); i++ {
		if (*ids)[i] == id {
			for j := i; j < len(*ids)-1; j++ {
				(*ids)[j] = (*ids)[j+1]
				(*names)[j] = (*names)[j+1]
				(*prices)[j] = (*prices)[j+1]
			}

			*ids = (*ids)[:len(*ids)-1]
			*names = (*names)[:len(*names)-1]
			*prices = (*prices)[:len(*prices)-1]

			fmt.Println("Đã xóa")
			return
		}
	}

	fmt.Printf("Không tìm thấy ID %d\n", id)
}
