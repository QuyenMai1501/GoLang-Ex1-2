package main

import "fmt"

const MAX_ITEMS = 100

func hienThiMenu() {
	fmt.Println("\n=== QUẢN LÝ MENU QUÁN CAFE ===")
	fmt.Println("1. Thêm món mới")
	fmt.Println("2. Sửa món")
	fmt.Println("3. Xóa món")
	fmt.Println("4. Hiển thị menu")
	fmt.Println("0. Thoát")
	fmt.Print("Chọn: ")
}

func hienThiTatCaMenu(ids [MAX_ITEMS]int, names [MAX_ITEMS]string, prices [MAX_ITEMS]float64, count int) {
	if count == 0 {
		fmt.Println("Menu hiện đang trống!")
		return
	}

	fmt.Println("\n=== MENU HIỆN TẠI ===")
	fmt.Printf("%-5s %-20s %s\n", "ID", "Tên món", "Giá")
	fmt.Println("-----------------------------")

	for i := 0; i < count; i++ {
		fmt.Printf("%-5d %-20s %.0f\n", ids[i], names[i], prices[i])
	}
	fmt.Printf("Tổng cộng: %d món\n", count)
}

func timViTriTheoID(ids [MAX_ITEMS]int, count int, id int) int {
	for i := 0; i < count; i++ {
		if ids[i] == id {
			return i
		}
	}
	return -1
}

func themMonMoi(ids *[MAX_ITEMS]int, names *[MAX_ITEMS]string, prices *[MAX_ITEMS]float64, count *int) {
	if *count >= MAX_ITEMS {
		fmt.Println("Menu đã đầy! Không thể thêm nữa.")
		return
	}

	var newID int
	var newName string
	var newPrice float64

	if *count == 0 {
		newID = 1
	} else {
		newID = ids[*count-1] + 1
	}

	fmt.Print("Nhập tên món mới: ")
	fmt.Scan(&newName)

	fmt.Print("Nhập giá món mới: ")
	fmt.Scan(&newPrice)

	ids[*count] = newID
	names[*count] = newName
	prices[*count] = newPrice
	*count++

	fmt.Printf("Đã thêm: %s (ID: %d, Giá: %.0f)\n", newName, newID, newPrice)
}

func suaMon(ids *[MAX_ITEMS]int, names *[MAX_ITEMS]string, prices *[MAX_ITEMS]float64, count int) {
	var id int
	fmt.Print("Nhập ID món cần sửa: ")
	fmt.Scan(&id)

	idx := timViTriTheoID(*ids, count, id)
	if idx == -1 {
		fmt.Printf("Không tìm thấy ID %d\n", id)
		return
	}

	fmt.Printf("Mon hien tai: %s - %.0f\n", names[idx], prices[idx])

	fmt.Print("Tên mới (0 để giữ nguyên): ")
	var tempName string
	fmt.Scan(&tempName)
	if tempName != "0" {
		names[idx] = tempName
	}

	fmt.Print("Giá mới (0 để giữ nguyên): ")
	var tempPrice float64
	fmt.Scan(&tempPrice)
	if tempPrice != 0 {
		prices[idx] = tempPrice
	}

	fmt.Println("Đã cập nhật!")
}

func xoaMon(ids *[MAX_ITEMS]int, names *[MAX_ITEMS]string, prices *[MAX_ITEMS]float64, count *int) {
	var id int
	fmt.Print("Nhập ID món cần xóa: ")
	fmt.Scan(&id)

	idx := timViTriTheoID(*ids, *count, id)
	if idx == -1 {
		fmt.Printf("Không tìm thấy ID %d\n", id)
		return
	}

	for j := idx; j < *count-1; j++ {
		ids[j] = ids[j+1]
		names[j] = names[j+1]
		prices[j] = prices[j+1]
	}
	*count--
	fmt.Println("Đã xóa!")
}

func main() {
	var ids [MAX_ITEMS]int
	var names [MAX_ITEMS]string
	var prices [MAX_ITEMS]float64
	var count int = 3

	ids[0] = 1
	names[0] = "Cà phê đen"
	prices[0] = 35000

	ids[1] = 2
	names[1] = "Cà phê sữa"
	prices[1] = 40000

	ids[2] = 3
	names[2] = "Trà đào"
	prices[2] = 45000

	var choice int

	for {
		hienThiMenu()
		fmt.Scan(&choice)

		if choice == 0 {
		fmt.Println("Thoát chương trình.")
		break
	}

		switch choice {
		case 1:
			themMonMoi(&ids, &names, &prices, &count)
		case 2:
			suaMon(&ids, &names, &prices, count)
		case 3:
			xoaMon(&ids, &names, &prices, &count)
		case 4:
			hienThiTatCaMenu(ids, names, prices, count)
		default:
			fmt.Println("Lựa chọn không hợp lệ")
		}
	}
}
