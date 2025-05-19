package main

import "fmt"

// Bağlı listenin düğümünü temsil eden yapı
type ListNode struct {
	Val  int
	Next *ListNode
}

// İki sıralı bağlı listeyi birleştirir
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // Geçici başlangıç düğümü
	tail := dummy        // tail, sonucu oluştururken kuyruğu takip eder

	// Her iki listede de eleman varken karşılaştırarak ekle
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// Kalan listeyi doğrudan ekle
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next // dummy.Next, birleşmiş sıralı listenin başlangıcıdır
}

// Slice'tan bağlı liste oluşturan yardımcı fonksiyon
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

// Bağlı listeyi yazdıran fonksiyon
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

// Ana fonksiyon (test)
func main() {
	// İki sıralı bağlı liste oluştur
	list1 := createList([]int{1, 2, 4})
	list2 := createList([]int{1, 3, 4})

	// Birleştir
	merged := mergeTwoLists(list1, list2)

	// Yazdır
	fmt.Print("Birleşmiş Liste: ")
	printList(merged)
}
