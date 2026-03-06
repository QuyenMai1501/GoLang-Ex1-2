package main

import "fmt"

func main() {
	menu := make(map[int][]string)

	menu[1] = []string{"Cà phê đen", "28000"}
	menu[2] = []string{"Cà phê sữa", "32000"}
	menu[3] = []string{"Trà đào", "45000"}

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
			themMon(menu)
		case 2:
			suaMon(menu)
		case 3:
			xoaMon(menu)
		case 4:
			hienThiMenu(menu)
		default:
			fmt.Println("Lựa chọn không hợp lệ!")
		}
	}
}

func hienThiMenu(menu map[int][]string) {
	if len(menu) == 0 {
		fmt.Println("Menu hiện đang trống!")
		return
	}

	fmt.Println("\n=== MENU HIỆN TẠI ===")
	fmt.Printf("%-5s %-20s %s\n", "ID", "Tên món", "Giá")
	fmt.Println("-----------------------------")

	for id, info := range menu {
		ten := info[0]
		gia := info[1]
		fmt.Printf("%3d %-25s %s\n", id, ten, gia)
	}

	fmt.Printf("Tổng cộng: %d món\n", len(menu))
}

func themMon(menu map[int][]string) {
	var id int
	var ten string
	var gia string

	fmt.Print("Nhập ID món: ")
	fmt.Scan(&id)

	if _, tonTai := menu[id]; tonTai {
		fmt.Println("ID này đã tồn tại!")
		return
	}

	fmt.Print("Nhập tên món mới: ")
	fmt.Scan(&ten)

	fmt.Print("Nhập giá món mới: ")
	fmt.Scan(&gia)

	menu[id] = []string{ten, gia}
	fmt.Printf("Đã thêm món: %s (ID: %d, Giá: %s)\n", ten, id, gia)
}

func suaMon(menu map[int][]string) {
	var id int
	fmt.Print("Nhập ID món cần sửa: ")
	fmt.Scan(&id)

	info, tonTai := menu[id]
	if !tonTai {
		fmt.Println("Không tìm thấy món có ID này.")
		return
	}

	tenCu := info[0]
	giaCu := info[1]
	fmt.Printf("Món hiện tại: %s - %s đồng\n", tenCu, giaCu)

	var tenMoi string
	fmt.Print("Tên mới (0 để giữ nguyên): ")
	fmt.Scan(&tenMoi)
	if tenMoi != "0" {
		info[0] = tenMoi
	}

	var giaMoi string
	fmt.Print("Giá mới (0 để giữ nguyên): ")
	fmt.Scan(&giaMoi)
	if giaMoi != "0" {
		info[1] = giaMoi
	}

	menu[id] = info
	fmt.Println("Đã cập nhật.")
}

func xoaMon(menu map[int][]string) {
	var id int
	fmt.Print("Nhập ID món cần xóa: ")
	fmt.Scan(&id)

	if _, tonTai := menu[id]; tonTai {
		delete(menu, id)
		fmt.Println("Đã xóa món thành công.")
	} else {
		fmt.Println("Không tìm thấy món có ID này.")
	}
}
