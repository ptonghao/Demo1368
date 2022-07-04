/*
 * @Author: Jimpu
 * @Description: main
 */

package quick_sort

// quickSort 快速排序
func QuickSort(arr []int, left, right int) {
	if left < right {
		pivot := arr[left]
		j := left
		for i := left; i < right; i++ {
			if arr[i] < pivot {
				j++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		arr[left], arr[j] = arr[j], arr[left]
		QuickSort(arr, left, j)
		QuickSort(arr, j+1, right)
	}
}
