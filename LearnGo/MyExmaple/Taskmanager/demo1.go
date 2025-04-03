package taskmanager

import "fmt"

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

func AddTask(tasks *[]Task, title string, desc string) {

	// ID oluştur: önceki en son ID'yi bul, bir arttır
	newID := 1
	if len(*tasks) > 0 {
		lastTask := (*tasks)[len(*tasks)-1]
		newID = lastTask.ID + 1
	}

	// Yeni görev oluştur
	newTask := Task{
		ID:          newID,
		Title:       title,
		Description: desc,
		Done:        false,
	}

	// Slice'a ekle
	*tasks = append(*tasks, newTask)

	fmt.Println("Görev eklendi:", newTask.Title)
}

func ListTasks(tasks []Task) {

	if len(tasks) == 0 {
		fmt.Println("Hiç görev yok.")
		return
	}

	fmt.Println("----- Görev Listesi -----")
	for i := 0; i < len(tasks); i++ {
		task := tasks[i]
		status := "❌ Tamamlanmadı"
		if task.Done {
			status = "✅ Tamamlandı"
		}

		fmt.Printf("ID: %d\nBaşlık: %s\nAçıklama: %s\nDurum: %s\n\n",
			task.ID, task.Title, task.Description, status)
	}
}

func MarkAsDone(tasks *[]Task, id int) {
	for i := 0; i < len(*tasks); i++ {
		if (*tasks)[i].ID == id {
			(*tasks)[i].Done = true
			fmt.Printf("Görev %d tamamlandı olarak işaretlendi.\n", id)
			return
		}
	}
	fmt.Printf("ID %d olan görev bulunamadı.\n", id)
}
